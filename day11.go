package adventofcode2021

// OctopusGrid represents the grid of octopus energy levels
type OctopusGrid [][]int

// NewDay11 parses the input lines into an octopus energy grid
func NewDay11(lines []string) (OctopusGrid, error) {
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
	return OctopusGrid(grid), nil
}

// Directions for adjacent cells (including diagonals)
var octopusDirections = [8][2]int{
	{-1, -1}, {-1, 0}, {-1, 1},
	{0, -1}, {0, 1},
	{1, -1}, {1, 0}, {1, 1},
}

func Day11(data OctopusGrid, part1 bool) uint {
	// Create a copy of the grid to avoid modifying the original
	grid := make([][]int, len(data))
	for i := range data {
		grid[i] = make([]int, len(data[i]))
		copy(grid[i], data[i])
	}

	rows, cols := len(grid), len(grid[0])
	totalFlashes := uint(0)

	// Pre-allocate arrays for better performance
	flashed := make([][]bool, rows)
	for i := range flashed {
		flashed[i] = make([]bool, cols)
	}

	// Use a more efficient queue implementation
	// Queue needs to be larger than grid size due to cascading flashes adding duplicates
	queue := make([][2]int, rows*cols*10)        // Allow for multiple additions of same cell
	flashedCells := make([][2]int, 0, rows*cols) // Track cells that flashed for efficient reset

	// Function to simulate one step
	simulateStep := func() uint {
		// Reset only previously flashed cells (more efficient than full grid scan)
		for _, cell := range flashedCells {
			flashed[cell[0]][cell[1]] = false
		}
		flashedCells = flashedCells[:0] // Clear the list

		queueHead, queueTail := 0, 0

		// Increase energy level and check for initial flashes
		for i := 0; i < rows; i++ {
			for j := 0; j < cols; j++ {
				grid[i][j]++
				if grid[i][j] > 9 && queueTail < len(queue) {
					queue[queueTail] = [2]int{i, j}
					queueTail++
				}
			}
		}

		stepFlashes := uint(0)

		// Process flashes using efficient queue operations
		for queueHead < queueTail {
			pos := queue[queueHead]
			queueHead++
			i, j := pos[0], pos[1]

			// Skip if already flashed
			if flashed[i][j] {
				continue
			}

			// Ensure energy level is still above 9
			if grid[i][j] > 9 {
				// Flash this octopus
				flashed[i][j] = true
				flashedCells = append(flashedCells, [2]int{i, j})
				stepFlashes++

				// Increase energy of all adjacent octopuses
				for _, dir := range octopusDirections {
					ni, nj := i+dir[0], j+dir[1]
					if ni >= 0 && ni < rows && nj >= 0 && nj < cols {
						grid[ni][nj]++
						// Add to queue if this causes a new flash
						if grid[ni][nj] > 9 && !flashed[ni][nj] && queueTail < len(queue) {
							queue[queueTail] = [2]int{ni, nj}
							queueTail++
						}
					}
				}
			}
		}

		// Reset all flashed octopuses to 0 (only the ones that actually flashed)
		for _, cell := range flashedCells {
			grid[cell[0]][cell[1]] = 0
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
