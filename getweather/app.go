package main

import (
	"fmt"
	"github.com/vascocosta/owm"
	"log"
	"os"
)

var (
	OPENWEATHER_API_KEY = os.Getenv("OPENWEATHER_API_KEY")
	CITY_NAME           = os.Getenv("CITY_NAME")
)

func main() {
	res := getWeather(OPENWEATHER_API_KEY, CITY_NAME, "metric")
	fmt.Println(res)
}

func getWeather(api_key, city, units string) string {

	// Create a new Client given an API key.
	client := owm.NewClient(api_key) // new OWM client interface

	// Decode the current weather of a location given the city name and
	// units. WeatherByName returns a Weather.
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
