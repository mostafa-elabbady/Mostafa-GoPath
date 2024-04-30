package api

import (
	"fmt"

	current_weather "github.com/mostafa-elabbady/Mostafa-GoPath/clients/current-weather"
	"github.com/mostafa-elabbady/Mostafa-GoPath/models"
)

type Service interface {
	GetCurrentWeather(lat, lon string) (models.CurrentWeatherResponse, error)
}

type service struct {
	currentWeatherClient current_weather.Client
}

func NewService(cwc current_weather.Client) service {
	return service{
		currentWeatherClient: cwc,
	}
}

func (s service) GetCurrentWeather(lat, lon string) (models.CurrentWeatherResponse, error) {
	response, err := s.currentWeatherClient.GetCurrentWeather(lat, lon)
	if err != nil {
		return models.CurrentWeatherResponse{}, err
	}
	var currentWeatherResponse models.CurrentWeatherResponse
	currentWeatherResponse.City = response.Name

	var weatherDesc string
	for i, value := range response.Weather {
		if i == 0 {
			weatherDesc += value.Description

		} else {
			weatherDesc += "," + value.Description + ","
		}
		currentWeatherResponse.Condition = append(currentWeatherResponse.Condition, value.Description)
	}

	var feelsLike string
	switch {
	case response.Feels_Like < 40:
		feelsLike = "cold"
	case response.Feels_Like >= 40 && response.Feels_Like < 50:
		feelsLike = "cool"
	case response.Feels_Like >= 50 && response.Feels_Like < 70:
		feelsLike = "slightly cool"
	case response.Feels_Like >= 70 && response.Feels_Like < 80:
		feelsLike = "slightly warm"
	case response.Feels_Like >= 80 && response.Feels_Like < 85:
		feelsLike = "warm"
	case response.Feels_Like >= 85:
		feelsLike = "hot"
	}

	currentWeatherResponse.Temperature = response.Temp

	currentWeatherResponse.Overview = fmt.Sprintf("It's currently %.0f degrees outside and %s with conditions of %s and %.0f mph winds. "+
		"Today's high will be %.0f, with a low of %.0f, and it feels like %.0f degrees.", response.Temp, feelsLike, weatherDesc,
		response.Speed, response.Main.Temp_Max, response.Main.Temp_Min, response.Main.Feels_Like)

	return currentWeatherResponse, nil
}
