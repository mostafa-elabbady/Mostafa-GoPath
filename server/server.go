package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
)

func BuildServer() (*http.Server, error) {
	router := chi.NewRouter()

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
