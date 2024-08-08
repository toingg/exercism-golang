// Package weather adalah package untuk memprediksi cuaca di kota Goblinius.
package weather

// CurrentCondition berisi kondisi cuaca sekarang.
var CurrentCondition string
// CurrentLocation berisi lokasi saat ini.
var CurrentLocation string

// Forecast adalah fungsi public yang digunakan untuk mengecek kondisi cuaca sekarang.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}


