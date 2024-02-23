package webserver

import (
	"encoding/json"
	"net/http"

	"github.com/renanmsantos/go-gateway/internal/usecases"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/trace"
)

func ValidationHandler(tracer trace.Tracer) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		carrier := propagation.HeaderCarrier(r.Header)
		ctx := r.Context()
		ctx = otel.GetTextMapPropagator().Extract(ctx, carrier)

		ctx, span := tracer.Start(ctx, "START go-gateway")
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
			if err != nil && err.Error() == "INVALID_CEP" {
				http.Error(w, "Invalid zipcode", http.StatusUnprocessableEntity)
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
