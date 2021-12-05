package adventofcode2021

import (
	"math"
	"strconv"
	"strings"
)

func Day05(lines []string, part1 bool) (int, error) {
	parse := func(line string) (complex128, error) {
		parts := strings.Split(line, ",")
		x, err := strconv.Atoi(parts[0])
		if err != nil {
			return 0, err
		}
		y, err := strconv.Atoi(parts[1])
		if err != nil {
			return 0, err
		}
		return complex(float64(x), float64(y)), nil
	}

	diagonal := func(c1, c2 complex128) bool {
		return real(c1) != real(c2) && imag(c1) != imag(c2)
	}

	m := make(map[complex128]int)
	for _, line := range lines {
		parts := strings.Fields(line)
		src, err := parse(parts[0])
		if err != nil {
			return 0, err
		}
		dst, err := parse(parts[2])
		if err != nil {
			return 0, err
		}

		// for part 1, only consider horizontal or vertical lines
		if part1 && diagonal(src, dst) {
			continue
		}

		dx := real(dst) - real(src)
		if dx != 0 {
			dx /= math.Abs(real(dst) - real(src))
		}
		dy := imag(dst) - imag(src)
		if dy != 0 {
			dy /= math.Abs(imag(dst) - imag(src))
		}
		inc := complex(dx, dy)

		c := src
		for {
			m[c]++
			if c == dst {
				break
			}
			c += inc
		}
	}

	var count int
	for _, v := range m {
		if v > 1 {
			count++
		}
	}
	return count, nil
}
