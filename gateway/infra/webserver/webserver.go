package webserver

import (
	"fmt"
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func StartWebserver() {
	http.Handle("/metrics", promhttp.Handler())
	http.HandleFunc("/", ValidationHandler())

	fmt.Printf("Starting GO-GATEWAY at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}

}
