package adventofcode2021

import (
	"math"
	"strings"
)

// NewDay14 parses the input lines into template and rules
func NewDay14(lines []string) (string, map[string]string) {
	rules := make(map[string]string)
	template := lines[0]

	for i := 1; i < len(lines); i++ {
		line := lines[i]
		if len(line) == 0 {
			continue
		}

		parts := strings.Split(line, " -> ")
		if len(parts) == 2 {
			rules[parts[0]] = parts[1]
		}
	}

	return template, rules
}

// Day14 solves day 14 part 1 or part 2 depending on the part1 flag
func Day14(template string, rules map[string]string, part1 bool) uint {
	steps := 40
	if part1 {
		steps = 10
	}

	// Count pairs instead of building the full polymer
	pairCounts := make(map[string]uint)
	for i := 0; i < len(template)-1; i++ {
		pair := template[i : i+2]
		pairCounts[pair]++
	}

	// Apply steps
	for range steps {
		newPairCounts := make(map[string]uint)
		for pair, count := range pairCounts {
			if insert, exists := rules[pair]; exists {
				// Create new pairs from the insertion
				newPair1 := string(pair[0]) + insert
				newPair2 := insert + string(pair[1])
				newPairCounts[newPair1] += count
				newPairCounts[newPair2] += count
			} else {
				newPairCounts[pair] += count
			}
		}
		pairCounts = newPairCounts
	}

	// Count elements from pairs
	elementCounts := make(map[byte]uint)
	for pair, count := range pairCounts {
		// Each pair contributes to the count of both its elements
		elementCounts[pair[0]] += count
		elementCounts[pair[1]] += count
	}

	// Add 1 to first and last elements to account for them being counted only once
	firstElement := template[0]
	lastElement := template[len(template)-1]
	elementCounts[firstElement]++
	elementCounts[lastElement]++

	// Now divide all counts by 2
	for element := range elementCounts {
		elementCounts[element] /= 2
	}

	// Find most and least common elements
	var maxCount, minCount uint
	minCount = math.MaxUint
	for _, count := range elementCounts {
		maxCount = max(maxCount, count)
		minCount = min(minCount, count)
	}

	return maxCount - minCount
}

// Day14PolymerAfterSteps returns the polymer string after applying the given number of steps
// This is a helper function for testing the step-by-step progression
func Day14PolymerAfterSteps(template string, rules map[string]string, steps int) string {
	polymer := template
	for step := 0; step < steps; step++ {
		var newPolymer strings.Builder
		for i := 0; i < len(polymer)-1; i++ {
			pair := polymer[i : i+2]
			newPolymer.WriteByte(polymer[i])
			if insert, exists := rules[pair]; exists {
				newPolymer.WriteString(insert)
			}
		}
		newPolymer.WriteByte(polymer[len(polymer)-1])
		polymer = newPolymer.String()
	}
	return polymer
}
