package cars

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	carsPerHour := float64(productionRate)*successRate/100

	return carsPerHour
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	carsPerHour := float64(productionRate)*successRate/100
	carsPerMinute := carsPerHour/60
	return int(carsPerMinute)
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
	individual := carsCount%10
	carsCost := (carsCount/10*95000) + individual*10000

	return uint(carsCost)
}
