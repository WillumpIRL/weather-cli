# Weather CLI

A command-line weather application built with Go that demonstrates the language's strengths and limitations.

## Features

- Real-time weather data for multiple cities
- Concurrent API calls using goroutines and channels
- Clean command-line interface
- Error handling demonstration
- Type safety showcase

## Prerequisites

- Go 1.21 or later
- OpenWeatherMap API key (create one at https://openweathermap.org/api)

## Setup

1. Clone the repository
2. Create a `.env` file in the root directory with your OpenWeatherMap API key:
   ```
   OPENWEATHER_API_KEY=your_api_key_here
   ```
3. Install dependencies:
   ```bash
   go mod tidy
   ```
4. Run the application:
   ```bash
   go run main.go
   ```

## Go Language Features Demonstrated

### Strengths
- Fast performance and low memory usage
- Built-in concurrency with goroutines and channels
- Strong type system and compile-time safety
- Excellent standard library
- Simple dependency management
- Channel-based communication

### Limitations
- Limited generics (though improved in recent versions)
- More verbose error handling
- Manual JSON marshaling
- Limited object-oriented features 