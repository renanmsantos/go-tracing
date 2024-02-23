package usecases

type Output struct {
	City                  string  `json:"city"`
	CelsiusTemperature    float64 `json:"temp_C"`
	FahrenheitTemperature float64 `json:"temp_F"`
	KelvinTemperature     float64 `json:"temp_K"`
}

func NewOutput(city string, celsiusTemperature float64, fahrenheitTemperature float64, kelvinTemperature float64) Output {
	return Output{
		City:                  city,
		CelsiusTemperature:    celsiusTemperature,
		FahrenheitTemperature: fahrenheitTemperature,
		KelvinTemperature:     kelvinTemperature,
	}
}
