package adventofcode2021

import (
	"strconv"
	"strings"
)

// Day02 solves both parts of day 2.
func Day02(lines []string, part1 bool) (int, error) {
	var position, depth, aim, n int
	var parts []string

	fn1 := func() {
		switch parts[0] {
		case "forward":
			position += n
		case "up":
			depth -= n
		case "down":
			depth += n
		}
	}
	fn2 := func() {
		switch parts[0] {
		case "forward":
			position += n
			depth += aim * n
		case "up":
			aim -= n
		case "down":
			aim += n
		}
	}

	for _, line := range lines {
		var err error
		parts = strings.Fields(line)
		n, err = strconv.Atoi(parts[1])
		if err != nil {
			return 0, err
		}
		if part1 {
			fn1()
		} else {
			fn2()
		}
	}
	return position * depth, nil
}
