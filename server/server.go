package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/mostafa-elabbady/Mostafa-GoPath/api"
	"github.com/mostafa-elabbady/Mostafa-GoPath/api/handlers"
	current_weather "github.com/mostafa-elabbady/Mostafa-GoPath/clients/current-weather"
)

func BuildServer() (*http.Server, error) {
	router := chi.NewRouter()
	currentWeatherClient, err := current_weather.NewCurrentWeatherClient()
	service := api.NewService(&currentWeatherClient)
	apiRouter := handlers.NewRouter(router, service)
	apiRouter.AddRoutes()

	if err != nil {
		return nil, err
	}

	port := "8080"
	server := &http.Server{
		Addr:         fmt.Sprintf("%s:%s", "0.0.0.0", port),
		Handler:      router,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	return server, nil
}
