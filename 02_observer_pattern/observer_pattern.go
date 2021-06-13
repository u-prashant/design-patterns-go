package main

import "design-patterns-go/02_observer_pattern/weather"

func main() {
	weatherData := &weather.WeatherData{}
	currentConditionsDisplay := weather.NewCurrentConditionsDisplay(weatherData)
	weatherData.SetMeasurements(80, 65, 30)
	currentConditionsDisplay.DeRegister()
	weatherData.SetMeasurements(80, 70, 30)
}
