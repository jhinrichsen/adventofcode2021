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
	return strings.ToLower(cave) == cave
}

// Day12 counts paths where small caves can be visited at most once (part1) or one small cave can be visited twice (part2)
func Day12(connections map[string][]string, part1 bool) uint {
	if part1 {
		return countPaths(connections, "start", "end", false)
	} else {
		return countPaths(connections, "start", "end", true)
	}
}

// countPaths iteratively counts all valid paths from start to end
func countPaths(connections map[string][]string, start, end string, canVisitTwice bool) uint {
	// Stack for DFS traversal, each element contains:
	// - current cave
	// - visited map (copy for this path)
	type stackItem struct {
		cave    string
		visited map[string]int
	}

	stack := []stackItem{{cave: start, visited: make(map[string]int)}}
	count := uint(0)

	for len(stack) > 0 {
		// Pop from stack
		current := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		// Mark current cave as visited
		current.visited[current.cave]++

		// If we've reached the end, we found a valid path
		if current.cave == end {
			count++
			continue
		}

		// Explore neighbors
		for _, neighbor := range connections[current.cave] {
			// Skip start cave (we never go back to start)
			if neighbor == "start" {
				continue
			}

			// For part 1, only visit small caves once
			// For part 2, visit small caves once, except one small cave which can be visited twice
			if !isSmallCave(neighbor) || current.visited[neighbor] == 0 {
				// Create a copy of the visited map for the new path
				visitedCopy := make(map[string]int)
				for k, v := range current.visited {
					visitedCopy[k] = v
				}
				stack = append(stack, stackItem{cave: neighbor, visited: visitedCopy})
			} else if canVisitTwice && current.visited[neighbor] == 1 {
				// For part 2, we can visit this small cave a second time
				// But only if no small cave has been visited twice already
				// For part 2, we can visit this small cave a second time
				// But only if we haven't already visited any small cave twice
				alreadyVisitedTwice := false
				for cave, times := range current.visited {
					if isSmallCave(cave) && times >= 2 {
						alreadyVisitedTwice = true
						break
					}
				}

				// If no small cave has been visited twice, we can visit this one twice
				if !alreadyVisitedTwice {
					// Create a copy of the visited map for the new path
					visitedCopy := make(map[string]int)
					for k, v := range current.visited {
						visitedCopy[k] = v
					}
					stack = append(stack, stackItem{cave: neighbor, visited: visitedCopy})
				}
			}
		}
	}

	return count
}
