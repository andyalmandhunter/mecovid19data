package transformservice

import (
	"time"

	"mecovid19data/service/dataservice"
)

type Dates map[string]Counties

type Counties map[string]Record

type Record struct {
	Confirmed        int
	Recovered        int
	Hospitalizations int
	Deaths           int
}

func ParseToDict(data []dataservice.Record, now time.Time) Dates {
	parsed := make(Dates)

	// If we have any new data from today (latest journal entry starts
	// today), consider the current day valid and move 'now' to
	// tomorrow.
	if len(data) > 0 && sameLocalDay(data[len(data)-1].Start, now) {
		now = now.AddDate(0, 0, 1)
	}

	for _, d := range data {
		end := d.End
		if end.IsZero() {
			end = now
		}

		for _, dateStr := range GetValidDates(d.Start, end) {
			_, ok := parsed[dateStr]
			if !ok {
				parsed[dateStr] = make(Counties)
			}
			parsed[dateStr][d.County] = Record{
				d.Confirmed,
				d.Recovered,
				d.Hospitalizations,
				d.Deaths,
			}
		}
	}

	return parsed
}

func sameLocalDay(a time.Time, b time.Time) bool {
	aa := a.In(maine)
	bb := b.In(maine)

	return aa.Year() == bb.Year() && aa.Month() == bb.Month() && aa.Day() == bb.Day()
}

// GetValidDates computes the dates for which a particular journal
// entry is valid.
func GetValidDates(start time.Time, end time.Time) []string {
	dates := make([]string, 0)

	s := start.In(maine)

	// Begin with the midnight following the start time
	d := time.Date(s.Year(), s.Month(), s.Day()+1, 0, 0, 0, 0, maine)
	for {
		if d.After(end) {
			break
		}

		// Since we're iterating over midnights, store the date of the
		// day preceding midnight
		dates = append(dates, d.AddDate(0, 0, -1).Format("2006-01-02"))
		d = d.AddDate(0, 0, 1)
	}

	return dates
}
