package adventofcode2021

import "strings"

// 0: 6 abcefg
// 1: 2 cf
// 2: 5 acdeg
// 3: 5 acdfg
// 4: 4 bcdf
// 5: 5 abdfg
// 6: 6 abdefg
// 7: 3 acf
// 8: 7 abcdefg
// 9: 6 abcdfg
func Day08(lines []string, part1 bool) uint {
	var n uint
	for _, line := range lines {
		parts := strings.Split(line, "|")
		outs := strings.Fields(parts[1])
		for _, out := range outs {
			switch len(out) {
			case 2: // 1
				n++
			case 4: // 4
				n++
			case 3: // 7
				n++
			case 7: // 8
				n++
			}
		}
	}
	return n
}
