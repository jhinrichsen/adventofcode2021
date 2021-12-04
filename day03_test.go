package adventofcode2021

import "testing"

func day03(t *testing.T, filename string, part1 bool, want int) {
	lines, err := linesFromFilename(filename)
	if err != nil {
		t.Fatal(err)
	}
	got, err := Day03(lines, part1)
	if err != nil {
		t.Fatal(err)
	}
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay03Part1Example(t *testing.T) {
	day03(t, exampleFilename(3), true, 198)
}

func TestDay03Part1(t *testing.T) {
	day03(t, filename(3), true, 4138664)
}
