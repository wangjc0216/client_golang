package m3

import (
	"bytes"
	"fmt"
	"net/http"
	"time"

	"github.com/prometheus/prometheus/prompb"
)

var writeRequestFixture = &prompb.WriteRequest{
	Timeseries: []prompb.TimeSeries{
		{
			Labels: []prompb.Label{
				{Name: "__name__", Value: "test_metric2"},
				{Name: "b", Value: "c"},
				{Name: "baz", Value: "qux"},
				{Name: "d", Value: "e"},
				{Name: "foo", Value: "bar"},
			},
			Samples: []prompb.Sample{{Value: 1, Timestamp: time.Now().UnixNano() / 1e6}},
		},
		{
			Labels: []prompb.Label{
				{Name: "__name__", Value: "test_metric3"},
				{Name: "b", Value: "c"},
				{Name: "baz", Value: "qux"},
				{Name: "d", Value: "e"},
				{Name: "foo", Value: "bar"},
			},
			Samples: []prompb.Sample{{Value: 2, Timestamp: time.Now().UnixNano() / 1e6}},
		},
	},
}

func ExampleRemoteWrite() {
	buf, _, err := buildWriteRequest(writeRequestFixture.Timeseries, nil, nil)
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
