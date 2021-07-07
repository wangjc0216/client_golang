package m3

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/prometheus/client_golang/prometheus/m3/prompb"
	dto "github.com/prometheus/client_model/go"

	"github.com/prometheus/client_golang/prometheus"
)

type HTTPDoer interface {
	Do(*http.Request) (*http.Response, error)
}

//M3客户端
type M3Pusher struct {
	error error

	//url, job string
	url, namespace, httpMethod string

	grouping map[string]string

	gatherers  prometheus.Gatherers
	registerer prometheus.Registerer

	client HTTPDoer

	useBasicAuth       bool
	username, password string

	ctx    context.Context
	cancel context.CancelFunc
}

//初始化选项，包括本地缓存空间设置
type M3PusherOpt func(*M3Pusher)

func WithHttpPutOpt() M3PusherOpt {
	return func(p *M3Pusher) {
		p.httpMethod = http.MethodPut
	}
}

func WithHttpPostOpt() M3PusherOpt {
	return func(p *M3Pusher) {
		p.httpMethod = http.MethodPost
	}
}

//DefaultRegister在初始化的时候就采集了process和go相关指标
func WithDefaultRegistererOpt() M3PusherOpt {
	return func(p *M3Pusher) {
		p.registerer = prometheus.DefaultRegisterer
		//p.gatherers = prometheus.Gatherers{prometheus.DefaultGatherer.(*prometheus.Registry)}
		p.gatherers = prometheus.Gatherers{prometheus.DefaultGatherer}
	}
}

//初始化M3Pusher客户端
func New(url, namespace string, opts ...M3PusherOpt) *M3Pusher {

	var (
		reg = prometheus.NewRegistry()
		err error
	)

	if !strings.Contains(url, "://") {
		url = "http://" + url
	}
	if strings.HasSuffix(url, "/") {
		url = url[:len(url)-1]
	}

	p := &M3Pusher{
		url:        url,
		namespace:  namespace,
		error:      err,
		registerer: reg,
		httpMethod: http.MethodPost,
		grouping:   map[string]string{},
		gatherers:  prometheus.Gatherers{reg},
		client:     &http.Client{},
	}
	for _, opt := range opts {
		opt(p)
	}
	return p
}

//配置HTTP Basic Authentication
func (p *M3Pusher) BasicAuth(username, password string) *M3Pusher {
	//todo 尚未使用到，需对M3进行设置后再使用
	//p.useBasicAuth = true
	//p.username = username
	//p.password = password
	return p
}

//为M3Pusher 设置自定义 HTTP 客户端
func (p *M3Pusher) Client(c HTTPDoer) *M3Pusher {
	p.client = c
	return p
}

//为M3Pusher添加Collector
func (p *M3Pusher) Collector(c prometheus.Collector) *M3Pusher {
	if p.error == nil {
		p.error = p.registerer.Register(c)
	}
	return p
}

//为M3Pusher添加Gatherer
func (p *M3Pusher) Gatherer(g prometheus.Gatherer) *M3Pusher {
	p.gatherers = append(p.gatherers, g)
	return p
}

//将数据推送到M3后仍把数据保存在本地
func (p *M3Pusher) Push() error {
	return p.push()
}

//todo  如何对该接口进行并发测试
func (p *M3Pusher) push() error {
	if p.error != nil {
		return p.error
	}
	mfs, err := p.gatherers.Gather()
	if err != nil {
		return err
	}
	timeSeries, err := transMFtoTS(mfs)
	if err != nil {
		return fmt.Errorf("translateMetricFamilytoWriteRequest error: %v", err)
	}
	buf, _, err := buildWriteRequest(timeSeries, nil, nil)
	if err != nil {
		return fmt.Errorf("buildWriteRequest error: %+v\n", err)
	}
	req, err := http.NewRequest(p.httpMethod, fmt.Sprintf("%s/api/v1/prom/remote/write", p.url), bytes.NewReader(buf))
	if err != nil {
		return fmt.Errorf("http NewRequest error: %v\n", err)
	}
	if p.useBasicAuth {
		req.SetBasicAuth(p.username, p.password)
	}
	resp, err := p.client.Do(req)
	if err != nil {
		return err
	}
	// Depending on version and configuration of the PGW, StatusOK or StatusAccepted may be returned.
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusAccepted {
		body, _ := ioutil.ReadAll(resp.Body) // Ignore any further error as this is for an error message only.
		return fmt.Errorf("unexpected status code %d while pushing to %s: %s", resp.StatusCode, p.url, body)
	}
	return nil
}

//将MetricFamily格式转换为TimeSeries
func transMFtoTS(mfs []*dto.MetricFamily) ([]prompb.TimeSeries, error) {
	timeSeries := []prompb.TimeSeries{}
	for _, mf := range mfs {
		name := mf.GetName()
		//value := getValueFunc(mf)
		ts := returnTimeSeriesFunc(mf)
		//todo 需要确认如何将help提示直接推送到remote storage中,可以抓包结合源码对 push 到pgw进行分析
		//help := mf.GetHelp()
		for _, m := range mf.Metric {
			labels := []prompb.Label{}
			labels = append(labels, prompb.Label{
				Name: "__name__", Value: name,
			})
			for _, label := range m.GetLabel() {
				labels = append(labels, prompb.Label{
					Name: label.GetName(), Value: label.GetValue(),
				})
			}
			//timeSeries = append(timeSeries, prompb.TimeSeries{
			//	Labels:  labels,
			//	Samples: []prompb.Sample{{Value: value(m), Timestamp: time.Now().UnixNano() / 1e6}},
			//})
			timeSeries = append(timeSeries, ts(labels, m, time.Now().UnixNano()/1e6)...)

		}
	}
	return timeSeries, nil
}

type getValue func(metric *dto.Metric) float64

type getTimeSeries func(labels []prompb.Label, metric *dto.Metric, timestamp int64) []prompb.TimeSeries

func returnTimeSeriesFunc(mf *dto.MetricFamily) getTimeSeries {
	switch *mf.Type {
	case dto.MetricType_GAUGE:
		return func(labels []prompb.Label, metric *dto.Metric, timestamp int64) []prompb.TimeSeries {
			gauge := metric.GetGauge()
			if gauge == nil {
				return []prompb.TimeSeries{}
			}
			return []prompb.TimeSeries{{
				Labels:  labels,
				Samples: []prompb.Sample{{Value: gauge.GetValue(), Timestamp: timestamp}},
			}}

		}
	case dto.MetricType_COUNTER:
		return func(labels []prompb.Label, metric *dto.Metric, timestamp int64) []prompb.TimeSeries {
			counter := metric.GetCounter()
			if counter == nil {
				return []prompb.TimeSeries{}
			}
			return []prompb.TimeSeries{{
				Labels:  labels,
				Samples: []prompb.Sample{{Value: counter.GetValue(), Timestamp: timestamp}},
			}}

		}
	//todo 如果是整数，quantile现在显示的是0.00，能否修改掉这个问题
	case dto.MetricType_SUMMARY:
		return func(labels []prompb.Label, metric *dto.Metric, timestamp int64) []prompb.TimeSeries {
			summary := metric.GetSummary()
			if summary == nil {
				return []prompb.TimeSeries{}
			}
			ts := make([]prompb.TimeSeries, 0)

			nameLabel := labels[0]
			//_sum
			ts = append(ts, prompb.TimeSeries{
				Labels: append([]prompb.Label{{
					Name: "__name__", Value: fmt.Sprintf("%s_sum", nameLabel.Value),
				}}, labels[1:]...),
				Samples: []prompb.Sample{
					{Value: summary.GetSampleSum(), Timestamp: timestamp},
				},
			})
			//_count
			ts = append(ts, prompb.TimeSeries{
				Labels: append([]prompb.Label{{
					Name: "__name__", Value: fmt.Sprintf("%s_count", nameLabel.Value),
				}}, labels[1:]...),
				Samples: []prompb.Sample{
					{Value: float64(summary.GetSampleCount()), Timestamp: timestamp},
				},
			})
			//quantile metric
			for _, quantile := range summary.GetQuantile() {
				ts = append(ts, prompb.TimeSeries{
					Labels: append(labels, prompb.Label{
						Name:  "quantile",
						Value: strconv.FormatFloat(quantile.GetQuantile(), 'f', 2, 64),
					}),
					Samples: []prompb.Sample{
						{Value: quantile.GetValue(), Timestamp: timestamp},
					},
				})
			}
			return ts
		}
	case dto.MetricType_HISTOGRAM:
		return func(labels []prompb.Label, metric *dto.Metric, timestamp int64) []prompb.TimeSeries {
			histogram := metric.GetHistogram()
			if histogram == nil {
				return []prompb.TimeSeries{}
			}
			ts := make([]prompb.TimeSeries, 0)
			ts = append(ts, prompb.TimeSeries{
				Labels: labels,
				Samples: []prompb.Sample{
					{Value: histogram.GetSampleSum(), Timestamp: timestamp},
					{Value: float64(histogram.GetSampleCount()), Timestamp: timestamp},
				},
			})
			//sum
			ts = append(ts, prompb.TimeSeries{
				Labels: append([]prompb.Label{{
					Name: "__name__", Value: fmt.Sprintf("%s_sum", labels[0].Value),
				}}, labels[1:]...),
				Samples: []prompb.Sample{
					{Value: histogram.GetSampleSum(), Timestamp: timestamp},
					//{Value: float64(histogram.GetSampleCount()), Timestamp: timestamp},
				},
			})
			//count
			ts = append(ts, prompb.TimeSeries{
				Labels: append([]prompb.Label{{
					Name: "__name__", Value: fmt.Sprintf("%s__count", labels[0].Value),
				}}, labels[1:]...),
				Samples: []prompb.Sample{
					{Value: float64(histogram.GetSampleCount()), Timestamp: timestamp},
				},
			})
			//le
			for _, bucket := range histogram.GetBucket() {
				ts = append(ts, prompb.TimeSeries{
					Labels: append(labels, prompb.Label{
						Name:  "le",
						Value: strconv.FormatFloat(bucket.GetUpperBound(), 'f', 2, 64),
					}),
					Samples: []prompb.Sample{
						{Value: float64(bucket.GetCumulativeCount()), Timestamp: timestamp},
					},
				})
			}
			return ts
		}
	default:
		return func(labels []prompb.Label, metric *dto.Metric, timestamp int64) []prompb.TimeSeries {
			return []prompb.TimeSeries{}
		}
	}
}

//根据当前Collector和Gatherer定时推送数据，推送后数据仍保存在本地
//需要考虑并发的问题，如果并发，会有怎么样的影响。
func PushTimed(p *M3Pusher, t time.Duration) error {
	if p.cancel != nil {
		p.cancel()
	}
	ctx, cancel := context.WithCancel(context.Background())
	ctx = context.WithValue(ctx, "timeDuration", t)
	p.cancel = cancel
	//go p.pushWithContext(ctx)

	go func() {
		ticker := time.NewTicker(t)

		for {
			select {
			case <-ctx.Done():
				break
			case <-ticker.C:
				p.push()
			}
		}
	}()
	return nil
}

//对定时任务做取消
func PushCanceld(p *M3Pusher) {
	if p.cancel == nil {
		return
	}
	p.cancel()
	p.cancel = nil
	return
}
