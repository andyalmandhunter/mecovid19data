package refreshservice

import (
	"log"
	"time"

	"mecovid19data/service/cache"
	"mecovid19data/service/dataservice"
	"mecovid19data/service/transformservice"
)

type RefreshService struct {
	Period time.Duration
}

func New(d time.Duration) RefreshService {
	return RefreshService{d}
}

func (r RefreshService) Run() {
	log.Printf("Refreshing cache every %v\n", r.Period)

	for {
		rawData, err := dataservice.Get()
		if err != nil {
			log.Fatal(err)
		}

		cache.Set(*transformservice.Parse(rawData))
		time.Sleep(r.Period)
	}
}
