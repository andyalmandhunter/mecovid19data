package transformservice

import (
	"github.com/google/go-cmp/cmp"
	"testing"
	"time"

	"mecovid19data/service/dataservice"
)

func TestParseToDict(t *testing.T) {
	tests := map[string]struct {
		data []dataservice.Record
		now  time.Time
		want Dates
	}{
		"happy path": {
			[]dataservice.Record{
				{
					"Kennebec",
					3,
					2,
					1,
					0,
					time.Date(2020, time.April, 8, 12, 0, 0, 0, time.UTC),
					time.Date(2020, time.April, 10, 12, 0, 0, 0, time.UTC),
				},
			},
			time.Date(2020, time.April, 11, 12, 0, 0, 0, time.UTC),
			Dates{
				"2020-04-08": {
					"Kennebec": {3, 2, 1, 0},
				},
				"2020-04-09": {
					"Kennebec": {3, 2, 1, 0},
				},
			},
		},
		"with null End": {
			[]dataservice.Record{
				{
					"Kennebec",
					3,
					2,
					1,
					0,
					time.Date(2020, time.April, 8, 12, 0, 0, 0, time.UTC),
					time.Time{},
				},
			},
			time.Date(2020, time.April, 9, 12, 0, 0, 0, time.UTC),
			Dates{
				"2020-04-08": {
					"Kennebec": {3, 2, 1, 0},
				},
			},
		},
		"with null End on same day": {
			[]dataservice.Record{
				{
					"Kennebec",
					3,
					2,
					1,
					0,
					time.Date(2020, time.April, 8, 12, 0, 0, 0, time.UTC),
					time.Time{},
				},
			},
			time.Date(2020, time.April, 8, 13, 0, 0, 0, time.UTC),
			Dates{
				"2020-04-08": {
					"Kennebec": {3, 2, 1, 0},
				},
			},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := ParseToDict(tc.data, tc.now)
			diff := cmp.Diff(tc.want, got)
			if diff != "" {
				t.Fatalf(diff)
			}
		})
	}
}

func TestGetValidDates(t *testing.T) {
	tests := map[string]struct {
		start time.Time
		end   time.Time
		want  []string
	}{
		"same day": {
			time.Date(2020, time.April, 8, 12, 0, 0, 0, time.UTC),
			time.Date(2020, time.April, 8, 13, 0, 0, 0, time.UTC),
			[]string{},
		},
		"one day difference": {
			time.Date(2020, time.April, 8, 12, 0, 0, 0, time.UTC),
			time.Date(2020, time.April, 9, 12, 0, 0, 0, time.UTC),
			[]string{"2020-04-08"},
		},
		"two days difference": {
			time.Date(2020, time.April, 8, 12, 0, 0, 0, time.UTC),
			time.Date(2020, time.April, 10, 12, 0, 0, 0, time.UTC),
			[]string{"2020-04-08", "2020-04-09"},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := GetValidDates(tc.start, tc.end)
			diff := cmp.Diff(tc.want, got)
			if diff != "" {
				t.Fatalf(diff)
			}
		})
	}
}
