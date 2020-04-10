package transformservice

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"sort"
	"strconv"
	"strings"
	"time"

	"mecovid19data/service/dataservice"
)

type Data struct {
	Dates []Date `json:"dates"`
}

type Date struct {
	Date     string   `json:"date"`
	Counties []County `json:"counties"`
}

type County struct {
	County           string `json:"county"`
	Confirmed        int    `json:"confirmed"`
	Recovered        int    `json:"recovered"`
	Hospitalizations int    `json:"hospitalizations"`
	Deaths           int    `json:"deaths"`
}

func Parse(data []dataservice.Record) *Data {
	dataDict := ParseToDict(data, time.Now())
	out := Convert(dataDict)
	return &out
}

func Convert(dataDict Dates) Data {
	var keys []string
	for k := range dataDict {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	var out Data
	for _, k := range keys {
		out.Dates = append(out.Dates, convertDate(k, dataDict[k]))
	}

	return out
}

func convertDate(date string, dataDict Counties) Date {
	var keys []string
	for k := range dataDict {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	var out Date
	out.Date = date
	for _, k := range keys {
		out.Counties = append(out.Counties, County{
			k,
			dataDict[k].Confirmed,
			dataDict[k].Recovered,
			dataDict[k].Hospitalizations,
			dataDict[k].Deaths,
		})
	}

	return out
}

func (data *Data) FilterCounty(c string) Data {
	countyStr := strings.ToLower(c)

	var filteredData Data
	filteredData.Dates = make([]Date, 0, len(data.Dates))

	for _, date := range data.Dates {
		filteredDate := Date{Date: date.Date}
		for _, county := range date.Counties {
			if strings.ToLower(county.County) == countyStr {
				filteredDate.Counties = append(filteredDate.Counties, county)
			}
		}

		if len(filteredDate.Counties) > 0 {
			filteredData.Dates = append(filteredData.Dates, filteredDate)
		}
	}

	return filteredData
}

func (data *Data) FilterLatest() Data {
	var filteredData Data
	if len(data.Dates) > 0 {
		filteredData.Dates = []Date{data.Dates[len(data.Dates)-1]}
	} else {
		filteredData.Dates = make([]Date, 0)
	}
	return filteredData
}

func (data *Data) WriteCsv(w io.Writer) error {
	csv := csv.NewWriter(w)

	csv.Write([]string{
		"date",
		"county",
		"confirmed",
		"recovered",
		"hospitalizations",
		"deaths",
	})

	for _, date := range data.Dates {
		for _, county := range date.Counties {
			if err := csv.Write([]string{
				date.Date,
				county.County,
				strconv.Itoa(county.Confirmed),
				strconv.Itoa(county.Recovered),
				strconv.Itoa(county.Hospitalizations),
				strconv.Itoa(county.Deaths),
			}); err != nil {
				return err
			}
		}
	}

	csv.Flush()
	return csv.Error()
}

func (data *Data) WriteJson(w io.Writer) error {
	res, err := json.Marshal(data)
	if err != nil {
		return err
	}

	fmt.Fprintf(w, "%s\n", res)
	return nil
}
