package adventofcode2021

import (
	"math"
	"strings"
)

func Day05Int(lines []string, part1 bool) (int, error) {
	parse := func(line string) (int, int, error) {
		numbers, err := ParseCommaSeparatedNumbers(line)
		if err != nil {
			return 0, 0, err
		}
		return numbers[0], numbers[1], nil
	}

	abs := func(i int) int {
		if i < 0 {
			return -i
		}
		return i
	}

	// Use a flat array instead of map - coordinates are bounded by ~1000
	const maxCoord = 1000
	grid := make([]byte, maxCoord*maxCoord)

	for _, line := range lines {
		parts := strings.Fields(line)
		srcX, srcY, err := parse(parts[0])
		if err != nil {
			return 0, err
		}
		dstX, dstY, err := parse(parts[2])
		if err != nil {
			return 0, err
		}

		// for part 1, only consider horizontal or vertical lines
		if part1 && srcX != dstX && srcY != dstY {
			continue
		}

		dx := dstX - srcX
		if dx != 0 {
			dx = dx / abs(dx)
		}
		dy := dstY - srcY
		if dy != 0 {
			dy = dy / abs(dy)
		}

		x, y := srcX, srcY
		for {
			grid[y*maxCoord+x]++
			if x == dstX && y == dstY {
				break
			}
			x += dx
			y += dy
		}
	}

	var count int
	for _, v := range grid {
		if v > 1 {
			count++
		}
	}
	return count, nil
}

func Day05Cmplx(lines []string, part1 bool) (int, error) {
	parse := func(line string) (complex128, error) {
		parts, err := ParseCommaSeparatedNumbers(line)
		if err != nil {
			return 0, err
		}
		return complex(float64(parts[0]), float64(parts[1])), nil
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
