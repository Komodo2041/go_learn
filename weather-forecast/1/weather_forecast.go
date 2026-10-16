// Package weather <package_description>.
package weather
 
 
var (
    // CurrentCondition <variable_description>.
	CurrentCondition string
    //CurrentLocation <variable_description>.
	CurrentLocation  string
)
// Forecast returns an string.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
