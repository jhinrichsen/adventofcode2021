package adventofcode2021

type SeaCucumberGrid struct {
	grid  [][]byte
	dimX  int
	dimY  int
}

func parseDay25(lines []string) SeaCucumberGrid {
	var grid [][]byte

	for _, line := range lines {
		if len(line) > 0 {
			row := make([]byte, len(line))
			copy(row, []byte(line))
			grid = append(grid, row)
		}
	}

	dimY := len(grid)
	dimX := 0
	if dimY > 0 {
		dimX = len(grid[0])
	}

	return SeaCucumberGrid{grid, dimX, dimY}
}

func (g *SeaCucumberGrid) step() bool {
	moved := false

	// Phase 1: Move east-facing cucumbers
	newGrid := make([][]byte, g.dimY)
	for y := range g.dimY {
		newGrid[y] = make([]byte, g.dimX)
		copy(newGrid[y], g.grid[y])
	}

	for y := range g.dimY {
		for x := range g.dimX {
			if g.grid[y][x] == '>' {
				nextX := (x + 1) % g.dimX
				if g.grid[y][nextX] == '.' {
					newGrid[y][x] = '.'
					newGrid[y][nextX] = '>'
					moved = true
				}
			}
		}
	}

	g.grid = newGrid

	// Phase 2: Move south-facing cucumbers
	newGrid = make([][]byte, g.dimY)
	for y := range g.dimY {
		newGrid[y] = make([]byte, g.dimX)
		copy(newGrid[y], g.grid[y])
	}

	for y := range g.dimY {
		for x := range g.dimX {
			if g.grid[y][x] == 'v' {
				nextY := (y + 1) % g.dimY
				if g.grid[nextY][x] == '.' {
					newGrid[y][x] = '.'
					newGrid[nextY][x] = 'v'
					moved = true
				}
			}
		}
	}

	g.grid = newGrid

	return moved
}

// Day25 solves day 25 puzzle
func Day25(lines []string, part1 bool) uint {
	grid := parseDay25(lines)

	if part1 {
		steps := uint(0)
		for {
			steps++
			if !grid.step() {
				return steps
			}
		}
	}

	// Part 2 doesn't exist for Day 25 (it's a freebie after completing 1-24)
	return 0
}
