package adventofcode2021

import "math"

func Day07(positions []int, part1 bool) int {
	min := positions[0]
	max := positions[0]
	for i := 1; i < len(positions); i++ {
		if positions[i] < min {
			min = positions[i]
		}
		if positions[i] > max {
			max = positions[i]
		}
	}

	burn := func(dist int) int {
		if part1 {
			return dist
		}
		// OEIS A000217 triangular number
		return dist * (dist + 1) / 2
	}

	fuel := func(pos int) int {
		sum := 0
		for i := range positions {
			dist := positions[i] - pos
			if dist < 0 {
				dist = -dist
			}
			sum += burn(dist)
		}
		return sum
	}

	fmin := math.MaxInt32
	for pos := min; pos <= max; pos++ {
		f := fuel(pos)
		if f < fmin {
			fmin = f
		}
	}
	return fmin
}
