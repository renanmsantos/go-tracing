package webserver

import (
	"fmt"
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.opentelemetry.io/otel/trace"
)

func StartWebserver(tracer trace.Tracer) {
	http.Handle("/metrics", promhttp.Handler())
	http.HandleFunc("/", GetTemperature(tracer))

	fmt.Printf("Starting GO-WEATHER at port 8081\n")
	if err := http.ListenAndServe(":8081", nil); err != nil {
		log.Fatal(err)
	}

}
