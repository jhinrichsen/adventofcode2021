package adventofcode2021

// NewDay11 parses the input lines into an octopus energy grid
func NewDay11(lines []string) [][]int {
	// Filter out empty lines
	var validLines []string
	for _, line := range lines {
		if len(line) > 0 {
			validLines = append(validLines, line)
		}
	}

	grid := make([][]int, len(validLines))
	for i, line := range validLines {
		grid[i] = make([]int, len(line))
		for j, char := range line {
			grid[i][j] = int(char - '0')
		}
	}
	return grid
}


func Day11(data [][]int, part1 bool) uint {
	rows, cols := len(data), len(data[0])
	totalFlashes := uint(0)

	// Use flat arrays for better cache locality
	gridSize := rows * cols
	grid := make([]int, gridSize)
	flashed := make([]bool, gridSize)

	// Copy data to flat array
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			grid[i*cols+j] = data[i][j]
		}
	}

	// Use integer queue for better performance (store flat indices)
	queue := make([]int, gridSize*10) // Allow for multiple additions of same cell
	flashedCells := make([]int, 0, gridSize) // Track cells that flashed for efficient reset

	// Function to simulate one step
	simulateStep := func() uint {
		// Reset only previously flashed cells
		for _, idx := range flashedCells {
			flashed[idx] = false
		}
		flashedCells = flashedCells[:0] // Clear the list

		queueHead, queueTail := 0, 0

		// Increase energy level and check for initial flashes
		for idx := 0; idx < gridSize; idx++ {
			grid[idx]++
			if grid[idx] > 9 {
				queue[queueTail] = idx
				queueTail++
			}
		}

		stepFlashes := uint(0)

		// Process flashes using efficient queue operations
		for queueHead < queueTail {
			idx := queue[queueHead]
			queueHead++

			// Skip if already flashed
			if flashed[idx] {
				continue
			}

			// Ensure energy level is still above 9
			if grid[idx] > 9 {
				// Flash this octopus
				flashed[idx] = true
				flashedCells = append(flashedCells, idx)
				stepFlashes++

				// Convert flat index to 2D coordinates for bounds checking
				i, j := idx/cols, idx%cols

				// Increase energy of all adjacent octopuses using precomputed flat offsets
				// Top row
				if i > 0 {
					if j > 0 { // Top-left
						nIdx := idx - cols - 1
						grid[nIdx]++
						if grid[nIdx] > 9 && !flashed[nIdx] {
							queue[queueTail] = nIdx
							queueTail++
						}
					}
					// Top
					nIdx := idx - cols
					grid[nIdx]++
					if grid[nIdx] > 9 && !flashed[nIdx] {
						queue[queueTail] = nIdx
						queueTail++
					}
					if j < cols-1 { // Top-right
						nIdx := idx - cols + 1
						grid[nIdx]++
						if grid[nIdx] > 9 && !flashed[nIdx] {
							queue[queueTail] = nIdx
							queueTail++
						}
					}
				}

				// Same row
				if j > 0 { // Left
					nIdx := idx - 1
					grid[nIdx]++
					if grid[nIdx] > 9 && !flashed[nIdx] {
						queue[queueTail] = nIdx
						queueTail++
					}
				}
				if j < cols-1 { // Right
					nIdx := idx + 1
					grid[nIdx]++
					if grid[nIdx] > 9 && !flashed[nIdx] {
						queue[queueTail] = nIdx
						queueTail++
					}
				}

				// Bottom row
				if i < rows-1 {
					if j > 0 { // Bottom-left
						nIdx := idx + cols - 1
						grid[nIdx]++
						if grid[nIdx] > 9 && !flashed[nIdx] {
							queue[queueTail] = nIdx
							queueTail++
						}
					}
					// Bottom
					nIdx := idx + cols
					grid[nIdx]++
					if grid[nIdx] > 9 && !flashed[nIdx] {
						queue[queueTail] = nIdx
						queueTail++
					}
					if j < cols-1 { // Bottom-right
						nIdx := idx + cols + 1
						grid[nIdx]++
						if grid[nIdx] > 9 && !flashed[nIdx] {
							queue[queueTail] = nIdx
							queueTail++
						}
					}
				}
			}
		}

		// Reset all flashed octopuses to 0 (only the ones that actually flashed)
		for _, idx := range flashedCells {
			grid[idx] = 0
		}

		return stepFlashes
	}

	if part1 {
		// Part 1: Count flashes after 100 steps
		for step := 0; step < 100; step++ {
			totalFlashes += simulateStep()
		}
		return totalFlashes
	} else {
		// Part 2: Find the step when all octopuses flash simultaneously
		expectedFlashes := uint(rows * cols)
		for step := uint(1); ; step++ {
			if simulateStep() == expectedFlashes {
				return step
			}
		}
	}
}
