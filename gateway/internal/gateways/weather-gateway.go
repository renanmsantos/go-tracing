package gateways

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
)

type Response struct {
	City              string  `json:"city"`
	Celsius           float64 `json:"temp_c"`
	Fahrenheit        float64 `json:"temp_f"`
	KelvinTemperature float64 `json:"temp_K"`
}

func GetLocationAndTemperature(ctx context.Context, cep string) (Response, error) {
	url := fmt.Sprintf("http://go-weather:8081?cep=%s", cep)
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return Response{}, err
	}

	otel.GetTextMapPropagator().Inject(ctx, propagation.HeaderCarrier(req.Header))

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return Response{}, err
	}
	defer res.Body.Close()

	resp, err := io.ReadAll(res.Body)
	if err != nil {
		return Response{}, err
	}

	var response Response
	err = json.Unmarshal(resp, &response)
	if err != nil {
		return Response{}, err
	}
	return response, nil

}
