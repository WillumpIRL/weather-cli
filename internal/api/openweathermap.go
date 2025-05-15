package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/WillumpIRL/weather-cli/internal/models"
)

type OpenWeatherMapAPI struct {
	apiKey string
}

func NewOpenWeatherMapAPI(apiKey string) *OpenWeatherMapAPI {
	return &OpenWeatherMapAPI{apiKey: apiKey}
}

func (api *OpenWeatherMapAPI) GetWeather(city string) (*models.WeatherResponse, error) {
	encodedCity := url.QueryEscape(city)
	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s&units=metric", encodedCity, api.apiKey)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API error: %s, response: %s", resp.Status, string(body))
	}

	var weather models.WeatherResponse
	if err := json.Unmarshal(body, &weather); err != nil {
		return nil, fmt.Errorf("failed to parse response: %v, body: %s", err, string(body))
	}

	if len(weather.Weather) == 0 {
		return nil, fmt.Errorf("no weather data found for %s", city)
	}

	// Calculate local time
	weather.LocalTime = time.Now().UTC().Add(time.Second * time.Duration(weather.Timezone))

	return &weather, nil
}
