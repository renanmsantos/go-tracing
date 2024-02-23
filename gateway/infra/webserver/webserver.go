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
	http.HandleFunc("/", ValidationHandler(tracer))

	fmt.Printf("Starting GO-GATEWAY at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}

}
