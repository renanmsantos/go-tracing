package webserver

import (
	"encoding/json"
	"net/http"

	"github.com/renanmoreirasan/go-weather/internal/usecases"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/trace"
)

func GetTemperature(tracer trace.Tracer) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		carrier := propagation.HeaderCarrier(r.Header)
		ctx := r.Context()
		ctx = otel.GetTextMapPropagator().Extract(ctx, carrier)

		ctx, span := tracer.Start(ctx, "START go-weather")
		defer span.End()

		w.Header().Set("Content-Type", "application/json")
		input := r.URL.Query().Get("cep")
		output, err := usecases.Execute(ctx, input)
		if err != nil && err.Error() == "INVALID_CEP" {
			http.Error(w, err.Error(), http.StatusUnprocessableEntity)
			json.NewEncoder(w).Encode("Invalid zipcode")
			return
		}
		if err != nil && err.Error() == "CEP_NOT_FOUND" {
			http.Error(w, err.Error(), http.StatusNotFound)
			json.NewEncoder(w).Encode("Can not found zipcode")
			return
		}
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			json.NewEncoder(w).Encode(err.Error())
			return
		}
		err = json.NewEncoder(w).Encode(output)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
