package purchase

// NeedsLicense determines whether a license is needed to drive a type of vehicle. Only "car" and "truck" require a license.
func NeedsLicense(kind string) bool {
	if kind == "car" || kind == "truck" {
		return true
	} else {
		return false
	}
}

// ChooseVehicle recommends a vehicle for selection. It always recommends the vehicle that comes first in lexicographical order.
func ChooseVehicle(option1, option2 string) string {
	var car string

	if option1 < option2 {
		car = option1
	} else {
		car = option2
	}

	message := car + " is clearly the better choice."

	return message
}

// CalculateResellPrice calculates how much a vehicle can resell for at a certain age.
func CalculateResellPrice(originalPrice, age float64) float64 {
	if age < 3 {
		return originalPrice*80/100
	} else if age >= 3 && age < 10 {
		return originalPrice*70/100
	} else {
		return originalPrice*50/100
	}
}
