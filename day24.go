package adventofcode2021

import (
	"strconv"
	"strings"
)

type MonadBlock struct {
	divz int
	addx int
	addy int
}

func parseDay24(lines []string) []MonadBlock {
	var blocks []MonadBlock
	var current MonadBlock
	lineNum := 0

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		parts := strings.Fields(line)
		if len(parts) == 0 {
			continue
		}

		if parts[0] == "inp" {
			if lineNum > 0 {
				blocks = append(blocks, current)
			}
			current = MonadBlock{}
			lineNum = 0
			continue
		}

		lineNum++

		// Line 5: div z <divz>
		if lineNum == 4 && parts[0] == "div" && parts[1] == "z" {
			val, err := strconv.Atoi(parts[2])
			if err == nil {
				current.divz = val
			}
		}

		// Line 6: add x <addx>
		if lineNum == 5 && parts[0] == "add" && parts[1] == "x" {
			val, err := strconv.Atoi(parts[2])
			if err == nil {
				current.addx = val
			}
		}

		// Line 16: add y <addy>
		if lineNum == 15 && parts[0] == "add" && parts[1] == "y" {
			val, err := strconv.Atoi(parts[2])
			if err == nil {
				current.addy = val
			}
		}
	}

	if lineNum > 0 {
		blocks = append(blocks, current)
	}

	return blocks
}

// solveMonad finds the largest (part1=true) or smallest (part1=false) valid model number
func solveMonad(blocks []MonadBlock, part1 bool) uint {
	// Build constraints from push/pop pairs
	type constraint struct {
		pos1  int // position that pushes
		pos2  int // position that pops
		delta int // pos2 must equal pos1 + delta
	}

	var constraints []constraint
	var stack []struct {
		pos   int
		addy  int
	}

	// Analyze blocks to find push/pop pairs
	for i, block := range blocks {
		if block.divz == 1 {
			// Push: save position and addy value
			stack = append(stack, struct {
				pos  int
				addy int
			}{i, block.addy})
		} else {
			// Pop: create constraint between this and previous push
			if len(stack) > 0 {
				top := stack[len(stack)-1]
				stack = stack[:len(stack)-1]

				// Constraint: digits[i] = digits[top.pos] + top.addy + block.addx
				constraints = append(constraints, constraint{
					pos1:  top.pos,
					pos2:  i,
					delta: top.addy + block.addx,
				})
			}
		}
	}

	// Solve constraints to find valid digits
	digits := make([]int, 14)

	// Initialize to maximize (part1) or minimize (part2)
	if part1 {
		for i := range digits {
			digits[i] = 9
		}
	} else {
		for i := range digits {
			digits[i] = 1
		}
	}

	// Apply constraints
	for _, c := range constraints {
		// digits[c.pos2] = digits[c.pos1] + c.delta
		// Both must be in range 1-9

		if part1 {
			// Maximize both digits
			d1 := 9
			d2 := d1 + c.delta
			if d2 > 9 {
				d2 = 9
				d1 = d2 - c.delta
			}
			if d1 < 1 {
				d1 = 1
				d2 = d1 + c.delta
			}
			digits[c.pos1] = d1
			digits[c.pos2] = d2
		} else {
			// Minimize both digits
			d1 := 1
			d2 := d1 + c.delta
			if d2 < 1 {
				d2 = 1
				d1 = d2 - c.delta
			}
			if d1 > 9 {
				d1 = 9
				d2 = d1 + c.delta
			}
			digits[c.pos1] = d1
			digits[c.pos2] = d2
		}
	}

	// Convert digits to number
	result := uint(0)
	for _, d := range digits {
		result = result*10 + uint(d)
	}

	return result
}

// Day24 solves day 24 puzzle
func Day24(lines []string, part1 bool) uint {
	blocks := parseDay24(lines)
	return solveMonad(blocks, part1)
}
