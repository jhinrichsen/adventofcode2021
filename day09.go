package adventofcode2021

import "fmt"

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
			fmt.Printf("(%d/%d) = %d", x, y, h)

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
				fmt.Printf(" <")
			}
			fmt.Println()
		}
	}

	fmt.Printf("%d low points\n", len(lows))
	var risk uint
	for i := range lows {
		risk += uint(lows[i]) + 1
	}
	// risk += uint(len(lows)) // risk = height + 1
	fmt.Printf("%d risk\n", risk)
	return risk, nil
}
