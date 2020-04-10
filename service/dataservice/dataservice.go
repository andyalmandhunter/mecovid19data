package dataservice

import (
	"encoding/csv"
	"io"
	"strconv"
	"time"

	"mecovid19data/service/s3client"
)

func Get() ([]Record, error) {
	b, err := s3client.Get()
	if err != nil {
		return nil, err
	}

	var records []Record

	csv := csv.NewReader(b)

	// Ignore header row
	_, err = csv.Read()
	if err == io.EOF {
		return records, nil
	}
	if err != nil {
		return nil, err
	}

	for {
		row, err := csv.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}

		rec, err := NewFromCsvRow(row)
		if err != nil {
			return nil, err
		}

		records = append(records, *rec)
	}

	return records, nil
}

type Record struct {
	County           string
	Confirmed        int
	Recovered        int
	Hospitalizations int
	Deaths           int
	Start            time.Time
	End              time.Time
}

func NewFromCsvRow(row []string) (*Record, error) {
	confirmed, err := strconv.Atoi(row[1])
	if err != nil {
		return nil, err
	}

	recovered, err := strconv.Atoi(row[2])
	if err != nil {
		return nil, err
	}

	hospitalizations, err := strconv.Atoi(row[3])
	if err != nil {
		return nil, err
	}

	deaths, err := strconv.Atoi(row[4])
	if err != nil {
		return nil, err
	}

	start, err := parseTime(row[5])
	if err != nil {
		return nil, err
	}

	end, err := parseTime(row[6])
	if err != nil {
		return nil, err
	}

	return &Record{
		row[0],
		confirmed,
		recovered,
		hospitalizations,
		deaths,
		start,
		end,
	}, nil
}

func parseTime(timeStr string) (time.Time, error) {
	if timeStr == "" {
		return time.Time{}, nil
	}

	return time.Parse("2006-01-02T15:04:05.000Z", timeStr)
}
