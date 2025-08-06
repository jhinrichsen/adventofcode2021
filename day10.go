package adventofcode2021

import (
	"sort"
	"strings"
)

func Day10(lines []string, part1 bool) uint {
	pairs := []string{"[]", "{}", "()", "<>"}
	points := []uint{57, 1197, 3, 25137}

	// corrupted points or 0 if just incomplete
	corrupted := func(s string) (total uint) {
		for i := range s {
			for j := range pairs {
				if s[i] == pairs[j][1] {
					return points[j]
				}
			}
		}
		return 0
	}

	var incompletes []string
	var total uint
	for _, line := range lines {
		for {
			changed := false
			for j := 0; j < len(points); j++ {
				before := len(line)
				line = strings.Replace(line, pairs[j], "", 1)
				if len(line) != before {
					changed = true
					break
				}
			}
			if !changed {
				break
			}
		}
		// Check if incomplete or corrupted
		n := corrupted(line)
		if n == 0 {
			incompletes = append(incompletes, line)
		} else {
			total += n
		}
	}
	if part1 {
		return total
	}

	// part 2
	totals := make([]uint, len(incompletes))
	for i, line := range incompletes {
		total = 0
		// We need to iterate from back to front when counting opening instead
		// of closing pairs, the order is important because of constant
		// multiplication factor
		// No C style backward loop as in `for j := len(line); --j >=0; `
		for j := len(line) - 1; j >= 0; j-- {
			b := line[j]
			total *= 5
			switch b {
			case '(':
				total += 1
			case '[':
				total += 2
			case '{':
				total += 3
			default:
				total += 4
			}
		}
		totals[i] = total
	}

	// we have sort.IntSlice, but no sort.UIntSlice
	sort.Slice(totals, func(i, j int) bool { return totals[i] < totals[j] })

	return totals[len(totals)/2] // middle element
}
