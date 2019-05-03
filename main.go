package main

import (
	"log"
	"math/rand"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/push"
)

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
