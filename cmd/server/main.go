package main

import (
	"log"
	"net/http"

	config "github.com/LucasBelusso1/go-temperatureByZipCode/configs"
	"github.com/LucasBelusso1/go-temperatureByZipCode/internal/webserver/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func init() {
	config.LoadConfig()
	log.Printf("Using config %+v", config.GetConfig())
}

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get(`/{cep:^\d{8}$}`, handlers.GetTemperatureByZipCode)

	configs := config.GetConfig()
	http.ListenAndServe(":"+configs.Port, r)
	log.Printf("Listening on port %s", configs.Port)
}
