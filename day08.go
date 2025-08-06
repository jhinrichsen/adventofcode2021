package adventofcode2021

import (
	"sort"
	"strings"
)

func sortString(s string) string {
	r := []rune(s)
	sort.Slice(r, func(i, j int) bool { return r[i] < r[j] })
	return string(r)
}

func containsAll(s, chars string) bool {
	for _, c := range chars {
		if !strings.ContainsRune(s, c) {
			return false
		}
	}
	return true
}

func Day08(lines []string, part1 bool) uint {
	var total uint

	for _, line := range lines {
		parts := strings.Split(line, " | ")
		patterns := strings.Fields(parts[0])
		outputs := strings.Fields(parts[1])

		// Sort all patterns and outputs for consistency
		for i, p := range patterns {
			patterns[i] = sortString(p)
		}
		for i, o := range outputs {
			outputs[i] = sortString(o)
		}

		// Map to store the digit to its pattern
		digitToPattern := make(map[int]string)

		// First pass: identify unique length digits
		var fives, sixes []string
		for _, p := range patterns {
			switch len(p) {
			case 2:
				digitToPattern[1] = p
			case 3:
				digitToPattern[7] = p
			case 4:
				digitToPattern[4] = p
			case 5:
				fives = append(fives, p)
			case 6:
				sixes = append(sixes, p)
			case 7:
				digitToPattern[8] = p
			}
		}

		// Identify 3 (the only 5-segment digit that contains all segments of 1)
		for i, p := range fives {
			if containsAll(p, digitToPattern[1]) {
				digitToPattern[3] = p
				fives = append(fives[:i], fives[i+1:]...)
				break
			}
		}

		// Identify 9 (the only 6-segment digit that contains all segments of 3)
		for i, p := range sixes {
			if containsAll(p, digitToPattern[3]) {
				digitToPattern[9] = p
				sixes = append(sixes[:i], sixes[i+1:]...)
				break
			}
		}

		// Identify 0 (the remaining 6-segment digit that contains all segments of 1)
		for i, p := range sixes {
			if containsAll(p, digitToPattern[1]) {
				digitToPattern[0] = p
				sixes = append(sixes[:i], sixes[i+1:]...)
				break
			}
		}

		// The remaining 6-segment digit is 6
		digitToPattern[6] = sixes[0]

		// Identify 5 (the remaining 5-segment digit that is contained in 6)
		for i, p := range fives {
			if containsAll(digitToPattern[6], p) {
				digitToPattern[5] = p
				fives = append(fives[:i], fives[i+1:]...)
				break
			}
		}

		// The remaining 5-segment digits are 2 and 3
		// 3 contains all segments of 1, 2 doesn't
		for _, p := range fives {
			if containsAll(p, digitToPattern[1]) {
				digitToPattern[3] = p
			} else {
				digitToPattern[2] = p
			}
		}

		// Create a reverse mapping from pattern to digit
		patternToDigit := make(map[string]int)
		for d, p := range digitToPattern {
			patternToDigit[p] = d
		}

		if part1 {
			// For part 1, just count 1, 4, 7, 8 in outputs
			for _, o := range outputs {
				switch len(o) {
				case 2, 3, 4, 7:
					total++
				}
			}
		} else {
			// For part 2, decode the output value
			var outputValue int
			for _, o := range outputs {
				digit := patternToDigit[o]
				// Calculate the place value (thousands, hundreds, tens, ones)
				outputValue = outputValue*10 + digit
			}
			total += uint(outputValue)
		}
	}

	return total
}
