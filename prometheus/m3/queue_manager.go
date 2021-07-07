package m3

import (
	"github.com/gogo/protobuf/proto"
	"github.com/golang/snappy"
	"github.com/prometheus/client_golang/prometheus/m3/prompb"
)

func buildWriteRequest(samples []prompb.TimeSeries, metadata []prompb.MetricMetadata, buf []byte) ([]byte, int64, error) {
	var highest int64
	for _, ts := range samples {
		// At the moment we only ever append a TimeSeries with a single sample or exemplar in it.
		if len(ts.Samples) > 0 && ts.Samples[0].Timestamp > highest {
			highest = ts.Samples[0].Timestamp
		}
		if len(ts.Exemplars) > 0 && ts.Exemplars[0].Timestamp > highest {
			highest = ts.Exemplars[0].Timestamp
		}
	}

	req := &prompb.WriteRequest{
		Timeseries: samples,
		Metadata:   metadata,
	}

	data, err := proto.Marshal(req)
	if err != nil {
		return nil, highest, err
	}

	// snappy uses len() to see if it needs to allocate a new slice. Make the
	// buffer as long as possible.
	if buf != nil {
		buf = buf[0:cap(buf)]
	}
	compressed := snappy.Encode(buf, data)
	return compressed, highest, nil
}
