package main

import (
	"fmt"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/push"
)

func main() {
	completionTime := prometheus.NewGauge(prometheus.GaugeOpts{
		Name:        "db_backup_last_completion_timestamp_seconds",
		Help:        "The timestamp of the last successful completion of a DB backup.",
		ConstLabels: map[string]string{"company_name": "armani", "company_id": "u39nd0s"},
	})
	completionTime.Set(1)

	xibeiDevGroupStatus := prometheus.NewGauge(prometheus.GaugeOpts{
		Name:        "devgroup_armani_status",
		ConstLabels: map[string]string{"company_name": "xibei2", "company_id": "0oh6mmdn5"},
		Help:        "company devgroup living status.",
	})
	xibeiDevGroupStatus.Set(0)

	pusher := push.New("http://localhost:9091", "db_backup").
		//Collector(completionTime).Collector(xibeiDevGroupStatus).
		//Collector(xibeiDevGroupStatus).
		Collector(completionTime).
		Grouping("db", "customers")

	if err := pusher.
		//Push(); err != nil {
		Add(); err != nil {
		fmt.Println("Could not push completion time to Pushgateway:", err)
	}

	time.Sleep(time.Second * 20)

	completionTime.Set(100)

	err := pusher.Add()
	if err != nil {
		fmt.Println("Could not push completion time to Pushgateway:", err)
	}

}
