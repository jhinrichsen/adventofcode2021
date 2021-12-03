package adventofcode2021

// Day01Part1 solves day 1 part 1
func Day01Part1(ns []int) int {
	count := 0
	for i := 1; i < len(ns); i++ {
		if ns[i] > ns[i-1] {
			count++
		}
	}
	return count
}

// Day01Part2 solves day 1 part 2
func Day01Part2(ns []int) int {
	count := 0
	for i := 1; i < len(ns)-2; i++ {
		lastWindow := ns[i-1] + ns[i] + ns[i+1]
		window := ns[i] + ns[i+1] + ns[i+2]

		if window > lastWindow {
			count++
		}
	}
	return count
}
