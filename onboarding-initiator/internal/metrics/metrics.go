package metrics

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (

	RequestCount = prometheus.NewCounter(

		prometheus.CounterOpts{

			Name: "onboarding_requests_total",

			Help: "Total onboarding API requests",
		})

	KafkaErrors = prometheus.NewCounter(

		prometheus.CounterOpts{

			Name: "kafka_errors_total",

			Help: "Total kafka publish errors",
		})
)

func Register() {

	prometheus.MustRegister(RequestCount)

	prometheus.MustRegister(KafkaErrors)
}

func Handler() http.Handler {

	return promhttp.Handler()
}
