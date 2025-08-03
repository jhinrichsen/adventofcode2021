// day20.go
package adventofcode2021

import "fmt"

// TrenchMap holds the enhancement algorithm and current image.
type TrenchMap struct {
	algo                   string
	img                    map[int]map[int]bool
	minY, maxY, minX, maxX int
}

// NewDay20 parses the raw lines into a TrenchMap.
func NewDay20(lines []string) (*TrenchMap, error) {
	if len(lines) < 3 {
		return nil, fmt.Errorf("input too short")
	}
	algo := lines[0]
	img := make(map[int]map[int]bool)
	for y, row := range lines[2:] {
		for x, ch := range row {
			if img[y] == nil {
				img[y] = make(map[int]bool)
			}
			img[y][x] = (ch == '#')
		}
	}
	h := len(lines) - 2
	w := len(lines[2])
	return &TrenchMap{
		algo: algo,
		img:  img,
		minY: 0, maxY: h - 1,
		minX: 0, maxX: w - 1,
	}, nil
}

// Day20 runs either part1 (2 steps) or part2 (50 steps) and returns the lit‐pixel count.
func Day20(tm *TrenchMap, part1 bool) int {
	steps := 50
	if part1 {
		steps = 2
	}
	fill := false
	for i := 0; i < steps; i++ {
		tm.img, tm.minY, tm.maxY, tm.minX, tm.maxX, fill = enhance(tm.algo, tm.img, tm.minY, tm.maxY, tm.minX, tm.maxX, fill)
	}
	return countLit(tm.img)
}

// enhance performs one enhancement step, expanding bounds by 1 in every direction.
func enhance(algo string, img map[int]map[int]bool, minY, maxY, minX, maxX int, fill bool) (map[int]map[int]bool, int, int, int, int, bool) {
	nbMinY, nbMaxY := minY-1, maxY+1
	nbMinX, nbMaxX := minX-1, maxX+1
	out := make(map[int]map[int]bool)

	for y := nbMinY; y <= nbMaxY; y++ {
		for x := nbMinX; x <= nbMaxX; x++ {
			idx := 0
			for dy := -1; dy <= 1; dy++ {
				for dx := -1; dx <= 1; dx++ {
					idx <<= 1
					yy, xx := y+dy, x+dx
					var bit bool
					if yy < minY || yy > maxY || xx < minX || xx > maxX {
						bit = fill
					} else {
						bit = img[yy][xx]
					}
					if bit {
						idx |= 1
					}
				}
			}
			if algo[idx] == '#' {
				if out[y] == nil {
					out[y] = make(map[int]bool)
				}
				out[y][x] = true
			}
		}
	}

	// flip fill if algorithm[0]=='#'
	newFill := fill
	if algo[0] == '#' {
		newFill = !fill
	}

	return out, nbMinY, nbMaxY, nbMinX, nbMaxX, newFill
}

// countLit counts the number of lit pixels.
func countLit(img map[int]map[int]bool) int {
	cnt := 0
	for _, row := range img {
		for _, on := range row {
			if on {
				cnt++
			}
		}
	}
	return cnt
}
