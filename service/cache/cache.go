package cache

import (
	"log"
	"time"

	"../config"
	"../dataservice"
	"../transformservice"
)

type requestType uint

const (
	get requestType = iota
	set
)

var requestChannel chan request
var value transformservice.Data

func init() {
	requestChannel = make(chan request)
	value.Dates = make([]transformservice.Date, 0)
	go run()
	go poll()
}

type request struct {
	Type requestType
	C    chan transformservice.Data
}

func run() {
	for {
		request := <-requestChannel
		switch request.Type {
		case get:
			request.C <- value
		case set:
			value = <-request.C
		}
	}
}

func poll() {
	log.Printf("Refreshing cache every %v", config.PollingPeriod)

	for {
		rawData, err := dataservice.Get()
		if err != nil {
			log.Fatal(err)
		}

		data := transformservice.Parse(rawData)

		c := make(chan transformservice.Data)
		requestChannel <- request{set, c}
		c <- *data

		time.Sleep(config.PollingPeriod)
	}
}

func Get() transformservice.Data {
	c := make(chan transformservice.Data)
	requestChannel <- request{get, c}
	return <-c
}
