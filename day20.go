// day20.go
package adventofcode2021

import (
	"fmt"
	"strings"
)

// TrenchMap holds the 512-byte algorithm and two flat buffers.
type TrenchMap struct {
	algo      [512]byte
	buf, next []byte
	size      int
}

// NewDay20 parses the input lines (single- or multi-line algorithm) and
// initializes flat buffers sized for 50 steps of padding.
func NewDay20(lines []string) (*TrenchMap, error) {
	// locate blank separator
	sep := -1
	for i, l := range lines {
		if l == "" {
			sep = i
			break
		}
	}

	var algoLines, imageLines []string
	if sep >= 0 {
		algoLines = lines[:sep]
		imageLines = lines[sep+1:]
	} else {
		if len(lines) < 2 {
			return nil, fmt.Errorf("not enough lines")
		}
		algoLines = lines[:1]
		imageLines = lines[1:]
	}

	algostr := strings.Join(algoLines, "")
	if len(algostr) != 512 {
		return nil, fmt.Errorf("algorithm length invalid: got %d, want 512", len(algostr))
	}

	var algoArr [512]byte
	for i := 0; i < 512; i++ {
		algoArr[i] = algostr[i]
	}

	h0 := len(imageLines)
	w0 := 0
	if h0 > 0 {
		w0 = len(imageLines[0])
	}

	padding := 50
	size := w0 + 2*padding
	buf := make([]byte, size*size)
	next := make([]byte, size*size)

	// draw initial image into center of buf
	offset := padding
	for y, row := range imageLines {
		base := (offset+y)*size + offset
		for x := 0; x < len(row); x++ {
			if row[x] == '#' {
				buf[base+x] = 1
			}
		}
	}

	return &TrenchMap{
		algo: algoArr,
		buf:  buf,
		next: next,
		size: size,
	}, nil
}

// Constants for Day20 implementation
const (
	neighborhood = 3 * 3
	darkPixel    = 0
	lightPixel   = 1
	outOfBounds  = -1 // Marker for out-of-bounds pixels
)

// Day20 runs 2 enhancement steps if part1==true, else 50, returning lit-pixel count.
func Day20(tm *TrenchMap, part1 bool) int {
	steps := 50
	if part1 {
		steps = 2
	}

	algo := tm.algo[:]
	size := tm.size
	bufSize := len(tm.buf)
	current := make([]byte, bufSize)
	next := make([]byte, bufSize)
	copy(current, tm.buf)

	// Track infinite grid state
	infiniteValue := byte(0)
	toggleInfinite := (algo[0] == '#' && algo[511] == '.')

	// Only precompute offsets for part 2 (50 steps)
	var offsets [][9]int
	if !part1 {
		offsets = make([][9]int, bufSize)
		for y := range size {
			for x := range size {
				i := y*size + x
				offsetIdx := 0
				for dy := -1; dy <= 1; dy++ {
					for dx := -1; dx <= 1; dx++ {
						ny, nx := y+dy, x+dx
						offset := ny*size + nx
						if ny < 0 || ny >= size || nx < 0 || nx >= size {
							offset = -1
						}
						offsets[i][offsetIdx] = offset
						offsetIdx++
					}
				}
			}
		}
	}

	for step := 0; step < steps; step++ {
		// Clear next buffer
		for i := range next {
			next[i] = 0
		}

		if part1 {
			// Optimized path for part 1 (2 steps)
			for y := range size {
				for x := range size {
					index := 0
					for dy := -1; dy <= 1; dy++ {
						for dx := -1; dx <= 1; dx++ {
							ny, nx := y+dy, x+dx
							bit := infiniteValue
							if ny >= 0 && ny < size && nx >= 0 && nx < size {
								bit = current[ny*size+nx]
							}
							index = (index << 1) | int(bit)
						}
					}
					if algo[index] == '#' {
						next[y*size+x] = 1
					}
				}
			}
		} else {
			// Optimized path for part 2 (50 steps)
			for i := 0; i < bufSize; i++ {
				index := 0
				for j := 0; j < 9; j++ {
					offset := offsets[i][j]
					bit := infiniteValue
					if offset != -1 {
						bit = current[offset]
					}
					index = (index << 1) | int(bit)
				}
				if algo[index] == '#' {
					next[i] = 1
				}
			}
		}

		current, next = next, current

		if toggleInfinite {
			infiniteValue = 1 - infiniteValue
			if infiniteValue == 1 {
				// Optimized border handling
				sizeMinus1 := size - 1
				for i := 0; i < size; i++ {
					// Top and bottom borders
					current[i] = 1
					current[sizeMinus1*size+i] = 1

					// Left and right borders
					current[i*size] = 1
					current[i*size+sizeMinus1] = 1
				}
			}
		}
	}

	// Count all lit pixels in the final image using a single loop for better performance
	cnt := 0
	for i := 0; i < len(current); i++ {
		cnt += int(current[i])
	}

	return cnt
}
