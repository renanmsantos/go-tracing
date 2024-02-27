package gateways

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/renanmoreirasan/go-weather/infra/configs"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
)

type Location struct {
	Lat  string `json:"lat"`
	Lng  string `json:"lng"`
	City string `json:"city"`
}

func GetLocation(ctx context.Context, cep string) (Location, error) {

	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	ctx, span := configs.GetTracer().Start(ctx, "START call to CEP API")
	defer span.End()
	req, _ := http.NewRequestWithContext(ctx, "GET", "http://cep.awesomeapi.com.br/json/"+cep, nil)
	otel.GetTextMapPropagator().Inject(ctx, propagation.HeaderCarrier(req.Header))
	res, _ := http.DefaultClient.Do(req)
	if res.StatusCode == http.StatusNotFound {
		return Location{}, errors.New("CEP_NOT_FOUND")
	}
	defer res.Body.Close()

	resp, err := io.ReadAll(res.Body)
	if err != nil {
		return Location{}, err
	}

	var coordinates Location
	err = json.Unmarshal(resp, &coordinates)
	if err != nil {
		return Location{}, err
	}
	return coordinates, nil
}
