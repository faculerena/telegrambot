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

type WeatherDataForecast struct {
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
	Forecast struct {
		Forecastday []struct {
			Date      string `json:"date"`
			DateEpoch int    `json:"date_epoch"`
			Day       struct {
				MaxtempC          float64 `json:"maxtemp_c"`
				MintempC          float64 `json:"mintemp_c"`
				AvgtempC          float64 `json:"avgtemp_c"`
				MaxwindKph        float64 `json:"maxwind_kph"`
				TotalprecipMm     float64 `json:"totalprecip_mm"`
				TotalsnowCm       float64 `json:"totalsnow_cm"`
				AvgvisKm          float64 `json:"avgvis_km"`
				Avghumidity       float64 `json:"avghumidity"`
				DailyWillItRain   int     `json:"daily_will_it_rain"`
				DailyChanceOfRain int     `json:"daily_chance_of_rain"`
				DailyWillItSnow   int     `json:"daily_will_it_snow"`
				DailyChanceOfSnow int     `json:"daily_chance_of_snow"`
				Condition         struct {
				} `json:"condition"`
				Uv float64 `json:"uv"`
			} `json:"day"`
			Astro struct {
				Sunrise          string `json:"sunrise"`
				Sunset           string `json:"sunset"`
				Moonrise         string `json:"moonrise"`
				Moonset          string `json:"moonset"`
				MoonPhase        string `json:"moon_phase"`
				MoonIllumination string `json:"moon_illumination"`
				IsMoonUp         int    `json:"is_moon_up"`
				IsSunUp          int    `json:"is_sun_up"`
			} `json:"astro"`
			Hour []struct {
				TimeEpoch int     `json:"time_epoch"`
				Time      string  `json:"time"`
				TempC     float64 `json:"temp_c"`
				IsDay     int     `json:"is_day"`
				Condition struct {
				} `json:"condition"`
				WindKph      float64 `json:"wind_kph"`
				WindDegree   int     `json:"wind_degree"`
				WindDir      string  `json:"wind_dir"`
				PressureMb   float64 `json:"pressure_mb"`
				PrecipMm     float64 `json:"precip_mm"`
				Humidity     int     `json:"humidity"`
				Cloud        int     `json:"cloud"`
				FeelslikeC   float64 `json:"feelslike_c"`
				WindchillC   float64 `json:"windchill_c"`
				HeatindexC   float64 `json:"heatindex_c"`
				DewpointC    float64 `json:"dewpoint_c"`
				WillItRain   int     `json:"will_it_rain"`
				ChanceOfRain int     `json:"chance_of_rain"`
				WillItSnow   int     `json:"will_it_snow"`
				VisKm        float64 `json:"vis_km"`
				GustKph      float64 `json:"gust_kph"`
				Uv           float64 `json:"uv"`
			} `json:"hour"`
		} `json:"forecastday"`
	} `json:"forecast"`
}

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

func makeUrl(key string, gps string, what string) (url string) {

	switch what {
	case "current":
		url = fmt.Sprintf("https://api.weatherapi.com/v1/current.json?key=" + key + "&q=" + gps + "&aqi=no")
		return url

	case "forecast":
		//http://api.weatherapi.com/v1/forecast.json?key= 579a11c94dc946fb84a230435230101&q=mar de ajo&days=1&aqi=no&alerts=no
		url = fmt.Sprintf("https://api.weatherapi.com/v1/forecast.json?key=" + key + "&q=" + gps + "&days=3" + "&aqi=no")
		return url
	}
	return ""
}

func GetCurrent(gps string) (WeatherDataCurrent, error) {
	var weatherData WeatherDataCurrent
	weatherKey := os.Getenv("API_KEY")

	response, err := http.Get(makeUrl(weatherKey, gps, "current"))
	if err != nil {
		log.Println("Error in response:", err)
		return weatherData, err
	}
	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&weatherData)
	if err != nil {
		log.Println("Error: decoding", err)
		return weatherData, err
	}

	return weatherData, nil

}

func Current(update tgbotapi.Update, gps string) tgbotapi.MessageConfig {
	weatherData, err := GetCurrent(gps)
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

func GetForecast(gps string) (WeatherDataForecast, error) {

	var weatherData WeatherDataForecast
	weatherKey := os.Getenv("API_KEY")

	response, err := http.Get(makeUrl(weatherKey, gps, "forecast"))
	if err != nil {
		log.Println("Error in response:", err)
		return weatherData, err
	}
	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&weatherData)
	if err != nil {
		log.Println("Error: decoding", err)
		return weatherData, err
	}

	return weatherData, nil
}

func Forecast(update tgbotapi.Update, gps string) tgbotapi.MessageConfig {
	weatherData, err := GetForecast(gps)
	if err != nil {
		log.Println(err)
	}
	/*

		t, err := time.Parse("2006-01-02 15:04", weatherData.Location.Localtime)
		if err != nil {
			log.Println(err)
		}
		formattedTime := t.Format("02/01")

	*/

	var allDaysForecast []string

	for i := 0; i <= 2; i++ {
		min := weatherData.Forecast.Forecastday[i].Day.MintempC
		max := weatherData.Forecast.Forecastday[i].Day.MaxtempC
		h := weatherData.Forecast.Forecastday[i].Day.Avghumidity
		rain := weatherData.Forecast.Forecastday[i].Day.DailyChanceOfRain

		ms := fmt.Sprintf("    -Max: %v°C\n    -Min: %v°C\n    -Humedad: %v%%\n    -Precipitaciones: %v%%\n", max, min, h, rain)
		allDaysForecast = append(allDaysForecast, ms)
	}

	msgContent := fmt.Sprintf("> Mañana:\n%v> Pasado Mañana:\n%v>PasadoPasado Mañana:\n%v", allDaysForecast[0], allDaysForecast[1], allDaysForecast[2])
	return tgbotapi.NewMessage(update.Message.Chat.ID, msgContent)

}
