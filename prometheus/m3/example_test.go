package m3

import (
	"bytes"
	"fmt"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"

	"github.com/prometheus/prometheus/prompb"
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

var metadata = &prompb.MetricMetadata{
	Type:             prompb.MetricMetadata_GAUGE,
	MetricFamilyName: "test_metric1",
	Help:             "this is test_metric1",
	Unit:             "unit_num",
}

//远程写入
func ExampleRemoteWrite() {
	buf, _, err := buildWriteRequest(writeRequestFixture.Timeseries, []prompb.MetricMetadata{*metadata, *metadata}, nil)
	if err != nil {
		fmt.Printf("buildWriteRequest error: %+v\n", err)
	}

	req, err := http.NewRequest(http.MethodPost, "http://localhost:7201/api/v1/prom/remote/write", bytes.NewReader(buf))
	if err != nil {
		fmt.Printf("http NewRequest error: %v\n", err)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Printf("http defaultClient  do error: %v\n", err)
	}
	fmt.Printf("resp: %v\n", resp)
	//output:
}

//直接对M3集群进行推送数据
func ExamplePushDate() {
	completionTime := prometheus.NewGauge(prometheus.GaugeOpts{
		Name:        "db_backup_last_completion_timestamp_seconds",
		Help:        "The timestamp of the last successful completion of a DB backup.",
		ConstLabels: map[string]string{"company_name": "armani", "company_id": "u39nd0s"},
	})
	completionTime.Set(10086)

	pusher := New("http://localhost:7201", "default", WithHttpPostOpt())
	pusher = pusher.Collector(completionTime)
	err := pusher.Push()
	if err != nil {
		fmt.Printf("pusher push error: %v\n", err)
	}
	//output:
}
