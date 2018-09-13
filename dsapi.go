// for calling into DarkSky API @ https://darksky.net/dev

package dsapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

const BaseForecastUrl = "https://api.darksky.net/forecast"

type Exclude string

const (
	ExcludeCurrently Exclude = "currently"
	ExcludeMinutely  Exclude = "minutely"
	ExcludeHourly    Exclude = "hourly"
	ExcludeDaily     Exclude = "daily"
	ExcludeAlerts    Exclude = "alerts"
	ExcludeFlags     Exclude = "flags"
)

var client = &http.Client{Timeout: 10 * time.Second}

type Forecast struct {
	Lat       float64  `json:"latitude"`
	Long      float64  `json:"longitude"`
	TimeZone  string   `json:"timezone"`
	Currently Data     `json:"currently"`
	Minutely  Minutely `json:"minutely"`
	Hourly    Hourly   `json:"hourly"`
	Daily     Daily    `json:"daily"`
}

type Data struct {
	Time    int    `json:"time"`
	Summary string `json:"summary"`
	Icon    string `json:"icon"`

	ApparentTemperature         float64 `json:"apparentTemperature"`
	ApparentTemperatureHigh     float64 `json:"apparentTemperatureHigh"`
	ApparentTemperatureHighTime int     `json:"apparentTemperatureHighTime"`
	ApparentTemperatureLow      float64 `json:"apparentTemperatureLow"`
	ApparentTemperatureLowTime  int     `json:"apparentTemperatureLowTime"`
	ApparentTemperatureMax      float64 `json:"apparentTemperatureMax"`
	ApparentTemperatureMaxTime  int     `json:"apparentTemperatureMaxTime"`
	ApparentTemperatureMin      float64 `json:"apparentTemperatureMin"`
	ApparentTemperatureMinTime  int     `json:"apparentTemperatureMinTime"`
	CloudCover                  float64 `json:"cloudCover"`
	DewPoint                    float64 `json:"dewPoint"`
	Humidity                    float64 `json:"humidity"`
	MoonPhase                   float64 `json:"moonPhase"`
	NearestStormBearing         float64 `json:"nearestStormBearing"`
	NearestStormDistance        float64 `json:"nearestStormDistance"`
	Ozone                       float64 `json:"ozone"`
	PrecipIntensity             float64 `json:"precipIntensity"`
	PrecipIntensityMax          float64 `json:"precipIntensityMax"`
	PrecipIntensityMaxTime      int     `json:"precipIntensityMaxTime"`
	PrecipProbability           float64 `json:"precipProbability"`
	PrecipAccumulation          float64 `json:"precipAccumulation"`
	PrecipType                  string  `json:"precipType"`
	Pressure                    float64 `json:"pressure"`
	SunriseTime                 int     `json:"sunriseTime"`
	SunsetTime                  int     `json:"sunsetTime"`
	Temperature                 float64 `json:"temperature"`
	TemperatureHigh             float64 `json:"temperatureHigh"`
	TemperatureHighTime         int     `json:"temperatureHighTime"`
	TemperatureLow              float64 `json:"temperatureLow"`
	TemperatureLowTime          int     `json:"temperatureLowTime"`
	TemperatureMax              float64 `json:"TemperatureMax"`
	TemperatureMaxTime          int     `json:"TemperatureMaxTime"`
	TemperatureMin              float64 `json:"temperatureMin"`
	TemperatureMinTime          int     `json:"temperatureMinTime"`
	UVIndex                     float64 `json:"uvIndex"`
	UVIndexTime                 int     `json:"uvIndexTime"`
	Visibility                  float64 `json:"visibility"`
	WindBearing                 float64 `json:"windBearing"`
	WindGust                    float64 `json:"windGust"`
	WindGustTime                int     `json:"windGustTime"`
	WindSpeed                   float64 `json:"windSpeed"`
}

type Minutely struct {
	Summary string `json:"summary"`
	Icon    string `json:"icon"`
	Data    []Data `json:"data"`
}

type Hourly struct {
	Summary string `json:"summary"`
	Icon    string `json:"icon"`
	Data    []Data `json:"data"`
}

type Daily struct {
	Summary string `json:"summary"`
	Icon    string `json:"icon"`
	Data    []Data `json:"data"`
}

func GetForecast(apiKey string, lat, long float64, excludes []Exclude) (Forecast, error) {
	url := forecastURL(apiKey, lat, long, excludes)
	r, err := client.Get(url)
	if err != nil {
		return Forecast{}, err
	}
	defer r.Body.Close()

	return parseForecastJSON(r.Body)
}

func parseForecastJSON(r io.Reader) (Forecast, error) {
	var resp Forecast
	err := json.NewDecoder(r).Decode(&resp)

	return resp, err
}

func forecastURL(apiKey string, lat, long float64, excludes []Exclude) string {
	ex := ""

	if len(excludes) > 0 {
		var excludeStrings []string

		for _, exclude := range excludes {
			excludeStrings = append(excludeStrings, string(exclude))
		}

		ex = "?exclude=" + strings.Join(excludeStrings, ",")
	}

	return fmt.Sprintf("%s/%s/%.4f,%.4f%s", BaseForecastUrl, apiKey, lat, long, ex)
}
