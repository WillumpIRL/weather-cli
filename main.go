package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"sync"
	"time"
	"unicode"

	"github.com/joho/godotenv"
)

// weather_data represents the weather information for a city
type weather_data struct {
	City        string    `json:"city"`
	Temperature float64   `json:"temperature"`
	Humidity    int       `json:"humidity"`
	Description string    `json:"description"`
	WindSpeed   float64   `json:"wind_speed"`
	LastUpdated time.Time `json:"last_updated"`
	LocalTime   time.Time `json:"local_time"`
	Error       string    `json:"error,omitempty"`
}

// open_weather_response represents the API response from OpenWeatherMap
type open_weather_response struct {
	Main struct {
		Temp     float64 `json:"temp"`
		Humidity int     `json:"humidity"`
	} `json:"main"`
	Weather []struct {
		Description string `json:"description"`
	} `json:"weather"`
	Wind struct {
		Speed float64 `json:"speed"`
	} `json:"wind"`
	Sys struct {
		Country string `json:"country"`
		Sunrise int64  `json:"sunrise"`
		Sunset  int64  `json:"sunset"`
	} `json:"sys"`
	Timezone int64 `json:"timezone"`
}

var (
	api_key string
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	api_key = os.Getenv("OPENWEATHER_API_KEY")
	if api_key == "" {
		log.Fatal("OPENWEATHER_API_KEY is required")
	}
}

// capitalise_words properly capitalises each word in a string
func capitalise_words(s string) string {
	words := strings.Fields(s)
	for i, word := range words {
		runes := []rune(word)
		if len(runes) > 0 {
			runes[0] = unicode.ToUpper(runes[0])
			for j := 1; j < len(runes); j++ {
				runes[j] = unicode.ToLower(runes[j])
			}
			words[i] = string(runes)
		}
	}
	return strings.Join(words, " ")
}

func fetch_weather(city string) weather_data {
	// Properly capitalise the city name
	proper_city := capitalise_words(city)

	// URL encode the city name to handle spaces and special characters
	encoded_city := url.QueryEscape(proper_city)
	url := fmt.Sprintf("http://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s&units=metric", encoded_city, api_key)

	resp, err := http.Get(url)
	if err != nil {
		return weather_data{
			City:  proper_city,
			Error: fmt.Sprintf("Error fetching weather: %v", err),
		}
	}
	defer resp.Body.Close()

	// Check if the response status is not OK
	if resp.StatusCode != http.StatusOK {
		return weather_data{
			City:  proper_city,
			Error: "City Not Found",
		}
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return weather_data{
			City:  proper_city,
			Error: fmt.Sprintf("Error reading response: %v", err),
		}
	}

	var weather_response open_weather_response
	if err := json.Unmarshal(body, &weather_response); err != nil {
		return weather_data{
			City:  proper_city,
			Error: fmt.Sprintf("Error parsing response: %v", err),
		}
	}

	// Calculate local time using the timezone offset from the API
	local_time := time.Now().UTC().Add(time.Duration(weather_response.Timezone) * time.Second)

	return weather_data{
		City:        proper_city,
		Temperature: weather_response.Main.Temp,
		Humidity:    weather_response.Main.Humidity,
		Description: weather_response.Weather[0].Description,
		WindSpeed:   weather_response.Wind.Speed,
		LastUpdated: time.Now(),
		LocalTime:   local_time,
	}
}

func display_weather(weather weather_data) {
	fmt.Printf("\n=== Weather for %s ===\n", weather.City)
	if weather.Error != "" {
		fmt.Printf("Error: %s\n", weather.Error)
		return
	}
	fmt.Printf("Temperature: %.1f°C\n", weather.Temperature)
	fmt.Printf("Humidity: %d%%\n", weather.Humidity)
	fmt.Printf("Conditions: %s\n", weather.Description)
	fmt.Printf("Wind Speed: %.1f m/s\n", weather.WindSpeed)
	fmt.Printf("Local Time: %s\n", weather.LocalTime.Format("15:04:05"))
	fmt.Printf("Last Updated: %s\n", weather.LastUpdated.Format("15:04:05"))
	fmt.Println("=========================")
}

func get_city_input() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("\nEnter a city name (or 'q' to quit): ")
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Printf("Error reading input: %v", err)
		return ""
	}
	return strings.TrimSpace(input)
}

func main() {
	fmt.Println("=======================================")
	fmt.Println("Weather CLI - Demonstrating Go's Features")
	fmt.Println("This programme demonstrates Go's concurrency features with goroutines and channels")
	fmt.Println("Type 'q' to quit the programme")
	fmt.Println("=======================================")

	for {
		city := get_city_input()
		if city == "q" {
			fmt.Println("\nThank you for using Weather CLI!")
			break
		}
		if city == "" {
			continue
		}

		// Create a channel to receive weather data
		weather_channel := make(chan weather_data)
		var wg sync.WaitGroup

		// Launch goroutine for the city
		wg.Add(1)
		go func(city string) {
			defer wg.Done()
			weather := fetch_weather(city)
			weather_channel <- weather
		}(city)

		// Launch a goroutine to close the channel when fetch is complete
		go func() {
			wg.Wait()
			close(weather_channel)
		}()

		// Display weather data as it arrives
		for weather := range weather_channel {
			display_weather(weather)
		}
	}

	fmt.Println("\nGo Language Features Demonstrated:")
	fmt.Println("=======================================")
	fmt.Println("Strengths:")
	fmt.Println("- Concurrent API calls using goroutines")
	fmt.Println("- Strong type system with structs")
	fmt.Println("- Efficient memory usage")
	fmt.Println("- Simple error handling")
	fmt.Println("- Channel-based communication")
	fmt.Println("\nLimitations:")
	fmt.Println("- Verbose error handling")
	fmt.Println("- Manual JSON marshalling")
	fmt.Println("- Limited generics usage")
	fmt.Println("=======================================")
}
