package adventofcode2021

// OctopusGrid represents the grid of octopus energy levels
type OctopusGrid [][]int

// NewDay11 parses the input lines into an octopus energy grid
func NewDay11(lines []string) (OctopusGrid, error) {
	grid := make([][]int, len(lines))
	for i, line := range lines {
		grid[i] = make([]int, len(line))
		for j, char := range line {
			grid[i][j] = int(char - '0')
		}
	}
	return OctopusGrid(grid), nil
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

	// Directions for adjacent cells (including diagonals)
	directions := [][]int{
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1},           {0, 1},
		{1, -1},  {1, 0},  {1, 1},
	}

	// Function to simulate one step
	simulateStep := func() uint {
		// First, increase energy level of all octopuses by 1
		for i := 0; i < rows; i++ {
			for j := 0; j < cols; j++ {
				grid[i][j]++
			}
		}

		// Track which octopuses have flashed this step
		flashed := make([][]bool, rows)
		for i := range flashed {
			flashed[i] = make([]bool, cols)
		}

		stepFlashes := uint(0)
		changed := true

		// Keep flashing until no more flashes occur
		for changed {
			changed = false
			for i := 0; i < rows; i++ {
				for j := 0; j < cols; j++ {
					if grid[i][j] > 9 && !flashed[i][j] {
						// Flash this octopus
						flashed[i][j] = true
						stepFlashes++
						changed = true

						// Increase energy of all adjacent octopuses
						for _, dir := range directions {
							ni, nj := i+dir[0], j+dir[1]
							if ni >= 0 && ni < rows && nj >= 0 && nj < cols {
								grid[ni][nj]++
							}
						}
					}
				}
			}
		}

		// Reset all flashed octopuses to 0
		for i := 0; i < rows; i++ {
			for j := 0; j < cols; j++ {
				if flashed[i][j] {
					grid[i][j] = 0
				}
			}
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
		step := uint(0)
		for {
			step++
			flashes := simulateStep()
			if flashes == uint(rows*cols) {
				return step
			}
		}
	}
}
