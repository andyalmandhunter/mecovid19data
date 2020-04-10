package main

import (
	"net/http"
	"net/url"

	"./cache"
	"./transformservice"
)

func handle(r *http.Request) (*transformservice.Data, error) {
	query, err := url.ParseQuery(r.URL.RawQuery)
	if err != nil {
		return nil, err
	}

	data := cache.Get()
	_, latest := query["latest"]
	if latest {
		data = data.FilterLatest()
	}
	if query.Get("county") != "" {
		data = data.FilterCounty(query.Get("county"))
	}

	return &data, nil
}
