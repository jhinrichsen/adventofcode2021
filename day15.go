package adventofcode2021

import (
	"container/heap"
	"errors"
)

// Day15 computes the lowest total risk of any path from the top left to the bottom right.
// The risk of a path is the sum of the risk levels of each position you enter; the starting
// position is not counted.
//
// lines is the raw grid input. part1 selects Part 1 vs Part 2 logic. Only Part 1 is implemented here.
func Day15(lines []string, part1 bool) (uint, error) {
	// Filter out empty lines
	var gridLines []string
	for _, line := range lines {
		if len(line) > 0 {
			gridLines = append(gridLines, line)
		}
	}
	baseRows := len(gridLines)
	baseCols := len(gridLines[0])

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
			row := gridLines[r]
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
				base := int(gridLines[br][bc] - '0')
				// wrap risk: 1..9
				v := ((base - 1 + inc) % 9) + 1
				weights[r*cols+c] = byte(v)
			}
		}
	}

	// Dijkstra on a 2D grid using a min-heap.
	const inf = int(^uint(0) >> 1) // max int
	dist := make([]int, N)
	for i := range dist {
		dist[i] = inf
	}
	start := 0
	target := N - 1
	dist[start] = 0

	pq := &minHeap{}
	heap.Push(pq, node{idx: start, dist: 0})

	inBounds := func(r, c int) bool { return r >= 0 && r < rows && c >= 0 && c < cols }

	for pq.Len() > 0 {
		n := heap.Pop(pq).(node)
		if n.idx == target {
			return uint(n.dist), nil
		}
		// If this entry is stale, skip
		if n.dist != dist[n.idx] {
			continue
		}
		r, c := n.idx/cols, n.idx%cols
		// 4 neighbors
		// Up
		if rr, cc := r-1, c; inBounds(rr, cc) {
			ni := rr*cols + cc
			nd := n.dist + int(weights[ni])
			if nd < dist[ni] {
				dist[ni] = nd
				heap.Push(pq, node{idx: ni, dist: nd})
			}
		}
		// Down
		if rr, cc := r+1, c; inBounds(rr, cc) {
			ni := rr*cols + cc
			nd := n.dist + int(weights[ni])
			if nd < dist[ni] {
				dist[ni] = nd
				heap.Push(pq, node{idx: ni, dist: nd})
			}
		}
		// Left
		if rr, cc := r, c-1; inBounds(rr, cc) {
			ni := rr*cols + cc
			nd := n.dist + int(weights[ni])
			if nd < dist[ni] {
				dist[ni] = nd
				heap.Push(pq, node{idx: ni, dist: nd})
			}
		}
		// Right
		if rr, cc := r, c+1; inBounds(rr, cc) {
			ni := rr*cols + cc
			nd := n.dist + int(weights[ni])
			if nd < dist[ni] {
				dist[ni] = nd
				heap.Push(pq, node{idx: ni, dist: nd})
			}
		}
	}

	return 0, errors.New("no path found")
}

// node represents a position in the grid with its current best-known distance.
type node struct {
	idx  int
	dist int
}

type minHeap []node

func (h minHeap) Len() int           { return len(h) }
func (h minHeap) Less(i, j int) bool { return h[i].dist < h[j].dist }
func (h minHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *minHeap) Push(x any)        { *h = append(*h, x.(node)) }
func (h *minHeap) Pop() any          { old := *h; n := len(old); x := old[n-1]; *h = old[:n-1]; return x }
