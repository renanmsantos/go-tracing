package webserver

import (
	"fmt"
	"log"
	"net/http"
)

func StartWebserver() {
	http.HandleFunc("/", GetTemperature())

	fmt.Printf("Starting GO-WEATHER at port 8081\n")
	if err := http.ListenAndServe(":8081", nil); err != nil {
		log.Fatal(err)
	}

}
