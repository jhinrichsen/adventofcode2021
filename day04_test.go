package adventofcode2021

import (
	"testing"
)

func day04(t *testing.T, filename string, part1 bool, want int) {
	lines, err := linesFromFilename(filename)
	if err != nil {
		t.Fatal(err)
	}

	draws, boards, err := NewDay04(lines)
	if err != nil {
		t.Fatal(err)
	}

	var f func([]int, []Bingo) int
	if part1 {
		f = Day04Part1
	} else {
		f = Day04Part2
	}
	got := f(draws, boards)
	if want != got {
		t.Fatalf("want %d but got %d\n", want, got)
	}
}

func TestDay04Part1Example(t *testing.T) {
	// "the final score, 188 * 24 = 4512."
	day04(t, exampleFilename(4), true, 4512)
}

func TestDay04Part1(t *testing.T) {
	day04(t, filename(4), true, 33462)
}

func TestDay04Part2Example(t *testing.T) {
	day04(t, exampleFilename(4), false, 1924)
}

func TestDay04Part2(t *testing.T) {
	day04(t, filename(4), false, 30070)
}

func BenchmarkDay04Part2(b *testing.B) {
	lines, err := linesFromFilename(filename(4))
	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		draws, boards, err := NewDay04(lines)
		if err != nil {
			b.Fatal(err)
		}
		_ = Day04Part2(draws, boards)
	}
}
