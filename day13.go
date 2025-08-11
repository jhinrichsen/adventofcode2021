package adventofcode2021

import "strings"

// Point represents a 2D point with X and Y coordinates
type Point struct {
	X, Y int
}

// BitImage represents a 2D grid of bits
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
func (img *BitImage) Count() int {
	count := 0
	for _, word := range img.data {
		for word != 0 {
			count += int(word & 1)
			word >>= 1
		}
	}
	return count
}

// ToASCII returns the image as a slice of strings, one per row
// '#' for set bits and '.' for unset bits
func (img *BitImage) ToASCII() []string {
	if img.height == 0 || img.width == 0 {
		return nil
	}

	// Find content bounds
	minX, maxX := img.width-1, 0
	minY, maxY := img.height-1, 0
	for y := 0; y < img.height; y++ {
		for x := 0; x < img.width; x++ {
			if img.Get(x, y) {
				if x < minX { minX = x }
				if x > maxX { maxX = x }
				if y < minY { minY = y }
				if y > maxY { maxY = y }
			}
		}
	}

	// If no points, return empty
	if minX > maxX || minY > maxY {
		return nil
	}

	// Calculate content dimensions
	contentWidth := maxX - minX + 1
	contentHeight := maxY - minY + 1

	// Create a grid with the content dimensions
	grid := make([][]rune, contentHeight)
	for y := 0; y < contentHeight; y++ {
		grid[y] = make([]rune, contentWidth)
		for x := 0; x < contentWidth; x++ {
			if img.Get(minX + x, minY + y) {
				grid[y][x] = '#'
			} else {
				grid[y][x] = '.'
			}
		}
	}

	// Convert grid to []string
	result := make([]string, contentHeight)
	for y := 0; y < contentHeight; y++ {
		result[y] = string(grid[y])
	}

	// Ensure minimum dimensions for OCR (6x4)
	minOCRWidth := 6
	minOCRHeight := 4

	// Pad width if needed
	if contentWidth < minOCRWidth {
		padding := (minOCRWidth - contentWidth) / 2
		for i := range result {
			result[i] = strings.Repeat(".", padding) + result[i] + strings.Repeat(".", minOCRWidth-contentWidth-padding)
		}
	}

	// Pad height if needed
	if contentHeight < minOCRHeight {
		paddingTop := (minOCRHeight - contentHeight) / 2
		paddingBottom := minOCRHeight - contentHeight - paddingTop

		// Pad top
		for i := 0; i < paddingTop; i++ {
			result = append([]string{strings.Repeat(".", len(result[0]))}, result...)
		}

		// Pad bottom
		for i := 0; i < paddingBottom; i++ {
			result = append(result, strings.Repeat(".", len(result[0])))
		}
	}

	return result
}

// reverseBits64 reverses the bits in a 64-bit word
func reverseBits64(x uint64) uint64 {
	x = (x>>1)&0x5555555555555555 | (x&0x5555555555555555)<<1
	x = (x>>2)&0x3333333333333333 | (x&0x3333333333333333)<<2
	x = (x>>4)&0x0F0F0F0F0F0F0F0F | (x&0x0F0F0F0F0F0F0F0F)<<4
	x = (x>>8)&0x00FF00FF00FF00FF | (x&0x00FF00FF00FF00FF)<<8
	x = (x>>16)&0x0000FFFF0000FFFF | (x&0x0000FFFF0000FFFF)<<16
	x = (x >> 32) | (x << 32)
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
	// Use a map-based approach for counting dots when we don't need the full grid
	if limit > 0 && limit <= uint(len(folds)) {
		// If we have a limit and it's less than the number of folds, use the map approach
		dotSet := make(map[Point]bool)
		for _, pt := range points {
			dotSet[pt] = true
		}

		// Apply up to 'limit' folds
		for i := 0; i < int(limit) && i < len(folds); i++ {
			fold := folds[i]
			newDotSet := make(map[Point]bool)
			if fold > 0 {
				// Vertical fold (fold along x=fold)
				for pt := range dotSet {
					if pt.X > fold {
						// Fold left - mirror the point
						newX := 2*fold - pt.X
						newDotSet[Point{X: newX, Y: pt.Y}] = true
					} else {
						// Keep as is
						newDotSet[pt] = true
					}
				}
			} else {
				// Horizontal fold (fold along y=-fold)
				fold = -fold
				for pt := range dotSet {
					if pt.Y > fold {
						// Fold up - mirror the point
						newY := 2*fold - pt.Y
						newDotSet[Point{X: pt.X, Y: newY}] = true
					} else {
						// Keep as is
						newDotSet[pt] = true
					}
				}
			}
			dotSet = newDotSet
		}
		return uint(len(dotSet)), nil
	}

	// If no limit or limit exceeds number of folds, use the BitImage implementation
	// Find grid size
	w, h := 0, 0
	for _, pt := range points {
		if pt.X > w {
			w = pt.X
		}
		if pt.Y > h {
			h = pt.Y
		}
	}
	w++
	h++

	// Create BitImage
	img := NewBitImage(w, h)

	// Fill points
	for _, pt := range points {
		img.Set(pt.X, pt.Y)
	}

	// Apply folds up to the limit (or all if limit is 0)
	for i, fold := range folds {
		if limit > 0 && uint(i) >= limit {
			break
		}
		if fold > 0 {
			// Vertical fold (fold along x=fold)
			img.FoldVertical(fold)
		} else {
			// Horizontal fold (fold along y=-fold)
			img.FoldHorizontal(-fold)
		}
	}

	// Get the ASCII representation as a slice of strings
	return uint(img.Count()), img.ToASCII()
}
