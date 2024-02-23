package usecases

import (
	"context"
	"errors"
	"regexp"

	"github.com/renanmsantos/go-gateway/internal/gateways"
)

func Execute(ctx context.Context, input Input) (Output, error) {
	//CEP validation
	if !isValidCep(input.Cep) {
		return Output{}, errors.New("INVALID_CEP")
	}
	//Get gateway
	response, err := gateways.GetLocationAndTemperature(ctx, input.Cep)
	if err != nil {
		return Output{}, err
	}

	//Print response
	return NewOutput(
		response.City,
		2,
		3,
		4,
	), nil
}

func isValidCep(cep string) bool {
	containsEightDigitsRegex := `^[0-9]{8}$`
	match, _ := regexp.MatchString(containsEightDigitsRegex, cep)
	return match
}
