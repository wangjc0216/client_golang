package m3

import (
	"bytes"
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"

	"github.com/prometheus/client_golang/prometheus/m3/prompb"
)

var writeRequestFixture = &prompb.WriteRequest{
	Timeseries: []prompb.TimeSeries{
		{
			Labels: []prompb.Label{
				{Name: "__name__", Value: "test_metric1"},
				{Name: "b", Value: "c"},
				{Name: "baz", Value: "qux"},
				{Name: "d", Value: "e"},
				{Name: "foo", Value: "bar"},
			},
			Samples: []prompb.Sample{{Value: 1, Timestamp: time.Now().UnixNano() / 1e6}},
		},
		{
			Labels: []prompb.Label{
				{Name: "__name__", Value: "test_metric1"},
				{Name: "b", Value: "c"},
				{Name: "baz", Value: "quax"},
				{Name: "d", Value: "ed"},
				{Name: "foo", Value: "bar"},
			},
			Samples: []prompb.Sample{{Value: 2, Timestamp: time.Now().UnixNano() / 1e6}},
		},
	},
}

var metadata = prompb.MetricMetadata{
	Type:             prompb.MetricMetadata_COUNTER,
	MetricFamilyName: "prometheus_remote_storage_sent_metadata_bytes_total",
	Help:             "a nice help text",
	Unit:             "",
}

//远程写入Sample
func ExampleRemoteWriteSample() {
	//buf, _, err := buildWriteRequest(writeRequestFixture.Timeseries, []prompb.MetricMetadata{*metadata, *metadata}, nil)
	buf, _, err := buildWriteRequest(writeRequestFixture.Timeseries, nil, nil)

	if err != nil {
		fmt.Printf("buildWriteRequest error: %+v\n", err)
	}

	req, err := http.NewRequest(http.MethodPost, "http://localhost:7201/api/v1/prom/remote/write", bytes.NewReader(buf))
	if err != nil {
		fmt.Printf("http NewRequest error: %v\n", err)
	}
	_, err = http.DefaultClient.Do(req)
	if err != nil {
		fmt.Printf("http defaultClient  do error: %v\n", err)
	}
	//output:
}

//todo 元数据如何推送，还是在思考
func ExampleRemoteWriteMetadata() {
	buf, _, err := buildWriteRequest(nil, []prompb.MetricMetadata{metadata}, nil)
	if err != nil {
		fmt.Printf("buildWriteRequest error: %+v\n", err)
	}

	req, err := http.NewRequest(http.MethodPost, "http://localhost:7201/api/v1/prom/remote/write", bytes.NewReader(buf))
	if err != nil {
		fmt.Printf("http NewRequest error: %v\n", err)
	}
	_, err = http.DefaultClient.Do(req)
	if err != nil {
		fmt.Printf("http defaultClient  do error: %v\n", err)
	}
	//output:
}

//直接对M3集群进行推送数据
func ExamplePushDate() {
	devgroupCount := prometheus.NewGauge(prometheus.GaugeOpts{
		Name:        "wop_monitor_all_groups_cnt",
		Help:        "count of  shop devgroup",
		ConstLabels: map[string]string{"shop_id": "4324870024005277952", "company_id": "00a38173-5e7d-40f0-b198-16f3cd5e1b53"},
	})
	devgroupCount.Set(10086)

	pusher := New("http://localhost:7201", "default", WithHttpPostOpt())
	pusher = pusher.Collector(devgroupCount).Gatherer(prometheus.DefaultGatherer)
	err := pusher.Push()
	if err != nil {
		fmt.Printf("pusher push error: %v\n", err)
	}
	//output:
}

//定时推送
func ExamplePushTimed() {
	rand.Seed(time.Now().UnixNano())

	devgroupCount := prometheus.NewGauge(prometheus.GaugeOpts{
		Name:        "wop_monitor_all_groups_cnt",
		Help:        "count of  shop devgroup",
		ConstLabels: map[string]string{"shop_id": "4324870024005277952", "company_id": "00a38173-5e7d-40f0-b198-16f3cd5e1b53"},
	})
	devgroupCount.Set(5)
	//使用WithDefaultRegistererOpt选项，可以在推送监控指标的同时加上go和process相关的监控指标
	pusher := New("localhost:7201", "default", WithHttpPostOpt(), WithDefaultRegistererOpt())
	//pusher := New("localhost:7201", "default", WithHttpPostOpt())

	pusher = pusher.Collector(devgroupCount)
	err := PushTimed(pusher, time.Second*1)
	if err != nil {
		fmt.Printf("pusher pushTimed error: %v\n", err)
	}
	for i := 0; i < 20; i++ {
		time.Sleep(time.Second)
		devgroupCount.Set(float64(rand.Intn(100)))
	}
	PushCanceld(pusher)

	time.Sleep(5 * time.Second)
	//output:
}
