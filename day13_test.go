package adventofcode2021

import "testing"

func bench13(b *testing.B, part1 bool) {
	b.Helper()
	lines, err := linesFromFilename(filename(13))
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for range b.N {
		dots, folds := NewDay13(lines)
		Day13(dots, folds, part1)
	}
}

func BenchmarkDay13Part1(b *testing.B) {
	bench13(b, true)
}

// TestDay13Part1 tests the actual puzzle input for Part 1
func TestDay13Part1(t *testing.T) {
	lines, err := linesFromFilename(filename(13))
	if err != nil {
		t.Fatal(err)
	}

	dots, folds := NewDay13(lines)
	got := Day13(dots, folds, true) // Part 1: only first fold
	// Expected result for Part 1 with the given input
	want := uint(729)
	if got != want {
		t.Errorf("Day13() part1 = %d, want %d", got, want)
	}
}

func TestDay13Example1(t *testing.T) {
	// Example from puzzle description
	lines, err := linesFromFilename("testdata/day13_example.txt")
	if err != nil {
		t.Fatal(err)
	}

	dots, folds := NewDay13(lines)
	got := Day13(dots, folds, true) // Part 1: only first fold
	want := uint(17)                // After first fold, 17 dots are visible
	if got != want {
		t.Errorf("Day13() part1 = %d, want %d", got, want)
	}
}

func TestDay13Example2(t *testing.T) {
	// Example from puzzle description - Part 2 should apply all folds
	lines, err := linesFromFilename("testdata/day13_example.txt")
	if err != nil {
		t.Fatal(err)
	}

	dots, folds := NewDay13(lines)
	got := Day13(dots, folds, false) // Part 2: all folds
	// After both folds, we get a 5x7 square pattern
	// The puzzle shows it forms a square, let's count the dots in the final pattern
	// Looking at the final pattern: #####, #...#, #...#, #...#, #####, ....., .....
	// That's 5+2+2+2+5 = 16 dots
	want := uint(16)
	if got != want {
		t.Errorf("Day13() part2 = %d, want %d", got, want)
	}
}
