package adventofcode2021

import (
	"testing"
)

func TestDay25Part1Example(t *testing.T) {
	const want = 58
	lines, err := linesFromFilename(example1Filename(25))
	if err != nil {
		t.Fatal(err)
	}

	got := Day25(lines, true)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay25Part1(t *testing.T) {
	const want = 300
	lines, err := linesFromFilename(filename(25))
	if err != nil {
		t.Fatal(err)
	}

	got := Day25(lines, true)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}
