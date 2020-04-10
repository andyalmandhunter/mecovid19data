package transformservice

import (
	"log"
	"time"
)

var maine *time.Location

func init() {
	var err error
	maine, err = time.LoadLocation("America/New_York")
	if err != nil {
		log.Fatal(err)
	}
}
