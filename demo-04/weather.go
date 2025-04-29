package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"time"
)

type Weather struct {
	Location Location `json:"location"`
	Current  Current  `json:"current"`
}

type Location struct {
	Name           string  `json:"name"`
	Region         string  `json:"region"`
	Country        string  `json:"country"`
	Lat            float32 `json:"lat"`
	Lon            float32 `json:"lon"`
	Timezone       string  `json:"tz_id"`
	LocaltimeEpoch int     `json:"localtime_epoch"`
	Localtime      string  `json:"localtime"`
}

type Current struct {
	Temp      float32   `json:"temp_c"`
	Condition Condition `json:"condition"`
}

type Condition struct {
	Text string `json:"text"`
}

const API_KEY = "cd99b2c687834ac0be1210113233005"
const API_URL = "http://api.weatherapi.com/v1"

func main() {
	// Build the request url
	location := url.QueryEscape("Rogil, Portugal")
	url := fmt.Sprintf("%s/current.json?key=%s&q=%s", API_URL, API_KEY, location)
	time.Now()

	// Makes the HTTP request
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	resp.Body.Close()
	if resp.StatusCode != 200 {
		log.Fatalf("Response failed with status code: %d and body: %s", resp.StatusCode, body)
	}
	if err != nil {
		log.Fatal(err)
	}

	// Parses the response json into the Weather struct
	var weather Weather
	err = json.Unmarshal(body, &weather)
	if err != nil {
		log.Fatal(err)
	}

	// Prints the weather condition and temperature for the location
	fmt.Printf("\n%s, %s\n", weather.Location.Name, weather.Location.Country)
	fmt.Printf("%s, %.0fºC\n\n", weather.Current.Condition.Text, weather.Current.Temp)

	fmt.Printf("Press \"Enter\" to exit.")
	fmt.Scanln()
}
