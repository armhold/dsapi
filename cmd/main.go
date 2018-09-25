package main

import (
	"flag"
	"fmt"
	"github.com/armhold/dsapi"
	"github.com/armhold/dsapi/format"
	"log"
	"os"
	"strconv"
	"strings"
)

// testing out some client code for weather APIs

func usage() {
	fmt.Fprintln(os.Stderr, "Usage: dsapi lat,long")
	fmt.Fprintln(os.Stderr, "E.g.: dsapi 41.48,-81.56")
	os.Exit(1)
}

// TODO: command line arg to specify lat/long arguments instead of zip

func main() {
	flag.Parse()

	apiKey := os.Getenv("DARK_SKY_APIKEY")
	if apiKey == "" {
		log.Fatal("DARK_SKY_APIKEY not set")
	}

	var err error
	var lat, long float64

	if len(os.Args) != 2 {
		usage()
	}

	// TODO: make this more robust, support both comma and spaces between lat/long
	latLong := os.Args[1]

	s := strings.Split(latLong, ",")
	lat, err = strconv.ParseFloat(s[0], 64)
	if err != nil {
		usage()
	}

	long, err = strconv.ParseFloat(s[1], 64)
	if err != nil {
		usage()
	}

	fmt.Printf("Forecast for lat: %f, long: %f:\n\n", lat, long)

	forecast, err := dsapi.GetForecast(apiKey, lat, long, []dsapi.Exclude{})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(format.Print(forecast))
	fmt.Println("Powered by Dark Sky: https://darksky.net/poweredby/")
}
