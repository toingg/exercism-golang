package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
	var totalBirds int
	for i := 0; i < len(birdsPerDay); i++ {
		totalBirds += birdsPerDay[i]
	}
	return totalBirds
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
	// birdsWeek := birdsPerDay[(week-1)*7:week*7]
	// var totalBirds int
	// for i := 0; i < len(birdsWeek); i++ {
	// 	totalBirds += birdsWeek[i]
	// }
	// return totalBirds

	birdsWeek := birdsPerDay[(week-1)*7 : week*7]

	return TotalBirdCount(birdsWeek)

}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {

	for i := 0; i < len(birdsPerDay); i++ {
		if i%2 == 0 {
			birdsPerDay[i] = birdsPerDay[i] + 1
		}
	}
	// alt
	// for i := 0; i< len(birdsPerDay); i+=2 {
	// 	birdsPerDay[i]++
	// }
	return birdsPerDay
}
