package adventofcode2021

import (
	"image"
)

// NewDay13 parses the input lines into dots and fold instructions
func NewDay13(lines []string) (map[image.Point]struct{}, []int) {
	dots := make(map[image.Point]struct{})
	folds := make([]int, 0)
	parsingDots := true

	for _, line := range lines {
		b := []byte(line)
		if len(b) == 0 {
			parsingDots = false
			continue
		}

		if parsingDots {
			// Parse "x,y"
			x, y, mode, val := 0, 0, 0, 0
			for _, c := range b {
				switch c {
				case ',':
					x = val
					val = 0
					mode = 1
				default:
					if c >= '0' && c <= '9' {
						val = val*10 + int(c-'0')
					}
				}
			}
			if mode == 1 {
				y = val
				dots[image.Point{X: x, Y: y}] = struct{}{}
			}
			continue
		}

		// Parse fold: "fold along x=number" or "fold along y=number"
		// We know line must begin: 'f','o','l','d',' ','a','l','o','n','g',' '
		if len(b) < 13 {
			continue
		}
		axisC := b[11]
		valIdx := 13 // character after 'x=' or 'y='
		val := 0
		for j := valIdx; j < len(b); j++ {
			c := b[j]
			if c < '0' || c > '9' {
				break
			}
			val = val*10 + int(c-'0')
		}
		if axisC == 'y' {
			val = -val
		}
		folds = append(folds, val)
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
