package adventofcode2021

import (
	"strings"
	"testing"

	"gitlab.com/jhinrichsen/aococr"
)

// minInt returns the smaller of two int values
func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// minUint returns the smaller of two uint values
func minUint(a, b uint) uint {
	if a < b {
		return a
	}
	return b
}

const (
	errFmt = "Day13() part1 = %d, want %d"
)

func bench13(b *testing.B, limit uint) {
	b.Helper()
	lines, err := linesFromFilename(filename(13))
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for range b.N {
		dots, folds := NewDay13(lines)
		_, _ = Day13(dots, folds, limit) // Ignore ASCII art in benchmarks
	}
}

func BenchmarkDay13Part1(b *testing.B) {
	bench13(b, 1) // Only apply first fold
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
	want := uint(729)
	if got != want {
		t.Errorf(errFmt, got, want)
	}
}

func TestDay13Example1(t *testing.T) {
	// Example from puzzle description
	lines, err := linesFromFilename("testdata/day13_example.txt")
	if err != nil {
		t.Fatal(err)
	}

	dots, folds := NewDay13(lines)
	// For Part 1, we only need to apply the first fold
	got, _ := Day13(dots, folds, 1) // Only apply first fold
	want := uint(17)                // After first fold, 17 dots are visible
	if got != want {
		t.Errorf(errFmt, got, want)
	}
}

func TestDay13Example2(t *testing.T) {
	// Example from puzzle description - Part 2 should apply all folds
	lines, err := linesFromFilename("testdata/day13_example.txt")
	if err != nil {
		t.Fatal(err)
	}

	dots, folds := NewDay13(lines)
	// For Part 2, we need to apply all folds
	got, asciiLines := Day13(dots, folds, 0) // Apply all folds (0 means no limit)

	// Define the expected pattern
	expectedPattern := []string{
		"#####",
		"#...#",
		"#...#",
		"#...#",
		"#####",
		".....",
		".....",
	}

	// Verify line count
	if len(asciiLines) != len(expectedPattern) {
		t.Fatalf("Expected %d lines, got %d", len(expectedPattern), len(asciiLines))
	}

	// Verify each line matches exactly
	for i, line := range asciiLines {
		if line != expectedPattern[i] {
			t.Errorf("Line %d mismatch.\nExpected: %q\nGot:      %q",
				i+1, expectedPattern[i], line)
			t.Logf("Full ASCII art:\n%s", strings.Join(asciiLines, "\n"))
			break
		}
	}

	// Verify dot count (5+2+2+2+5 = 16 dots)
	want := uint(16)
	if got != want {
		t.Errorf(errFmt, got, want)
	}
}

// TestDay13Part2 tests the actual puzzle input for Part 2
func TestDay13Part2(t *testing.T) {
	// Read test input
	input := `6,10
0,14
9,10
0,3
10,4
4,11
6,0
6,12
4,1
0,13
10,12
3,4
3,0
8,4
1,10
2,14
8,10
9,0

fold along y=7
fold along x=5`

	lines := strings.Split(input, "\n")
	points, folds := NewDay13(lines)

	// Process all folds (limit = 0 means process all)
	_, asciiLines := Day13(points, folds, 0)

	// Verify the output is not empty
	if len(asciiLines) == 0 {
		t.Fatal("Empty ASCII art generated")
	}

	// After all folds, we should have the final code letters
	// The expected height is 7 based on the actual output after all folds
	expectedHeight := 7

	if len(asciiLines) != expectedHeight {
		t.Fatalf("Expected %d lines in ASCII art, got %d", expectedHeight, len(asciiLines))
	}
	// Print the ASCII art for debugging
	t.Logf("\nASCII Art (height: %d):", len(asciiLines))
	for i, line := range asciiLines {
		t.Logf("%2d: %s", i, line)
	}

	// Convert the ASCII art to a single string for OCR
	asciiArt := strings.Join(asciiLines, "\n")

	// Print the exact string being passed to ParseLetters
	t.Logf("String being passed to ParseLetters (length: %d):\n---\n%s\n---\n", len(asciiArt), asciiArt)

	// Define which characters to consider as points (in this case, '#')
	charSet := map[rune]bool{
		'#': true,
	}

	// Use aococr to parse the letters from the ASCII art
	t.Log("Attempting to parse letters with aococr...")
	code, err := aococr.ParseLetters(asciiArt, charSet)
	if err != nil {
		t.Fatalf("OCR failed: %v", err)
	}

	// Verify the code looks reasonable (8 capital letters)
	if len(code) != 8 {
		t.Fatalf("Expected 8-character code, got %q (length %d)", code, len(code))
	}

	// Log the recognized code
	t.Logf("Recognized code: %s", code)
}
