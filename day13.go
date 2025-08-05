package adventofcode2021

import (
	"image"
	"math/bits"
)

// BitImage represents a 2D bit image stored as a flat array of uint64 words
type BitImage struct {
	data        []uint64
	width       int
	height      int
	wordsPerRow int
}

// NewBitImage creates a new BitImage with the given dimensions
func NewBitImage(width, height int) *BitImage {
	wordsPerRow := (width + 63) / 64
	totalWords := wordsPerRow * height
	return &BitImage{
		data:        make([]uint64, totalWords),
		width:       width,
		height:      height,
		wordsPerRow: wordsPerRow,
	}
}

// Set sets the bit at the given coordinates
func (img *BitImage) Set(x, y int) {
	if x >= 0 && x < img.width && y >= 0 && y < img.height {
		wordIdx := x / 64
		bitIdx := x % 64
		img.data[y*img.wordsPerRow+wordIdx] |= 1 << bitIdx
	}
}

// Get returns true if the bit at the given coordinates is set
func (img *BitImage) Get(x, y int) bool {
	if x >= 0 && x < img.width && y >= 0 && y < img.height {
		wordIdx := x / 64
		bitIdx := x % 64
		return (img.data[y*img.wordsPerRow+wordIdx] & (1 << bitIdx)) != 0
	}
	return false
}

// Clear clears all bits in the image
func (img *BitImage) Clear() {
	for i := range img.data {
		img.data[i] = 0
	}
}

// Count returns the number of set bits in the image
func (img *BitImage) Count() uint {
	var count uint
	for _, word := range img.data {
		count += uint(bits.OnesCount64(word))
	}
	return count
}

// reverseBits64 reverses the bits in a 64-bit word
func reverseBits64(x uint64) uint64 {
	x = (x>>1)&0x5555555555555555 | (x&0x5555555555555555)<<1
	x = (x>>2)&0x3333333333333333 | (x&0x3333333333333333)<<2
	x = (x>>4)&0x0F0F0F0F0F0F0F0F | (x&0x0F0F0F0F0F0F0F0F)<<4
	x = (x>>8)&0x00FF00FF00FF00FF | (x&0x00FF00FF00FF00FF)<<8
	x = (x>>16)&0x0000FFFF0000FFFF | (x&0x0000FFFF0000FFFF)<<16
	x = (x>>32) | (x<<32)
	return x
}

// FoldVertical folds the image along a vertical line in-place
func (img *BitImage) FoldVertical(foldLine int) {
	newWidth := foldLine
	newWordsPerRow := (newWidth + 63) / 64
	oldWordsPerRow := img.wordsPerRow

	for y := 0; y < img.height; y++ {
		rowOffset := y * img.wordsPerRow

		// Handle right side folding with word-level bit manipulation
		rightStartX := foldLine + 1
		if rightStartX < img.width {
			rightStartWord := rightStartX / 64

			// Process complete words from the right side
			for w := rightStartWord; w < oldWordsPerRow; w++ {
				if img.data[rowOffset+w] == 0 {
					continue // Skip empty words
				}

				word := img.data[rowOffset+w]
				wordStartX := w * 64
				wordEndX := wordStartX + 63

				// Check if this entire word can be processed with word-level operations
				if wordStartX >= rightStartX && wordEndX < img.width {
					// This word is entirely in the fold region
					// Calculate the mirror position for this word
					mirrorWordEndX := 2*foldLine - wordStartX
					mirrorWordStartX := 2*foldLine - wordEndX

					// Check if the mirrored word fits entirely within the new width
					if mirrorWordStartX >= 0 && mirrorWordEndX < newWidth {
						mirrorWordIdx := mirrorWordStartX / 64
						mirrorBitOffset := mirrorWordStartX % 64

						if mirrorBitOffset == 0 {
							// Perfect alignment - reverse bits and OR directly
							img.data[rowOffset+mirrorWordIdx] |= reverseBits64(word)
						} else {
							// Handle bit offset with word-level operations
							reversedWord := reverseBits64(word)

							// Split the reversed word across two target words
							img.data[rowOffset+mirrorWordIdx] |= reversedWord << mirrorBitOffset
							if mirrorWordIdx+1 < newWordsPerRow {
								img.data[rowOffset+mirrorWordIdx+1] |= reversedWord >> (64 - mirrorBitOffset)
							}
						}
						continue // Skip bit-by-bit processing for this word
					}
				}

				// Fall back to bit-by-bit processing (for words that span boundaries or don't fit entirely)
				for bit := 0; bit < 64 && word != 0; bit++ {
					if (word & (1 << bit)) != 0 {
						x := wordStartX + bit
						if x >= rightStartX {
							mirrorX := 2*foldLine - x
							if mirrorX >= 0 && mirrorX < newWidth {
								wordIdx := mirrorX / 64
								bitIdx := mirrorX % 64
								img.data[rowOffset+wordIdx] |= 1 << bitIdx
							}
						}
						word &= ^(1 << bit)
					}
				}
			}
		}

		// Clear words beyond the fold line (but don't truncate slice to avoid allocations)
		for w := newWordsPerRow; w < img.wordsPerRow; w++ {
			img.data[rowOffset+w] = 0
		}
		if newWordsPerRow > 0 {
			lastWordBits := newWidth % 64
			if lastWordBits > 0 {
				mask := (uint64(1) << lastWordBits) - 1
				img.data[rowOffset+newWordsPerRow-1] &= mask
			}
		}
	}

	img.width = newWidth
	img.wordsPerRow = newWordsPerRow
}

// FoldHorizontal folds the image along a horizontal line in-place
func (img *BitImage) FoldHorizontal(foldLine int) {
	newHeight := foldLine

	for r := 0; r < newHeight; r++ {
		mirrorRow := 2*foldLine - r
		if mirrorRow < img.height {
			// Use word-level operations for efficiency
			rowOffset := r * img.wordsPerRow
			mirrorRowOffset := mirrorRow * img.wordsPerRow
			for w := 0; w < img.wordsPerRow; w++ {
				img.data[rowOffset+w] |= img.data[mirrorRowOffset+w]
			}
		}
		// Note: if mirrorRow >= img.height, we just keep the existing row as-is
	}

	// Clear rows beyond the fold line (but don't truncate slice to avoid allocations)
	for r := newHeight; r < img.height; r++ {
		rowOffset := r * img.wordsPerRow
		for w := 0; w < img.wordsPerRow; w++ {
			img.data[rowOffset+w] = 0
		}
	}
	img.height = newHeight
}

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

	// Create BitImage
	img := NewBitImage(w, h)

	// fill points
	for _, pt := range points {
		img.Set(pt.X, pt.Y)
	}

	for _, fold := range folds {
		if fold > 0 {
			// Vertical fold (fold along x=fold)
			img.FoldVertical(fold)
		} else {
			// Horizontal fold (fold along y=-fold)
			img.FoldHorizontal(-fold)
		}

		if part1 {
			break
		}
	}

	return img.Count()
}
