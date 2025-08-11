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
        for j := range len(b) - valIdx {
            c := b[valIdx+j]
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
    // Determine how many folds to apply
    applyCount := len(folds)
    if limit > 0 {
        applyCount = min(len(folds), int(limit))
    }

    // Compute initial maxima to size Part 1 bitmap precisely
    maxX0, maxY0 := 0, 0
    for _, pt := range points {
        if pt.X > maxX0 {
            maxX0 = pt.X
        }
        if pt.Y > maxY0 {
            maxY0 = pt.Y
        }
    }

    // Determine target dimensions
    w, h := 0, 0
    if applyCount == 0 {
        // No folds: width/height are just maxima + 1 (edge case; not used by tests)
        w, h = maxX0+1, maxY0+1
    } else if limit > 0 {
        // Part 1: size after the first fold only
        f := folds[0]
        if f > 0 {
            // x-fold, width becomes f, height unchanged
            w = f
            h = maxY0 + 1
        } else {
            // y-fold, height becomes -f, width unchanged
            w = maxX0 + 1
            h = -f
        }
    } else {
        // Part 2: final paper size from smallest fold lines
        minXF, minYF := -1, -1
        for _, f := range folds {
            if f > 0 {
                if minXF == -1 || f < minXF {
                    minXF = f
                }
            } else {
                y := -f
                if minYF == -1 || y < minYF {
                    minYF = y
                }
            }
        }
        if minXF >= 0 {
            w = minXF
        } else {
            w = maxX0 + 1
        }
        if minYF >= 0 {
            h = minYF
        } else {
            h = maxY0 + 1
        }
    }

    // Bitset helpers
    area := w * h
    bits := make([]uint64, (area+63)>>6)
    setBit := func(i int) bool {
        idx := i >> 6
        mask := uint64(1) << (uint(i) & 63)
        old := bits[idx]
        bits[idx] = old | mask
        return (old & mask) == 0
    }

    // Fold a single point through N folds; return final coords and ok=false if it lands on a fold line
    foldPoint := func(x, y int) (nx, ny int, ok bool) {
        nx, ny = x, y
        for i := range applyCount {
            f := folds[i]
            if f > 0 {
                // x-fold at x=f
                if nx > f {
                    nx = 2*f - nx
                } else if nx == f {
                    return 0, 0, false
                }
            } else {
                yf := -f
                if ny > yf {
                    ny = 2*yf - ny
                } else if ny == yf {
                    return 0, 0, false
                }
            }
        }
        return nx, ny, true
    }

    // Populate bitset and count unique dots
    var count uint
    if applyCount == 0 {
        // No folds: just set existing points (not used by AoC but keep correctness)
        for _, p := range points {
            if p.X >= 0 && p.X < w && p.Y >= 0 && p.Y < h {
                if setBit(p.Y*w + p.X) {
                    count++
                }
            }
        }
    } else {
        for _, p := range points {
            x, y, ok := foldPoint(p.X, p.Y)
            if !ok || x < 0 || x >= w || y < 0 || y >= h {
                continue
            }
            if setBit(y*w + x) {
                count++
            }
        }
    }

    // Part 1 path: return count only
    if limit > 0 {
        return count, nil
    }

    // Part 2: render full image using '.' and '#'
    result := make([]string, h)
    row := make([]byte, w)
    for y := range h {
        base := y * w
        for x := range w {
            i := base + x
            if (bits[i>>6]>>(uint(i)&63))&1 == 1 {
                row[x] = '#'
            } else {
                row[x] = '.'
            }
        }
        result[y] = string(row)
    }
    return count, result
}
