package gateways

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
)

type Temperatures struct {
	Current Current `json:"current"`
}

type Current struct {
	Celsius    float64 `json:"temp_c"`
	Fahrenheit float64 `json:"temp_f"`
}

const apiKey = "369d49d6a467440489b202319242601"

func GetLocationTemperature(ctx context.Context, location Location) (Temperatures, error) {
	latitude, err := strconv.ParseFloat(location.Lat, 64)
	if err != nil {
		return Temperatures{}, err
	}
	longitude, err := strconv.ParseFloat(location.Lng, 64)
	if err != nil {
		return Temperatures{}, err
	}
	url := fmt.Sprintf("http://api.weatherapi.com/v1/current.json?key=%s&q=%v,%v", apiKey, latitude, longitude)
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return Temperatures{}, err
	}

	otel.GetTextMapPropagator().Inject(ctx, propagation.HeaderCarrier(req.Header))
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return Temperatures{}, err
	}
	defer res.Body.Close()

	resp, err := io.ReadAll(res.Body)
	if err != nil {
		return Temperatures{}, err
	}

	var temperatures Temperatures
	err = json.Unmarshal(resp, &temperatures)
	if err != nil {
		return Temperatures{}, err
	}
	return temperatures, nil

}
