package adventofcode2021

import (
	"bufio"
	"strconv"
	"strings"
)

type Day17Input struct {
	x1, x2 int
	y1, y2 int
}

// NewDay17: parses input lines into Day17Input. Expected format: "target area: x=20..30, y=-10..-5"
func NewDay17(lines []string) Day17Input {
	line := strings.TrimSpace(strings.Join(lines, " "))
	nums := []int{}
	cur := ""
	for _, r := range line {
		if (r >= '0' && r <= '9') || r == '-' {
			cur += string(r)
		} else {
			if cur != "" {
				v, _ := strconv.Atoi(cur)
				nums = append(nums, v)
				cur = ""
			}
		}
	}
	if cur != "" {
		v, _ := strconv.Atoi(cur)
		nums = append(nums, v)
	}
	return Day17Input{x1: nums[0], x2: nums[1], y1: nums[2], y2: nums[3]}
}

func inTarget(px, py int, d Day17Input) bool {
	return px >= d.x1 && px <= d.x2 && py >= d.y1 && py <= d.y2
}

// Day17 computes the solution. For part1=true, solve Part 1; for part1=false, Part 2 would be computed if unlocked.
func Day17(data Day17Input, part1 bool) uint {
	maxHeight := 0
	hitCount := 0

	// Reasonable bounds: vx0 in [0..x2], vy0 in [y1..-y1]
	// The upper bound -y1 works for typical inputs where y1 is negative.
	upperVy := -data.y1
	if upperVy < 0 {
		upperVy = 0
	}

	for vx0 := 0; vx0 <= data.x2; vx0++ {
		for vy0 := data.y1; vy0 <= upperVy; vy0++ {
			vx, vy := vx0, vy0
			x, y := 0, 0
			localMax := 0
			hit := false
			// Simulate steps; stop when probe passes beyond target window
			for step := 0; step < 1000; step++ {
				x += vx
				y += vy
				if y > localMax {
					localMax = y
				}
				if inTarget(x, y, data) {
					hit = true
					break
				}
				if x > data.x2 || y < data.y1 {
					break
				}
				if vx > 0 {
					vx--
				} else if vx < 0 {
					vx++
				}
				vy--
			}
			if hit {
				if localMax > maxHeight {
					maxHeight = localMax
				}
				hitCount++
			}
		}
	}

	if part1 {
		return uint(maxHeight)
	}
	return uint(hitCount)
}

// Helpers for tests
func ReadLinesFromReader(r *bufio.Scanner) []string {
	lines := []string{}
	for r.Scan() {
		lines = append(lines, r.Text())
	}
	return lines
}
