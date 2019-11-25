package main

/**
Given variables listed below, fetches and prints data from
openweathermap.org API using specified format
  CITY_NAME
  OPENWEATHER_API_KEY
**/

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

var (
	OPENWEATHER_API_KEY = os.Getenv("OPENWEATHER_API_KEY")
	CITY_NAME           = os.Getenv("CITY_NAME")
)

func main() {
	if OPENWEATHER_API_KEY == "" || CITY_NAME == "" {
		log.Fatalln("ERROR: OPENWEATHER_API_KEY and/or CITY_NAME environment variable missing")
	}
	res := getWeather(OPENWEATHER_API_KEY, CITY_NAME, "metric")
	fmt.Println(res)
}

// Gets JSON from provided URL and returns []byte with JSON blob reply
// Expects ony HTTP 200 code, exits on other response codes
func getJson(url string) []byte {
	res, err := http.Get(url)
	// If there is an error, print it and terminate the program.
	if err != nil {
		log.Fatal("ERROR: sorry, something happened: ", err)
	}

	jsonBlob, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		// If there is an error, print it and terminate the program.
		log.Fatal(err)
	}
	if res.StatusCode != 200 {
		var error Errors
		err := json.Unmarshal(jsonBlob, &error)
		// If there is an error, print it and terminate the program.
		if err != nil {
			log.Fatalln("ERROR: ", err)
		}
		log.Fatalf("ERROR: code: %v, message: %v\n", error.Cod, error.Message)
	}
	return jsonBlob
}

// Accepts API key, city and units strings, returns formatted string
// with weather data or stops on errors
func getWeather(apiKey, city, units string) string {

	OPENWEATHER_API_URL := "https://api.openweathermap.org/data/2.5/weather?q=" +
		CITY_NAME + "&APPID=" + OPENWEATHER_API_KEY + "&units=" + units

	jsonBlob := getJson(OPENWEATHER_API_URL)

	//fmt.Printf("%s\n\n", jsonBlob)

	var weather Weather
	err := json.Unmarshal(jsonBlob, &weather)
	// If there is an error, print it and terminate the program.
	if err != nil {
		log.Fatalln("ERROR: ", err)
	}

	return fmt.Sprintf("source=openweathermap, city=\"%v\", description=\"%v\", temp=%v, humidity=%v",
		weather.Name, weather.Weather[0].Description, weather.Main.Temp, weather.Main.Humidity)
}
