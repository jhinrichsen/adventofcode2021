package adventofcode2021

import (
	"image"
	"slices"
)

// NewDay13 parses the input lines into dots and fold instructions
func NewDay13(lines []string) ([]image.Point, []int) {
	dots := make([]image.Point, 0, 1024)
	folds := make([]int, 0, 32)
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
				dots = append(dots, image.Point{X: x, Y: y})
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
func Day13(points []image.Point, folds []int, part1 bool) uint {
	// Find grid size
	w, h := 0, 0
	for _, pt := range points {
		w = max(w, pt.X)
		h = max(h, pt.Y)
	}
	w++
	h++
	gridA := make([]bool, w*h)
	gridB := slices.Clone(gridA)

	// fill points
	for _, pt := range points {
		gridA[pt.Y*w+pt.X] = true
	}

	// handles for buffer switching
	grid := &gridA
	buffer := &gridB

	for _, fold := range folds {
		if fold > 0 {
			foldX := fold
			for y := 0; y < h; y++ {
				for x := 0; x < foldX; x++ {
					r := 2*foldX - x
					(*buffer)[y*foldX+x] = (*grid)[y*w+x] || (r < w && (*grid)[y*w+r])
				}
			}
			w = foldX
		} else {
			foldY := -fold
			for y := 0; y < foldY; y++ {
				for x := 0; x < w; x++ {
					r := 2*foldY - y
					(*buffer)[y*w+x] = (*grid)[y*w+x] || (r < h && (*grid)[r*w+x])
				}
			}
			h = foldY
		}

		// Swap buffers
		buffer, grid = grid, buffer

		if part1 {
			break
		}

		// Clear the buffer
		for i := range *buffer {
			(*buffer)[i] = false
		}
	}

	var count uint
	for i := range *grid {
		if (*grid)[i] {
			count++
		}
	}
	return count
}
