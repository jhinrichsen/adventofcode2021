package adventofcode2021

import (
	"image"
	"strconv"
	"strings"
)

// parseNumbers converts a list of comma separated numbers.
func parseNumbers(line string) ([]int, error) {
	var ns []int
	parts := strings.Split(line, ",")
	for n := range parts {
		n, err := strconv.Atoi(parts[n])
		if err != nil {
			return ns, err
		}
		ns = append(ns, n)
	}
	return ns, nil
}

func NewDay04(lines []string) ([]int, []Bingo, error) {
	// first line is comma separated list of numbers
	draws, err := parseNumbers(lines[0])
	if err != nil {
		return nil, nil, err
	}

	nBoards := 0
	for i := range lines {
		if len(lines[i]) == 0 {
			nBoards++
		}
	}

	dimY := ((len(lines) - 1) / nBoards) - 1

	var boards []Bingo
	for y := 0; y < nBoards; y++ {
		idx := 2 + y*(dimY+1)
		b, err := NewBingo(lines[idx : idx+dimY])
		if err != nil {
			return nil, nil, err
		}
		boards = append(boards, b)
	}
	return draws, boards, nil
}

func Day04Part1(draws []int, boards []Bingo) int {
	for _, draw := range draws {
		for _, b := range boards {
			bingo := b.Draw(draw)
			if bingo {
				return Sum(b.Unmarked()) * draw
			}
		}
	}
	return 0
}

type Bingo struct {
	DimX, DimY int
	Drawn      map[int]bool
	Cols, Rows []int // number of hits per col/ row
	Numbers    map[int]image.Point
}

// NewBingo parses a rectangular table of newline and whitespace separated
// numbers.
func NewBingo(lines []string) (Bingo, error) {
	var b Bingo
	b.DimY = len(lines)
	b.Numbers = make(map[int]image.Point)
	for y, line := range lines {
		parts := strings.Fields(line)
		b.DimX = len(parts)
		for x := 0; x < b.DimX; x++ {
			n, err := strconv.Atoi(parts[x])
			if err != nil {
				return b, err
			}
			b.Numbers[n] = image.Point{x, y}
		}
	}
	b.Drawn = make(map[int]bool)
	b.Cols = make([]int, b.DimX)
	b.Rows = make([]int, b.DimY)

	return b, nil
}

// Draw optionally crosses a number and returns true if Bingo!.
func (a *Bingo) Draw(n int) bool {
	a.Drawn[n] = true

	c, ok := a.Numbers[n]
	if !ok {
		return false
	}
	a.Rows[c.X]++
	if a.Rows[c.X] == a.DimX {
		return true
	}
	a.Cols[c.Y]++
	if a.Cols[c.Y] == a.DimY {
		return true
	}
	return false
}

// Unmarked returns a list of unmarked numbers.
func (a Bingo) Unmarked() map[int]bool {
	m := make(map[int]bool)

	// list of all numbers ...
	for k := range a.Numbers {
		m[k] = true
	}

	// ... minus drawn...
	for k := range a.Drawn {
		delete(m, k)
	}

	// ... is list of unmarked
	return m
}

func Sum(m map[int]bool) (sum int) {
	for k := range m {
		sum += k
	}
	return
}

func Day04Part2(draws []int, boards []Bingo) int {
	active := make(map[int]Bingo, len(boards))
	for i := range boards {
		active[i] = boards[i]
	}
	for _, draw := range draws {
		for k, v := range active {
			bingo := v.Draw(draw)
			if bingo {
				delete(active, k)
				if len(active) == 0 {
					return Sum(v.Unmarked()) * draw
				}
			}
		}
	}
	return 0
}
