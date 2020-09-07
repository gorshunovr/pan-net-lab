package main

/*
Given variables listed below, fetches and prints data from
openweathermap.org API using specified format
  CITY_NAME
  OPENWEATHER_API_KEY
*/

import (
	"fmt"
	"log"
	"os"

	"github.com/vascocosta/owm"
)

var (
	OpenweatherAPIKey = os.Getenv("OPENWEATHER_API_KEY")
	CityName          = os.Getenv("CITY_NAME")
)

func main() {
	res := getWeather(OpenweatherAPIKey, CityName, "metric")
	fmt.Println(res)
}

func getWeather(apiKey, city, units string) string {

	// Create a new Client given an API key.
	client := owm.NewClient(apiKey) // new OWM client interface

	// Decode the current weather of a location given the city name and
	// units. WeatherByName returns a Weather and error (nil or other)
	weather, err := client.WeatherByName(city, units)
	// If there is an error, print it and terminate the program.
	if err != nil {
		log.Fatal(err)
		return "ERROR: Could not get weather"
	}

	// Print a string representation of w using the Stringer interface.
	//fmt.Println(&weather)

	return fmt.Sprintf("source=openweathermap, city=\"%v\", description=\"%v\", temp=%v, humidity=%v",
		weather.Name, weather.Weather[0].Description, weather.Main.Temp, weather.Main.Humidity)
}
