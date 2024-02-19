package main

import (
	"log"
	"net/http"

	config "github.com/LucasBelusso1/go-temperatureByZipCode/configs"
	"github.com/LucasBelusso1/go-temperatureByZipCode/internal/webserver/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	err := config.LoadConfig(".")

	if err != nil {
		log.Fatalln("Couldn't read configurations", err)
		panic(err)
	}

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get(`/{cep:^\d{8}$}`, handlers.GetTemperatureByZipCode)

	http.ListenAndServe(":8080", r)
}
