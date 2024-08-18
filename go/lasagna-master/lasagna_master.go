package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, avgTimePerLayer int) int {
	if avgTimePerLayer == 0 {
		return len(layers) * 2
	}
	return len(layers) * avgTimePerLayer
}

// TODO: define the 'Quantities()' function
func Quantities(layers []string) (int, float64) {
	var noodle, sauce int

	for i := 0; i < len(layers); i++ {
		if layers[i] == "noodles" {
			noodle++
		} else if layers[i] == "sauce" {
			sauce++
		}
	}

	return noodle * 50, float64(sauce) * 0.2

}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendRecipe []string, ownRecipe []string) {
	ownRecipe[len(ownRecipe)-1] = friendRecipe[len(friendRecipe)-1]
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(quantities []float64, portions int) []float64 {
	// newQuantities := []float64{}
    // newQuantities = append(newQuantities, quantities...)

    // best perfomance
    newQuantities := make([]float64, len(quantities))
	for i := 0; i < len(quantities); i++ {
		newQuantities[i] = quantities[i] * float64(portions) / 2
	}
	return newQuantities
}
