// Package weather declares project scope.
package weather

// CurrentCondition delcares variable current condition for pulled coordinates.
var CurrentCondition string

// CurrentLocation delcares variable location for pulled coordinates.
var CurrentLocation string

// Forecast function returns a string of currentLocation and the currentCondition (Ex. Shawnee, OK - current weather condition: Rainy).
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
