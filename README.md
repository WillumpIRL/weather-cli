# Weather CLI

A command-line interface application that provides real-time weather information for cities worldwide using the OpenWeatherMap API.

## Features

- Real-time weather data retrieval
- Support for cities worldwide
- Beautiful emoji-enhanced display
- Temperature in Celsius
- Humidity information
- Weather condition descriptions
- Local time for the queried city

## Prerequisites

- Go 1.24 or higher
- OpenWeatherMap API key

## Installation

1. Clone the repository:
```bash
git clone https://github.com/yourusername/weather-cli.git
cd weather-cli
```

2. Create a copy of the example environment file and add your OpenWeatherMap API key:
```bash
cp .env.example .env
# Edit .env and add your API key
```

3. Build the application:
```bash
go build -o weather ./cmd/weather
```

## Usage

Run the application:
```bash
./weather
```

Enter a city name when prompted. The application will display:
- City name
- Local time
- Weather condition
- Temperature in Celsius
- Humidity percentage

Type 'x' to exit the application.

## Example Output
```
[------------------Weather Information--------------]
 🌍 City: London
 🕒 Local Time: 2024-03-27 04:14:47
 🌤️  Condition: few clouds
 🌡️  Temperature: 5.22°C
 💧 Humidity: 90%
[---------------------------------------------------]
```

## Environment Variables

- `OPENWEATHER_API_KEY`: Your OpenWeatherMap API key

## Project Structure

```
.
├── cmd/
│   └── weather/
│       └── main.go
├── internal/
│   ├── api/
│   │   └── openweathermap.go
│   └── models/
│       └── weather.go
├── .env.example
├── .gitignore
├── go.mod
└── README.md
```

## License

MIT License - feel free to use this project as you wish.

## Contributing

Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change. 