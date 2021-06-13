package weather

import "fmt"

type Observer interface {
	update()
}

type DisplayElement interface {
	display()
}

type CurrentConditionsDisplay struct {
	temperature float64
	humidity    float64
	weatherData *WeatherData
}

func (c *CurrentConditionsDisplay) display() {
	fmt.Printf("Current Conditions: temperature %f F degrees and humidity %f\n", c.temperature, c.humidity)
}

func (c *CurrentConditionsDisplay) update() {
	c.temperature = c.weatherData.GetTemperature()
	c.humidity = c.weatherData.GetHumidity()
	c.display()
}

func (c *CurrentConditionsDisplay) DeRegister() {
	c.weatherData.removeObserver(c)
}

func NewCurrentConditionsDisplay(weatherData *WeatherData) *CurrentConditionsDisplay {
	c := CurrentConditionsDisplay{
		weatherData: weatherData,
	}
	c.weatherData.registerObserver(&c)
	return &c
}
