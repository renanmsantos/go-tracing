package usecases

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShouldExecutionReturnErrorWhenLocationIsNotFound(t *testing.T) {

	_, err := Execute("12345678")

	assert.Equal(t, "CEP_NOT_FOUND", err.Error())
}

func TestShouldExecutionReturnErrorWhenCepIsInvalid(t *testing.T) {
	cep := "1234567"
	_, err := Execute(cep)
	assert.Equal(t, "INVALID_CEP", err.Error())
}

func TestShouldBeValidCEP(t *testing.T) {
	cep := "12345678"
	isValid := isValidCep(cep)
	assert.True(t, isValid)
}

func TestShouldNotBeValidCEP(t *testing.T) {
	cep := "1234567"
	isValid := isValidCep(cep)
	assert.False(t, isValid)
}

func TestShouldConvertCelsiusToKelvin(t *testing.T) {
	celsiusTemperature := 20.0
	kelvinTemperature := convertCelsiusToKelvin(celsiusTemperature)
	assert.Equal(t, 293.15, kelvinTemperature)
}
