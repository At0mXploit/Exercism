// Package weather provides tools to forecast and track the current weather conditions
// for cities in Goblinocus.
package weather

// CurrentCondition stores the current weather condition for a city, e.g., "sunny", "rainy".
var CurrentCondition string

// CurrentLocation stores the name of the city for which the weather condition is recorded.
var CurrentLocation string

// Forecast sets the current location and weather condition, then returns a formatted string
// describing the current weather for the specified city.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
