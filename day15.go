package adventofcode2021

import (
	"errors"
	"math"
)

// Day15 computes the lowest total risk of any path from the top left to the bottom right.
// The risk of a path is the sum of the risk levels of each position you enter; the starting
// position is not counted.
//
// lines is the raw grid input. part1 selects Part 1 vs Part 2 logic. Only Part 1 is implemented here.
func Day15(lines []string, part1 bool) (uint, error) {
	// Filter out empty lines
	/*
		var gridLines []string
		for _, line := range lines {
			if len(line) > 0 {
				gridLines = append(gridLines, line)
			}
		}
	*/

	baseRows := len(lines)
	baseCols := len(lines[0])

	rows, cols := baseRows, baseCols
	if !part1 {
		rows *= 5
		cols *= 5
	}
	N := rows * cols

	// Build weights. For Part 1, copy directly. For Part 2, expand 5x with wrapping risk increments.
	weights := make([]byte, N)
	if part1 {
		for r := 0; r < rows; r++ {
			row := lines[r]
			for c := 0; c < cols; c++ {
				weights[r*cols+c] = row[c] - '0'
			}
		}
	} else {
		for r := 0; r < rows; r++ {
			for c := 0; c < cols; c++ {
				br := r % baseRows
				bc := c % baseCols
				inc := (r / baseRows) + (c / baseCols)
				base := int(lines[br][bc] - '0')
				// wrap risk: 1..9
				v := ((base - 1 + inc) % 9) + 1
				weights[r*cols+c] = byte(v)
			}
		}
	}

	// Dijkstra on a 2D grid using Dial's algorithm (bucketed for weights 1..9).
	const inf = int32(math.MaxInt32)
	dist := make([]int32, N)
	for i := range dist {
		dist[i] = inf
	}
	visited := make([]bool, N)
	start := 0
	target := N - 1
	dist[start] = 0

	const C = 9            // max edge weight
	B := C * (rows + cols) // safe upper bound on shortest-path cost
	buckets := make([][]int, B)
	buckets[0] = append(buckets[0], start)
	var curDist int32 = 0
	processed := 0

	for processed < N {
		// advance to next non-empty bucket; if we run out, no path
		for int(curDist) < B && len(buckets[curDist]) == 0 {
			curDist++
		}
		if int(curDist) >= B {
			break
		}
		b := &buckets[curDist]
		// pop from bucket (LIFO is fine)
		idx := (*b)[len(*b)-1]
		*b = (*b)[:len(*b)-1]

		if visited[idx] || dist[idx] != curDist {
			continue
		}
		visited[idx] = true
		processed++
		if idx == target {
			return uint(dist[idx]), nil
		}

		r, c := idx/cols, idx%cols
		// Up
		if rr, cc := r-1, c; rr >= 0 {
			ni := rr*cols + cc
			if !visited[ni] {
				nd := curDist + int32(weights[ni])
				if nd < dist[ni] {
					dist[ni] = nd
					if int(nd) < B {
						buckets[nd] = append(buckets[nd], ni)
					}
				}
			}
		}
		// Down
		if rr, cc := r+1, c; rr < rows {
			ni := rr*cols + cc
			if !visited[ni] {
				nd := curDist + int32(weights[ni])
				if nd < dist[ni] {
					dist[ni] = nd
					if int(nd) < B {
						buckets[nd] = append(buckets[nd], ni)
					}
				}
			}
		}
		// Left
		if rr, cc := r, c-1; cc >= 0 {
			ni := rr*cols + cc
			if !visited[ni] {
				nd := curDist + int32(weights[ni])
				if nd < dist[ni] {
					dist[ni] = nd
					if int(nd) < B {
						buckets[nd] = append(buckets[nd], ni)
					}
				}
			}
		}
		// Right
		if rr, cc := r, c+1; cc < cols {
			ni := rr*cols + cc
			if !visited[ni] {
				nd := curDist + int32(weights[ni])
				if nd < dist[ni] {
					dist[ni] = nd
					if int(nd) < B {
						buckets[nd] = append(buckets[nd], ni)
					}
				}
			}
		}
	}

	return 0, errors.New("no path found")
}
