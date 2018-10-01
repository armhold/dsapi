package format

import (
	"encoding/json"
	"github.com/armhold/dsapi"
	"github.com/google/go-cmp/cmp"
	"os"
	"testing"
)

func TestPrint(t *testing.T) {
	expected :=
		`-------------------------------------------------------
|   70|   71|   73|   75|   75|   75|   74|   73|   71|
|   ⛅|   ⛅|   ⛅|   ☀️ |   ☀️ |   ☀️ |   ☀️ |   ☀️ |   🌒|
|  Now| 12pm|  1pm|  2pm|  3pm|  4pm|  5pm|  6pm|  7pm|
-------------------------------------------------------
 Mostly Cloudy

 Feels like:        70°  |   Sunrise:         6:25am
 Precipitation:      5%  |   Sunset:          6:55pm
 Humidity:          85%  |   Wind:            1mph
 Cloud Cover:       67%  |   UV Index:        5

-------------------------------------------------------
Mostly cloudy for the hour.
`

	r, err := os.Open("../sample_forecast.json")
	if err != nil {
		t.Fatal(err)
	}

	var resp dsapi.Forecast
	err = json.NewDecoder(r).Decode(&resp)
	if err != nil {
		t.Fatal(err)
	}

	actual := Print(resp)

	if expected != actual {
		t.Errorf(cmp.Diff(expected, actual))
	}
}
