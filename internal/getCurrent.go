package internal

import (
	"encoding/json"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"net/http"
	"os"
	"time"
)

type WeatherDataCurrent struct {
	Location struct {
		Name           string  `json:"name"`
		Region         string  `json:"region"`
		Country        string  `json:"country"`
		Lat            float64 `json:"lat"`
		Lon            float64 `json:"lon"`
		TzId           string  `json:"tz_id"`
		LocaltimeEpoch int     `json:"localtime_epoch"`
		Localtime      string  `json:"localtime"`
	} `json:"location"`
	Current struct {
		TempC     float64 `json:"temp_c"`
		IsDay     int     `json:"is_day"`
		Condition struct {
			Text string `json:"text"`
		} `json:"condition"`
		WindKph    float64 `json:"wind_kph"`
		WindDegree int     `json:"wind_degree"`
		WindDir    string  `json:"wind_dir"`
		PressureMb float64 `json:"pressure_mb"`
		PrecipMm   float64 `json:"precip_mm"`
		Humidity   int     `json:"humidity"`
		Cloud      int     `json:"cloud"`
		FeelslikeC float64 `json:"feelslike_c"`
		VisKm      float64 `json:"vis_km"`
		Uv         float64 `json:"uv"`
		GustKph    float64 `json:"gust_kph"`
	} `json:"current"`
}

func makeUrl(key string) (url string) {

	gps := "Buenos_Aires"
	url = fmt.Sprintf("https://api.weatherapi.com/v1/current.json?key=" + key + "&q=" + gps + "&aqi=no")
	fmt.Println(url)
	return url
}

func GetCurrent() (WeatherDataCurrent, error) {
	var weatherData WeatherDataCurrent
	weatherKey := os.Getenv("API_KEY")

	response, err := http.Get(makeUrl(weatherKey))
	if err != nil {
		fmt.Println("Error in response:", err)
		return weatherData, err
	}
	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&weatherData)
	if err != nil {
		fmt.Println("Error: decoding", err)
		return weatherData, err
	}

	return weatherData, nil

}

func Current(update tgbotapi.Update) tgbotapi.MessageConfig {
	weatherData, err := GetCurrent()
	if err != nil {
		log.Println(err)
	}
	t, err := time.Parse("2006-01-02 15:04", weatherData.Location.Localtime)
	if err != nil {
		log.Println(err)
	}
	formattedTime := t.Format("15:04")

	msgContent := fmt.Sprintf("Hora: %s\nTemperatura actual en %s: %v°C\nHumedad: %v%%", formattedTime, weatherData.Location.Name, weatherData.Current.TempC, weatherData.Current.Humidity)
	return tgbotapi.NewMessage(update.Message.Chat.ID, msgContent)

}
