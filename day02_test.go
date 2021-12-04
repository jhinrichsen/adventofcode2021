package adventofcode2021

import (
	"testing"
)

// 6:30

func day02(t *testing.T, filename string, part1 bool, want int) {
	lines, err := linesFromFilename(filename)
	if err != nil {
		t.Fatal(err)
	}
	got, err := Day02(lines, part1)
	if err != nil {
		t.Fatal(err)
	}
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay02Part1Example(t *testing.T) {
	day02(t, exampleFilename(2), true, 150)
}

func TestDay02Part1(t *testing.T) {
	day02(t, filename(2), true, 1938402)
}

func TestDay02Part2Example(t *testing.T) {
	day02(t, exampleFilename(2), false, 900)
}

func TestDay02Part2(t *testing.T) {
	day02(t, filename(2), false, 1947878632)
}
