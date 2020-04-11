package main

import (
	"log"
	"net/http"

	"mecovid19data/service/config"
	"mecovid19data/service/refreshservice"
)

const port = "8080"

func main() {
	rs := refreshservice.New(config.PollingPeriod)
	go rs.Run()

	http.HandleFunc("/api/v0/countydata.json", jsonContent(jsonHandler))
	http.HandleFunc("/api/v0/countydata.csv", csvContent(csvHandler))

	log.Printf("Starting server on localhost:%s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func jsonHandler(w http.ResponseWriter, r *http.Request) error {
	data, err := handle(r)
	if err != nil {
		return err
	}

	return data.WriteJson(w)
}

func csvHandler(w http.ResponseWriter, r *http.Request) error {
	data, err := handle(r)
	if err != nil {
		return err
	}
	return data.WriteCsv(w)
}
