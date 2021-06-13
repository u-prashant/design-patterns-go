package weather

type Subject interface {
	registerObserver(observer Observer)
	removeObserver(observer Observer)
	notifyObservers()
}

type WeatherData struct {
	observers   []Observer
	temperature float64
	humidity    float64
	pressure    float64
}

func (w *WeatherData) registerObserver(observer Observer) {
	w.observers = append(w.observers, observer)
}

func (w *WeatherData) removeObserver(observer Observer) {
	for i, o := range w.observers {
		if o == observer {
			w.observers = append(w.observers[:i], w.observers[i+1:]...)
		}
	}
}

func (w *WeatherData) notifyObservers() {
	for _, observer := range w.observers {
		observer.update()
	}
}

func (w *WeatherData) measurementsChanged() {
	w.notifyObservers()
}

func (w *WeatherData) GetTemperature() float64 {
	return w.temperature
}

func (w *WeatherData) GetHumidity() float64 {
	return w.humidity
}

func (w *WeatherData) GetPressure() float64 {
	return w.pressure
}

func (w *WeatherData) SetMeasurements(temperature float64, humidity float64, pressure float64) {
	w.temperature = temperature
	w.humidity = humidity
	w.pressure = pressure
	w.measurementsChanged()
}
