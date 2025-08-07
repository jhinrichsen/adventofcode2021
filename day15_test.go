package adventofcode2021

import "testing"

func day15(t *testing.T, filename string, part1 bool, want uint) {
	lines, err := linesFromFilename(filename)
	if err != nil {
		t.Fatal(err)
	}
	got, err := Day15(lines, part1)
	if err != nil {
		t.Fatal(err)
	}
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay15Part1Example(t *testing.T) {
	// "In the above example, the lowest total risk is 40."
	day15(t, exampleFilename(15), true, 40)
}

func TestDay15Part1(t *testing.T) {
	// Expected value should be derived from running the solution on the actual input once.
	// We'll compute it and then lock it in here.
	day15(t, filename(15), true, 435)
}
