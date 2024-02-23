package usecases

import (
	"context"
	"errors"
	"regexp"

	"github.com/renanmoreirasan/go-weather/internal/gateways"
)

func Execute(ctx context.Context, cep string) (OutputDTO, error) {
	//CEP validation
	if !isValidCep(cep) {
		return OutputDTO{}, errors.New("INVALID_CEP")
	}
	//Call API to get location
	location, err := gateways.GetLocation(ctx, cep)
	if err != nil {
		return OutputDTO{}, err
	}
	//Call API to get temperature
	temperatures, err := gateways.GetLocationTemperature(ctx, location)
	if err != nil {
		return OutputDTO{}, err
	}

	//Print temperature
	return NewOutputDTO(
		location.City,
		temperatures.Current.Celsius,
		temperatures.Current.Fahrenheit,
		convertCelsiusToKelvin(temperatures.Current.Celsius),
	), nil
}

func isValidCep(cep string) bool {
	containsEightDigitsRegex := `^[0-9]{8}$`
	match, _ := regexp.MatchString(containsEightDigitsRegex, cep)
	return match
}

func convertCelsiusToKelvin(celsiusTemperature float64) float64 {
	return celsiusTemperature + 273.15
}
