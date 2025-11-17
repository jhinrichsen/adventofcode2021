package adventofcode2021

import (
	"testing"
)

func TestDay24Part1(t *testing.T) {
	const want = 39924989499969
	lines, err := linesFromFilename(filename(24))
	if err != nil {
		t.Fatal(err)
	}

	got := Day24(lines, true)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay24Part2(t *testing.T) {
	const want = 16811412161117
	lines, err := linesFromFilename(filename(24))
	if err != nil {
		t.Fatal(err)
	}

	got := Day24(lines, false)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}
