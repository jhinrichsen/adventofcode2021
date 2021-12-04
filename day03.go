package adventofcode2021

import (
	"fmt"
	"strconv"
)

func Day03(lines []string, part1 bool) (int, error) {
	if part1 {
		return day03Part1(lines)
	}
	return day03Part2(lines)
}

func day03Part1(lines []string) (int, error) {
	bits := len(lines[0])

	countBits := func(bit int) (zeroes, ones int) {
		for _, line := range lines {
			if line[bit] == '0' {
				zeroes++
			} else {
				ones++
			}
		}
		return
	}

	// simulate Rust's multitypes by hungarian notation
	var sgamma, sepsilon string

	for bit := 0; bit < bits; bit++ {
		zeroes, ones := countBits(bit)
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

func day03Part2(lines []string) (int, error) {
	// map reduce
	var m map[string]bool

	firstEntry := func() (int64, error) {
		for k := range m {
			return strconv.ParseInt(k, 2, 64)
		}
		return 0, fmt.Errorf("internal error: map is empty")
	}

	// del removes entry from map unless one entry left.
	del := func(bit int, value byte) {
		if len(m) == 1 {
			return
		}
		for k := range m {
			if k[bit] == value {
				delete(m, k)
			}
		}
	}

	countBits := func(bit int) (zeroes, ones int) {
		for k := range m {
			if k[bit] == '0' {
				zeroes++
			} else {
				ones++
			}
		}
		return
	}

	rating := func(lines []string, dominant, recessive byte) (int64, error) {
		m = make(map[string]bool)
		for _, line := range lines {
			m[line] = true
		}
		bits := len(lines[0])
		for bit := 0; bit < bits; bit++ {
			zeroes, ones := countBits(bit)
			if zeroes > ones {
				del(bit, dominant)
			} else if ones > zeroes {
				del(bit, recessive)
			} else { // zeroes == ones
				del(bit, recessive)
			}
			if len(m) == 1 {
				return firstEntry()
			}
		}
		return 0, fmt.Errorf("want 1 entry in map but got %d: %+v\n",
			len(m), m)
	}

	oxygenGeneratorRating, err := rating(lines, '1', '0')
	if err != nil {
		return 0, err
	}
	CO2ScrubberGeneratorRating, err := rating(lines, '0', '1')
	if err != nil {
		return 0, err
	}

	return int(oxygenGeneratorRating * CO2ScrubberGeneratorRating), nil
}
