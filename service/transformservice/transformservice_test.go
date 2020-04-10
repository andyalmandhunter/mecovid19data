package transformservice

import (
	"github.com/google/go-cmp/cmp"
	"testing"
)

func TestConvert(t *testing.T) {
	tests := map[string]struct {
		dataDict Dates
		want     Data
	}{
		"happy path": {
			Dates{
				"2020-04-08": {
					"Kennebec": {3, 2, 1, 0},
					"Waldo":    {6, 4, 2, 0},
				},
				"2020-04-09": {
					"Kennebec": {3, 2, 1, 0},
				},
			},
			Data{
				[]Date{
					{
						"2020-04-08",
						[]County{
							{"Kennebec", 3, 2, 1, 0},
							{"Waldo", 6, 4, 2, 0},
						},
					},
					{
						"2020-04-09",
						[]County{
							{"Kennebec", 3, 2, 1, 0},
						},
					},
				},
			},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := Convert(tc.dataDict)
			diff := cmp.Diff(tc.want, got)
			if diff != "" {
				t.Fatalf(diff)
			}
		})
	}
}
