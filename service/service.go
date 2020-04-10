package main

import (
	"log"
	"net/http"
)

const port = "8080"

func main() {
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
