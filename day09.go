package adventofcode2021

import "sort"

func Day09(lines []string, part1 bool) (uint, error) {
	Y := len(lines)
	X := len(lines[0])

	height := func(x, y int) byte {
		return lines[y][x] - '0'
	}

	// Find low points
	var lowPoints [][2]int
	for y := range lines {
		for x := range lines[y] {
			h := height(x, y)

			low := true
			// N
			if y > 0 && height(x, y-1) <= h {
				low = false
			}
			// E
			if x < X-1 && height(x+1, y) <= h {
				low = false
			}
			// S
			if y < Y-1 && height(x, y+1) <= h {
				low = false
			}
			// W
			if x > 0 && height(x-1, y) <= h {
				low = false
			}
			if low {
				lowPoints = append(lowPoints, [2]int{x, y})
			}
		}
	}

	if part1 {
		var risk uint
		for _, point := range lowPoints {
			x, y := point[0], point[1]
			risk += uint(height(x, y)) + 1
		}
		return risk, nil
	}

	// Part 2: Find basin sizes
	visited := make([][]bool, Y)
	for i := range visited {
		visited[i] = make([]bool, X)
	}

	var basinSizes []int
	for _, point := range lowPoints {
		x, y := point[0], point[1]
		size := exploreBasin(lines, visited, x, y, X, Y)
		basinSizes = append(basinSizes, size)
	}

	// Sort basin sizes in descending order
	sort.Sort(sort.Reverse(sort.IntSlice(basinSizes)))

	// Multiply the three largest
	result := uint(basinSizes[0] * basinSizes[1] * basinSizes[2])
	return result, nil
}

func exploreBasin(lines []string, visited [][]bool, startX, startY, X, Y int) int {
	// Use iterative approach with a stack to avoid recursion
	stack := [][2]int{{startX, startY}}
	size := 0

	for len(stack) > 0 {
		// Pop from stack
		current := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		x, y := current[0], current[1]

		// Check bounds and if already visited
		if x < 0 || x >= X || y < 0 || y >= Y || visited[y][x] {
			continue
		}

		// Check if height is 9 (not part of any basin)
		height := lines[y][x] - '0'
		if height == 9 {
			continue
		}

		// Mark as visited and increment size
		visited[y][x] = true
		size++

		// Add all 4 neighbors to stack
		stack = append(stack, [2]int{x - 1, y}) // West
		stack = append(stack, [2]int{x + 1, y}) // East
		stack = append(stack, [2]int{x, y - 1}) // North
		stack = append(stack, [2]int{x, y + 1}) // South
	}

	return size
}
