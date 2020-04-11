package cache

import (
	"log"

	"mecovid19data/service/transformservice"
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
}

type request struct {
	Type requestType
	C    chan transformservice.Data
}

func Run() {
	log.Println("Starting cache")

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

func Set(data transformservice.Data) {
	c := make(chan transformservice.Data)
	requestChannel <- request{set, c}
	c <- data
}

func Get() transformservice.Data {
	c := make(chan transformservice.Data)
	requestChannel <- request{get, c}
	return <-c
}
