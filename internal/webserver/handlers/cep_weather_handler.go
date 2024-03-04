package handlers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/url"
	"regexp"

	config "github.com/LucasBelusso1/go-temperatureByZipCode/configs"
	"github.com/LucasBelusso1/go-temperatureByZipCode/internal/dto"
	"github.com/go-chi/chi/v5"
)

func GetTemperatureByZipCode(w http.ResponseWriter, r *http.Request) {
	cep := chi.URLParam(r, "cep")
	matched, err := regexp.MatchString(`^[0-9]{8}$`, cep)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	if !matched {
		w.WriteHeader(http.StatusUnprocessableEntity)
		w.Write([]byte("Invalid CEP"))
		return
	}

	cepResponse, err := requestCEP(cep)
	log.Printf("CEP: %s", cep)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	if cepResponse.Localidade == "" {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Can not find zipcode"))
		return
	}

	log.Printf("Localidade %s", cepResponse.Localidade)

	weatherResponse, err := requestWeather(cepResponse)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	if weatherResponse == nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Something went wrong"))
		return
	}

	apiResponse := dto.ApiOutput{
		TempC: weatherResponse.Current.TempC,
		TempF: weatherResponse.Current.TempF,
		TempK: weatherResponse.Current.TempC + 273.15,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(apiResponse)
}

func requestCEP(cep string) (dto.ViaCepOutput, error) {
	var viaCepDto dto.ViaCepOutput

	req, err := http.NewRequest("GET", "http://viacep.com.br/ws/"+cep+"/json/", nil)

	if err != nil {
		return viaCepDto, err
	}

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		return viaCepDto, err
	}

	if res.StatusCode != 200 {
		return viaCepDto, err
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)

	if err != nil {
		return viaCepDto, err
	}

	err = json.Unmarshal(body, &viaCepDto)

	if err != nil {
		return viaCepDto, err
	}

	return viaCepDto, nil
}

func requestWeather(data dto.ViaCepOutput) (*dto.WeatherOutput, error) {
	log.Printf("Got into requestWeather")
	var weatherDto *dto.WeatherOutput
	configs := config.GetConfig()

	url := "http://api.weatherapi.com/v1/current.json?key=" + configs.WeatherApiKey + "&q=" + url.QueryEscape(data.Localidade)

	res, err := http.Get(url)

	log.Printf("Made the weather request")

	if err != nil {
		return nil, err
	}

	if res.StatusCode != 200 {
		return nil, err
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)

	if err != nil {
		return weatherDto, err
	}

	err = json.Unmarshal(body, &weatherDto)

	if err != nil {
		return weatherDto, err
	}

	return weatherDto, nil
}
