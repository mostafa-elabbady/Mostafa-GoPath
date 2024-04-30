package current_weather

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/mostafa-elabbady/Mostafa-GoPath/models"
)

type Client interface {
	GetCurrentWeather(lat, lon string) (models.OpenWeatherCurrentWeatherResponse, error)
}

func (c *currentWeatherClient) GetCurrentWeather(lat, lon string) (models.OpenWeatherCurrentWeatherResponse, error) {
	appId, ok := os.LookupEnv("APP_ID_CURRENT_WEATHER")
	if !ok {
		err := errors.New("error getting environment variable APP_ID_CURRENT_WEATHER")
		return models.OpenWeatherCurrentWeatherResponse{}, err
	}

	// Create an http.request object to use when the client call is made
	newRequest, err := http.NewRequest(http.MethodGet, c.baseUri.String(), nil)
	if err != nil {
		return models.OpenWeatherCurrentWeatherResponse{}, err
	}

	// Take the query parameters passed in and add those to the URL, formatted correctly
	query := newRequest.URL.Query()
	query.Add("lat", lat)
	query.Add("lon", lon)
	query.Add("appId", appId)
	query.Add("units", "imperial")
	newRequest.URL.RawQuery = query.Encode()

	// Actually do the client request
	response, err := c.client.Do(newRequest)
	if err != nil {
		log.Printf("Current weather client: Error making request: %v", err)
		return models.OpenWeatherCurrentWeatherResponse{}, err
	}
	defer response.Body.Close()

	var weatherResponse models.OpenWeatherCurrentWeatherResponse
	switch response.StatusCode {
	case http.StatusOK:
		// Read body sent back from open weather call
		body, err := io.ReadAll(response.Body)
		if err != nil {
			return models.OpenWeatherCurrentWeatherResponse{}, err
		}

		// Unmarshal json from open weather response into our currentWeatherResponse struct
		err = json.Unmarshal(body, &weatherResponse)
		if err != nil {
			return models.OpenWeatherCurrentWeatherResponse{}, err
		}
	}

	// Successful return of weather data
	return weatherResponse, nil

}
