package webserver

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/renanmsantos/go-gateway/infra/configs"
	"github.com/renanmsantos/go-gateway/internal/usecases"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
)

func ValidationHandler() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		carrier := propagation.HeaderCarrier(r.Header)
		ctx := r.Context()
		ctx = otel.GetTextMapPropagator().Extract(ctx, carrier)

		ctx, span := configs.GetTracer().Start(ctx, "START go-gateway")
		defer span.End()

		w.Header().Set("Content-Type", "application/json")
		switch r.Method {
		case http.MethodPost:
			var input usecases.Input
			err := json.NewDecoder(r.Body).Decode(&input)
			if err != nil {
				http.Error(w, "Invalid zipcode", http.StatusUnprocessableEntity)
				return
			}
			output, err := usecases.Execute(ctx, input)
			fmt.Println(output, err)
			if err != nil && err.Error() == "INVALID_CEP" {
				http.Error(w, "Invalid zipcode", http.StatusUnprocessableEntity)
				return
			}
			if err != nil && err.Error() == "CEP_NOT_FOUND" {
				http.Error(w, "Can not found zipcode", http.StatusNotFound)
				return
			}
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			err = json.NewEncoder(w).Encode(output)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		default:
			http.Error(w, "METHOD_NOT_ALLOWED", http.StatusMethodNotAllowed)
			return
		}

	}
}
