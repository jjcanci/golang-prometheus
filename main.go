package main

import (
	"log"
	"math/rand"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/push"
)

var (
	completionTime = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "db_backup_last_completion_timestamp_seconds",
		Help: "The timestamp of the last completion of a DB backup, successful or not.",
	})
	successTime = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "db_backup_last_success_timestamp_seconds",
		Help: "The timestamp of the last successful completion of a DB backup.",
	})
	duration = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "db_backup_duration_seconds",
		Help: "The duration of the last DB backup in seconds.",
	})
	records = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "db_backup_records_processed",
		Help: "The number of records processed in the last DB backup.",
	})
)

func performBackup() (int, error) {
	// Perform the backup and return the number of backed up records and any
	// applicable error.
	// ...
	return 42, nil
}

func main() {
	log.Println("Start")
	rand.Seed(time.Now().UTC().UnixNano())

	sales := prometheus.NewCounter(prometheus.CounterOpts{
		Name: "sales_promo_tradein",
		Help: "Sales promo TRADEIN",
	})

	for index := 0; index < 1000; index++ {
		sales.Add(float64(rand.Intn(10)))

		if err := push.New("http://localhost:9091", "job_sales").
			Collector(sales).
			Grouping("channel", "18").
			Grouping("sourceSystem", "ipos").
			Push(); err != nil {
			log.Fatalln("Could not push completion time to Pushgateway:", err)
		}
		log.Printf("%v) Pushed!!\n", index)

		time.Sleep(time.Duration(rand.Intn(10)) * time.Second)
	}

	log.Println("End")
}
