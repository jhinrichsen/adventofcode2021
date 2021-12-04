package adventofcode2021

import (
	"strconv"
)

func Day03(lines []string, part1 bool) (int, error) {
	bits := len(lines[0])

	// simulate Rust's multitypes by hungarian notation
	var sgamma, sepsilon string

	for bit := 0; bit < bits; bit++ {
		var zeroes, ones int
		for _, line := range lines {
			if line[bit] == '0' {
				zeroes++
			} else {
				ones++
			}
		}
		if zeroes > ones {
			sgamma += "0"
			sepsilon += "1"
		} else {
			sgamma += "1"
			sepsilon += "0"
		}
	}
	gamma, err := strconv.ParseInt(sgamma, 2, 64)
	if err != nil {
		return 0, err
	}
	epsilon, err := strconv.ParseInt(sepsilon, 2, 64)
	if err != nil {
		return 0, err
	}
	powerConsumption := int(epsilon * gamma)
	return powerConsumption, nil
}
