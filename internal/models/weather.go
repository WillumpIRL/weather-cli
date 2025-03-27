package models

import "time"

type WeatherResponse struct {
	Name      string `json:"name"`
	Timezone  int    `json:"timezone"` // Timezone offset in seconds
	LocalTime time.Time

	Main struct {
		Temp     float64 `json:"temp"`
		Humidity int     `json:"humidity"`
	} `json:"main"`

	Weather []struct {
		Description string `json:"description"`
	} `json:"weather"`
}
