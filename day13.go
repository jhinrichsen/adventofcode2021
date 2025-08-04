package adventofcode2021

import (
	"image"
	"strconv"
	"strings"
)

// NewDay13 parses the input lines into dots and fold instructions
func NewDay13(lines []string) (map[image.Point]struct{}, []int) {
	dots := make(map[image.Point]struct{})
	folds := make([]int, 0)

	parsingDots := true

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			parsingDots = false
			continue
		}
		if parsingDots {
			// Parse dot coordinates: "x,y"
			parts := strings.Split(line, ",")
			if len(parts) == 2 {
				x, _ := strconv.Atoi(parts[0])
				y, _ := strconv.Atoi(parts[1])
				dots[image.Point{X: x, Y: y}] = struct{}{}
			}
		} else {
			// Parse fold instructions: "fold along x=5" or "fold along y=7"
			if strings.HasPrefix(line, "fold along ") {
				instruction := strings.TrimPrefix(line, "fold along ")
				parts := strings.Split(instruction, "=")
				if len(parts) == 2 {
					axis := parts[0]
					value, _ := strconv.Atoi(parts[1])
					if axis == "x" {
						folds = append(folds, value)
					} else if axis == "y" {
						folds = append(folds, -value)
					}
				}
			}
		}
	}
	return dots, folds
}

// Day13 solves the transparent origami puzzle
func Day13(dotsIn map[image.Point]struct{}, folds []int, part1 bool) uint {
	// Make a copy of dots to avoid modifying original data
	dots := make(map[image.Point]struct{})
	for point := range dotsIn {
		dots[point] = struct{}{}
	}

	// Apply folds
	foldsToApply := folds
	if part1 {
		// Part 1: only apply the first fold
		foldsToApply = folds[:1]
	}

	for _, fold := range foldsToApply {
		dots = applyFold(dots, fold)
	}

	return uint(len(dots))
}

// applyFold applies a single fold to the set of dots
// Positive value = X fold (vertical); Negative value = Y fold (horizontal)
func applyFold(dots map[image.Point]struct{}, fold int) map[image.Point]struct{} {
	newDots := make(map[image.Point]struct{})
	if fold == 0 {
		return dots
	}

	if fold > 0 {
		// X axis fold at fold=X
		line := fold
		for point := range dots {
			var newPoint image.Point
			if point.X < line {
				newPoint = point
			} else {
				newPoint = image.Point{X: 2*line - point.X, Y: point.Y}
			}
			newDots[newPoint] = struct{}{}
		}
	} else {
		// Y axis fold at fold=-fold
		line := -fold
		for point := range dots {
			var newPoint image.Point
			if point.Y < line {
				newPoint = point
			} else {
				newPoint = image.Point{X: point.X, Y: 2*line - point.Y}
			}
			newDots[newPoint] = struct{}{}
		}
	}
	return newDots
}
