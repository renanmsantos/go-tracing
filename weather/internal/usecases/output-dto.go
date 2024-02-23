package usecases

type OutputDTO struct {
	City                  string  `json:"city"`
	CelsiusTemperature    float64 `json:"temp_C"`
	FahrenheitTemperature float64 `json:"temp_F"`
	KelvinTemperature     float64 `json:"temp_K"`
}

func NewOutputDTO(city string, celsiusTemperature float64, fahrenheitTemperature float64, kelvinTemperature float64) OutputDTO {
	return OutputDTO{
		City:                  city,
		CelsiusTemperature:    celsiusTemperature,
		FahrenheitTemperature: fahrenheitTemperature,
		KelvinTemperature:     kelvinTemperature,
	}
}
