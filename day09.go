package adventofcode2021

func Day09(lines []string, part1 bool) (uint, error) {
	Y := len(lines)
	X := len(lines[0])

	height := func(x, y int) byte {
		return lines[y][x] - '0'
	}

	var lows []byte
	for y := range lines {
		for x := range lines[y] {
			h := height(x, y)

			low := true
			// N
			if y > 0 && height(x, y-1) <= h {
				low = false
				// continue
			}
			// E
			if x < X-1 && height(x+1, y) <= h {
				low = false
				// continue
			}
			// S
			if y < Y-1 && height(x, y+1) <= h {
				low = false
				// continue
			}
			// E
			if x > 0 && height(x-1, y) <= h {
				low = false
				// continue
			}
			if low {
				lows = append(lows, h)
			}
		}
	}

	var risk uint
	for i := range lows {
		risk += uint(lows[i]) + 1
	}
	// risk += uint(len(lows)) // risk = height + 1
	return risk, nil
}
