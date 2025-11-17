package adventofcode2021

import (
	"strings"
)

// NewDay12 creates a new cave system from input lines
func NewDay12(lines []string) map[string][]string {
	connections := make(map[string][]string)

	for _, line := range lines {
		if line == "" {
			continue
		}
		parts := strings.Split(line, "-")
		if len(parts) != 2 {
			continue
		}

		from, to := parts[0], parts[1]

		// Add bidirectional connections
		connections[from] = append(connections[from], to)
		connections[to] = append(connections[to], from)
	}

	return connections
}

// isSmallCave checks if a cave name represents a small cave (lowercase)
func isSmallCave(cave string) bool {
	if len(cave) == 0 {
		return false
	}
	// Check if all characters are lowercase (no allocation)
	for i := range len(cave) {
		if cave[i] < 'a' || cave[i] > 'z' {
			return false
		}
	}
	return true
}

// Day12 counts paths where small caves can be visited at most once (part1) or one small cave can be visited twice (part2)
func Day12(connections map[string][]string, part1 bool) uint {
	if part1 {
		return countPaths(connections, "start", "end", false)
	} else {
		return countPaths(connections, "start", "end", true)
	}
}

// countPaths iteratively counts all valid paths from start to end using backtracking
func countPaths(connections map[string][]string, start, end string, canVisitTwice bool) uint {
	type stackItem struct {
		cave            string
		usedDoubleVisit bool
		neighborIdx     int  // Index of next neighbor to explore
		entering        bool // true when entering, false when backtracking
	}

	visited := make(map[string]int, 16)
	stack := make([]stackItem, 1, 128)
	stack[0] = stackItem{cave: start, usedDoubleVisit: false, neighborIdx: 0, entering: true}
	var count uint

	for len(stack) > 0 {
		current := &stack[len(stack)-1]

		if current.entering {
			// Entering this cave
			visited[current.cave]++

			// Check if we reached the end
			if current.cave == end {
				count++
				visited[current.cave]--
				stack = stack[:len(stack)-1]
				continue
			}

			// Mark as no longer entering (next iteration will explore neighbors)
			current.entering = false
		}

		neighbors := connections[current.cave]
		if current.neighborIdx < len(neighbors) {
			// Explore next neighbor
			neighbor := neighbors[current.neighborIdx]
			current.neighborIdx++

			// Skip start cave (never return to start)
			if neighbor == "start" {
				continue
			}

			visitCount := visited[neighbor]

			// Large caves or unvisited small caves can always be visited
			if !isSmallCave(neighbor) || visitCount == 0 {
				stack = append(stack, stackItem{
					cave:            neighbor,
					usedDoubleVisit: current.usedDoubleVisit,
					neighborIdx:     0,
					entering:        true,
				})
			} else if canVisitTwice && visitCount == 1 && !current.usedDoubleVisit {
				// For part 2: visit this small cave twice (first double visit)
				stack = append(stack, stackItem{
					cave:            neighbor,
					usedDoubleVisit: true,
					neighborIdx:     0,
					entering:        true,
				})
			}
		} else {
			// Done exploring all neighbors, backtrack
			visited[current.cave]--
			stack = stack[:len(stack)-1]
		}
	}

	return count
}
