package adventofcode2021

import (
	"testing"
)

func TestDay17Part1Example(t *testing.T) {
	const want = 45
	lines, err := linesFromFilename(exampleFilename(17))
	if err != nil {
		t.Fatal(err)
	}
	data := NewDay17(lines)
	got := Day17(data, true)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay17Part1(t *testing.T) {
	const want = 7750
	lines, err := linesFromFilename(filename(17))
	if err != nil {
		t.Fatal(err)
	}
	data := NewDay17(lines)
	got := Day17(data, true)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}
