package adventofcode2021

import (
	"testing"
)

func TestDay21Part1Example(t *testing.T) {
	const want = 739785
	lines, err := linesFromFilename(example1Filename(21))
	if err != nil {
		t.Fatal(err)
	}

	got := Day21(lines, true)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay21Part1(t *testing.T) {
	const want = 428736
	lines, err := linesFromFilename(filename(21))
	if err != nil {
		t.Fatal(err)
	}

	got := Day21(lines, true)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}
