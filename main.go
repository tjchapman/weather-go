package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/fatih/color"
	"github.com/joho/godotenv"
)


type Weather struct {
	Location struct{
		Name string `json:"name"`
		Region string `json:"region"`
		Country string `json:"country"`
	}`json:"location"`

	Current struct{
		LastUpdated string `json:"last_updated"`
		TempC float64 `json:"temp_c"`
		Condition struct {
			Text string `json:"text"`
		} `json:"condition"`
		Wind float64 `json:"wind_mph"`
		FeelsLike float64 `json:"feelslike_c"`
	} `json:"current"`

	Forecast struct{
		ForecastDay []struct {
			Hour []struct{
				TimeEpoch int64 `json:"time_epoch"`
				TempC float64 `json:"temp_c"`
				Condition struct {
					Text string `json:"text"`
				} `json:"condition"`
				ChanceOfRain float64 `json:"chance_of_rain"`
			} `json:"hour"`
		}`json:"forecastday"`
	} `json:"forecast"`
	
}

func goDotEnvVar(key string) string {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	  }

	return os.Getenv(key)
}

func main() {
	q := "London"
	if len(os.Args) >= 2{
		q = os.Args[1]
	}
	API_KEY := goDotEnvVar("WEATHERAPI_KEY")
	res, err := http.Get("http://api.weatherapi.com/v1/forecast.json?q="+ q +"&days=1&key="+API_KEY)
	

	if err !=nil {
		panic(err)
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		panic("Error connecting to Weather API")

	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	// fmt.Println(string(body))
	var weather Weather 
	err = json.Unmarshal(body, &weather)
	if err != nil {
		panic(err)
	}
	location, current, hours := weather.Location, weather.Current, weather.Forecast.ForecastDay[0].Hour


	fmt.Printf(
		"%s, %s: %s, %.0f°C but feels like %.0f°C, %.0fmph winds\n",
		location.Name,
		location.Country,
		current.Condition.Text,
		current.TempC,
		current.FeelsLike,
		current.Wind,
	)

	for _, hour := range hours {
		date := time.Unix(hour.TimeEpoch,0)

		if date.Before(time.Now()){
			continue
		}

		message := fmt.Sprintf(
			"%s - %.0f°C, %.0f%% chance of rain, %s\n",
		date.Format("15:04"),
		hour.TempC,
		hour.ChanceOfRain,
		hour.Condition.Text,
		)

		if hour.ChanceOfRain < 50 {
			fmt.Print(message)
		} else {
			color.Red(message)
		}
	}
}
