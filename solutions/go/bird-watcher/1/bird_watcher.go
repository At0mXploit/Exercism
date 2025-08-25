package birdwatcher

// TotalBirdCount returns the total bird count for all days.
func TotalBirdCount(birdsPerDay []int) int {
	total := 0
	for i := 0; i < len(birdsPerDay); i++ {
		total += birdsPerDay[i]
	}
	return total
}

// BirdsInWeek returns the total bird count for a specific week.
// Weeks are 1-based: week 1 = days 0-6, week 2 = days 7-13, etc.
func BirdsInWeek(birdsPerDay []int, week int) int {
	start := (week - 1) * 7
	end := start + 7

	total := 0
	for i := start; i < end; i++ {
		total += birdsPerDay[i]
	}
	return total
}

// FixBirdCountLog returns the corrected bird counts
// by adding 1 to every other day starting from the first day.
func FixBirdCountLog(birdsPerDay []int) []int {
    for i := 0; i < len(birdsPerDay); i += 2 {
        birdsPerDay[i] += 1
    }
    return birdsPerDay
}

