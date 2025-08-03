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
	size := tm.size
	buf := tm.buf
	next := tm.next
	algo := tm.algo[:]

	// neighbor offsets in the flat buffer
	nbs := [9]int{
		-size - 1, -size, -size + 1,
		-1, 0, +1,
		+size - 1, +size, +size + 1,
	}

	fill := byte(0)
	toggle := (algo[0] == '#')

	for step := 0; step < steps; step++ {
		// pre-fill next[] to the infinite background value
		if fill == 1 {
			for i := range next {
				next[i] = 1
			}
		} else {
			for i := range next {
				next[i] = 0
			}
		}

		// enhancement pass: only update inner pixels [1..size-2]
		for y := 1; y < size-1; y++ {
			base := y * size
			for x := 1; x < size-1; x++ {
				idx := base + x
				code := 0
				for k, off := range nbs {
					if buf[idx+off] == 1 {
						code |= 1 << (8 - k)
					}
				}
				if algo[code] == '#' {
					next[idx] = 1
				}
			}
		}

		// swap buffers and toggle fill if needed
		buf, next = next, buf
		if toggle {
			fill ^= 1
		}
	}

	// count lit pixels in the final buffer
	cnt := 0
	for _, v := range buf {
		cnt += int(v)
	}

	// store buffers back into tm
	tm.buf = buf
	tm.next = next
	return cnt
}
