package adventofcode2021

import (
	"strings"
	"testing"

	"gitlab.com/jhinrichsen/aococr"
)

func TestDay13Part1Example(t *testing.T) {
	const want = 17
	lines, err := linesFromFilename("testdata/day13_example.txt")
	if err != nil {
		t.Fatal(err)
	}

	dots, folds := NewDay13(lines)
	got, _ := Day13(dots, folds, 1)
	if got != want {
		t.Errorf("Day13() part1 = %d, want %d", got, want)
	}
}

func TestDay13Part2Example(t *testing.T) {
	want := []string{
		"#####",
		"#...#",
		"#...#",
		"#...#",
		"#####",
		".....",
		".....",
	}

	lines, err := linesFromFilename("testdata/day13_example.txt")
	if err != nil {
		t.Fatal(err)
	}

	dots, folds := NewDay13(lines)
	_, got := Day13(dots, folds, 0)

	if len(want) != len(got) {
		t.Fatalf("want %d lines, got %d", len(want), len(got))
	}

	for i := range want {
		if want[i] != got[i] {
			t.Errorf("line %d mismatch.\nwant: %q\ngot:  %q", i+1, want[i], got[i])
			break
		}
	}
}

// TestDay13Part1 tests the actual puzzle input for Part 1
func TestDay13Part1(t *testing.T) {
	lines, err := linesFromFilename(filename(13))
	if err != nil {
		t.Fatal(err)
	}

	dots, folds := NewDay13(lines)
	// For Part 1, we only need to apply the first fold
	got, _ := Day13(dots, folds, 1) // Only apply first fold
	// Expected result for Part 1 with the given input
	const want = 729
	if got != want {
		t.Errorf("Day13() part1 = %d, want %d", got, want)
	}
}

// TestDay13Part2 tests the actual puzzle input for Part 2
func TestDay13Part2(t *testing.T) {
	const want = "RGZLBHFP"
	lines, err := linesFromFilename(filename(13))
	if err != nil {
		t.Fatal(err)
	}
	dots, folds := NewDay13(lines)
	_, ascii := Day13(dots, folds, 0)

	// run OCR against result
	got, err := aococr.ParseLetters(strings.Join(ascii, "\n"), map[rune]bool{'#': true})
	if err != nil {
		t.Fatalf("OCR failed on real input: %v", err)
	}
	if want != got {
		t.Fatalf("unexpected code: want %q, got %q", want, got)
	}
}

func BenchmarkDay13Part1(b *testing.B) {
	bench13(b, 1) // Only apply first fold
}

func BenchmarkDay13Part2(b *testing.B) {
	bench13(b, 0) // Apply all folds and render
}

func bench13(b *testing.B, limit uint) {
	b.Helper()
	lines, err := linesFromFilename(filename(13))
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		dots, folds := NewDay13(lines)
		_, _ = Day13(dots, folds, limit)
	}
}
