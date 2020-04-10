package main

import (
	"fmt"
	"net/http"
)

func jsonContent(
	handle func(http.ResponseWriter, *http.Request) error,
) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		if err := handle(w, r); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "{\"message\":\"%s\"}\n", err)
		}
	}
}

func csvContent(
	handle func(http.ResponseWriter, *http.Request) error,
) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/csv")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		if err := handle(w, r); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "%s\n", err)
		}
	}
}
