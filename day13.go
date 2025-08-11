package adventofcode2021

import "image"

// Point is an alias to image.Point for text-based image coordinates
type Point = image.Point

// NewDay13 parses the input lines into dots and fold instructions
func NewDay13(lines []string) ([]Point, []int) {
    dots := make([]Point, 0, 1024)
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
				dots = append(dots, Point{X: x, Y: y})
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
// Parameters:
//   - points: slice of points representing the dots
//   - folds: slice of fold instructions (positive for x, negative for y)
//   - limit: maximum number of folds to apply (0 means apply all)
//
// Returns:
//   - count of visible dots
//   - the final grid as a slice of strings, one string per line
func Day13(points []Point, folds []int, limit uint) (uint, []string) {
    // Single algorithm: map-based folding for both parts
    dotSet := make(map[Point]bool)
    for _, pt := range points {
        dotSet[pt] = true
    }

    // Determine how many folds to apply
    applyCount := len(folds)
    if limit > 0 {
        applyCount = min(len(folds), int(limit))
    }

    // Apply folds
    for i := range applyCount {
        fold := folds[i]
        newDotSet := make(map[Point]bool)
        if fold > 0 {
            // Vertical fold along x=fold
            for pt := range dotSet {
                switch {
                case pt.X > fold:
                    newX := 2*fold - pt.X
                    newDotSet[Point{X: newX, Y: pt.Y}] = true
                case pt.X < fold:
                    newDotSet[pt] = true
                // points on the fold are discarded
                }
            }
        } else {
            // Horizontal fold along y=-fold
            fold = -fold
            for pt := range dotSet {
                switch {
                case pt.Y > fold:
                    newY := 2*fold - pt.Y
                    newDotSet[Point{X: pt.X, Y: newY}] = true
                case pt.Y < fold:
                    newDotSet[pt] = true
                // points on the fold are discarded
                }
            }
        }
        dotSet = newDotSet
    }

    // Part 1 path (limit > 0): return count only
    if limit > 0 {
        return uint(len(dotSet)), nil
    }

    // Part 2 path (limit == 0): render full image using paper size implied by folds
    // Determine paper bounds from folds: width is the smallest x-fold value,
    // height is the smallest y-fold value. If none for an axis, use dot maxima.
    minXF, minYF := -1, -1
    for _, f := range folds {
        if f > 0 { // x fold at x=f
            if minXF == -1 || f < minXF {
                minXF = f
            }
        } else { // y fold at y=-f
            y := -f
            if minYF == -1 || y < minYF {
                minYF = y
            }
        }
    }

    maxX, maxY := 0, 0
    for pt := range dotSet {
        maxX = max(maxX, pt.X)
        maxY = max(maxY, pt.Y)
    }
    w := maxX + 1
    h := maxY + 1
    if minXF >= 0 {
        w = minXF
    }
    if minYF >= 0 {
        h = minYF
    }

    result := make([]string, h)
    row := make([]rune, w)
    for y := 0; y < h; y++ {
        for x := 0; x < w; x++ {
            if dotSet[Point{X: x, Y: y}] {
                row[x] = '#'
            } else {
                row[x] = '.'
            }
        }
        result[y] = string(row)
    }
    return uint(len(dotSet)), result
}
