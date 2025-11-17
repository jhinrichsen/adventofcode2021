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

	for step := 0; step < steps; step++ {
		// Clear next buffer
		clear(next)

		// Process all pixels - optimize by manually unrolling the 3x3 window
		for y := range size {
			yBase := y * size
			for x := range size {
				index := 0

				// Manually unroll 3x3 window for better performance
				// Row -1
				if y > 0 {
					prevRow := (y-1) * size
					if x > 0 {
						index = int(current[prevRow+x-1])
					} else {
						index = int(infiniteValue)
					}
					index = (index << 1) | int(current[prevRow+x])
					if x < size-1 {
						index = (index << 1) | int(current[prevRow+x+1])
					} else {
						index = (index << 1) | int(infiniteValue)
					}
				} else {
					index = int(infiniteValue)
					index = (index << 1) | int(infiniteValue)
					index = (index << 1) | int(infiniteValue)
				}

				// Row 0 (current)
				if x > 0 {
					index = (index << 1) | int(current[yBase+x-1])
				} else {
					index = (index << 1) | int(infiniteValue)
				}
				index = (index << 1) | int(current[yBase+x])
				if x < size-1 {
					index = (index << 1) | int(current[yBase+x+1])
				} else {
					index = (index << 1) | int(infiniteValue)
				}

				// Row +1
				if y < size-1 {
					nextRow := (y+1) * size
					if x > 0 {
						index = (index << 1) | int(current[nextRow+x-1])
					} else {
						index = (index << 1) | int(infiniteValue)
					}
					index = (index << 1) | int(current[nextRow+x])
					if x < size-1 {
						index = (index << 1) | int(current[nextRow+x+1])
					} else {
						index = (index << 1) | int(infiniteValue)
					}
				} else {
					index = (index << 1) | int(infiniteValue)
					index = (index << 1) | int(infiniteValue)
					index = (index << 1) | int(infiniteValue)
				}

				if algo[index] == '#' {
					next[yBase+x] = 1
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
