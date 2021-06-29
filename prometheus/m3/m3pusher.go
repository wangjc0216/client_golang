package m3

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	dto "github.com/prometheus/client_model/go"
	"github.com/prometheus/prometheus/prompb"

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

	ctx context.Context
}

//初始化选项，包括本地缓存空间设置
type M3PusherOpt func(*M3Pusher)

func WithHttpPutOpt() M3PusherOpt {
	return func(p *M3Pusher) {
		p.httpMethod = "PUT"
	}
}

func WithHttpPostOpt() M3PusherOpt {
	return func(p *M3Pusher) {
		p.httpMethod = "POST"
	}
}

//初始化M3Pusher客户端
func New(url, namespace string, opts ...M3PusherOpt) *M3Pusher {

	var (
		reg = prometheus.NewRegistry()
		err error
	)

	p := &M3Pusher{
		url:        url,
		namespace:  namespace,
		error:      err,
		registerer: reg,
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
	p.useBasicAuth = true
	p.username = username
	p.password = password
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

//将数据推送到M3后清空本地缓存
//func (p *M3Pusher) PushWithEmpty() error { return nil }

//将数据推送到M3后仍把数据保存在本地
func (p *M3Pusher) Push() error {
	return p.push()
}

func (p *M3Pusher) push() error {
	if p.error != nil {
		return p.error
	}
	mfs, err := p.gatherers.Gather()
	if err != nil {
		return err
	}
	timeSeries, err := translateMetricFamilytoWriteRequest(mfs)
	if err != nil {
		return fmt.Errorf("translateMetricFamilytoWriteRequest error: %v", err)
	}
	buf, _, err := buildWriteRequest(timeSeries, nil, nil)
	if err != nil {
		fmt.Printf("buildWriteRequest error: %+v\n", err)
	}
	//todo 对url需要进行检查并重新构造
	req, err := http.NewRequest(p.httpMethod, fmt.Sprintf("%s/api/v1/prom/remote/write", p.url), bytes.NewReader(buf))
	if err != nil {
		fmt.Printf("http NewRequest error: %v\n", err)
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

func translateMetricFamilytoWriteRequest(mfs []*dto.MetricFamily) ([]prompb.TimeSeries, error) {
	timeSeries := []prompb.TimeSeries{}
	for _, mf := range mfs {
		name := mf.GetName()
		value := getValueFunc(mf)
		//todo 需要确认如何将help提示直接推送到remote storage中
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
			timeSeries = append(timeSeries, prompb.TimeSeries{
				Labels:  labels,
				Samples: []prompb.Sample{{Value: value(m), Timestamp: time.Now().UnixNano() / 1e6}},
			})
		}
	}
	return timeSeries, nil
}

type getValue func(metric *dto.Metric) float64

func getValueFunc(mf *dto.MetricFamily) getValue {
	switch *mf.Type {
	case dto.MetricType_COUNTER:
		return func(metric *dto.Metric) float64 {
			return *metric.Counter.Value
		}
	case dto.MetricType_GAUGE:
		return func(metric *dto.Metric) float64 {
			return *metric.Gauge.Value
		}
	//todo  当前只支持对Gauge和Counter的指标写入
	//case dto.MetricType_SUMMARY:
	//	getValue = func(metric *dto.Metric) float64 {
	//		return *metric.Summary.SampleSum
	//	}
	//case dto.MetricType_UNTYPED:
	//	getValue = func(metric *dto.Metric) float64 {
	//		return *metric.Untyped.Value
	//	}
	//case dto.MetricType_HISTOGRAM:
	//	getValue = func(metric *dto.Metric) float64 {
	//		return *metric.Histogram.SampleSum
	//	}
	default:
		return func(*dto.Metric) float64 {
			return 0
		}
	}

}

//根据当前Collector和Gatherer定时推送数据，推送后数据仍保存在本地
//todo 需要对goroutine做管理，使用context做管理，可以参考毛大的Context
func (p *M3Pusher) PushTimed(t time.Duration) {
	p.ctx = context.Background()

	timeTicker := time.NewTicker(t)
	select {
	case <-timeTicker.C:
		p.push()
	}

}
func Go(f func()) {
	go f()
}

//对定时任务做取消
func (p *M3Pusher) PushCanceld() {

}

//根据当前Collector和Gatherer定时推送数据，推送后数据清空
func (p *M3Pusher) PushTimedWithEmpty(t *time.Duration) error { return nil }
