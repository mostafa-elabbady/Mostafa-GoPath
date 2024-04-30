package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (router *router) GetCurrentWeatherHandler(w http.ResponseWriter, r *http.Request) {
	lat := r.URL.Query().Get("lat")
	lon := r.URL.Query().Get("lon")

	if len(lat) == 0 || len(lon) == 0 {
		fmt.Printf("Missing required longitude or latitude parameters")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	weatherResponse, err := router.Service.GetCurrentWeather(lat, lon)
	if err != nil {
		fmt.Printf("Error getting the weather: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(weatherResponse)
	if err != nil {
		fmt.Printf("Error JSON encoding response body: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

}
