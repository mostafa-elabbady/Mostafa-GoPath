package handlers

import (
	"github.com/go-chi/chi/v5"
	"github.com/mostafa-elabbady/Mostafa-GoPath/api"
)

const URL_CURRENT_WEATHER = "/current-weather"

type router struct {
	Mux     chi.Router
	Service api.Service
}

func NewRouter(
	mux chi.Router,
	service api.Service) router {

	return router{
		Mux:     mux,
		Service: service,
	}
}

func (router *router) AddRoutes() {
	router.Mux.Group(func(r chi.Router) {
		r.Get(URL_CURRENT_WEATHER, router.GetCurrentWeatherHandler)
	})
}
