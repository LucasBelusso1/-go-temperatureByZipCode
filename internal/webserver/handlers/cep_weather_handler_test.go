package handlers

import (
	"os"
	"path/filepath"
	"testing"

	config "github.com/LucasBelusso1/go-temperatureByZipCode/configs"
	"github.com/LucasBelusso1/go-temperatureByZipCode/internal/dto"
	"github.com/stretchr/testify/assert"
)

func TestRequestCEP(t *testing.T) {
	responseDto, err := requestCEP("04566-090")
	assert.NoError(t, err)
	assert.NotNil(t, responseDto)
	assert.IsType(t, responseDto, dto.ViaCepOutput{})
	assert.NotEmpty(t, responseDto.Localidade)
}

func TestRequestWeather(t *testing.T) {
	cmdDir, err := filepath.Abs("../../../cmd/server/")
	assert.NoError(t, err)

	err = os.Chdir(cmdDir)
	assert.NoError(t, err)

	defer func() {
		err := os.Chdir(filepath.Dir(cmdDir))
		assert.NoError(t, err)
	}()

	config.LoadConfig()

	assert.NoError(t, err, "Error returned when trying to read .env file.")

	responseDto, err := requestWeather(dto.ViaCepOutput{Localidade: "São Paulo"})
	assert.NoError(t, err)
	assert.NotNil(t, responseDto)
	assert.IsType(t, responseDto, dto.WeatherOutput{})
	assert.NotEmpty(t, responseDto.Current.TempC)
}
