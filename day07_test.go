package adventofcode2021

import "testing"

func day07(t *testing.T, filename string, part1 bool, want int) {
	lines, err := linesFromFilename(filename)
	if err != nil {
		t.Fatal(err)
	}
	is, err := ParseCommaSeparatedNumbers(lines[0])
	if err != nil {
		t.Fatal(err)
	}
	got := Day07(is, part1)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay07Part1Example(t *testing.T) {
	// "This costs a total of 37 fuel."
	day07(t, exampleFilename(7), true, 37)
}

func TestDay07Part1(t *testing.T) {
	day07(t, filename(7), true, 326132)
}

func TestDay07Part2Example(t *testing.T) {
	day07(t, exampleFilename(7), false, 168)
}

func TestDay07Part2(t *testing.T) {
	day07(t, filename(7), false, 88612508)
}

func BenchmarkDay07Part1(b *testing.B) {
	lines, err := linesFromFilename(filename(7))
	if err != nil {
		b.Fatal(err)
	}

	is, err := ParseCommaSeparatedNumbers(lines[0])
	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()
	for range b.N {
		_ = Day07(is, true)
	}
}

func BenchmarkDay07Part2(b *testing.B) {
	lines, err := linesFromFilename(filename(7))
	if err != nil {
		b.Fatal(err)
	}

	is, err := ParseCommaSeparatedNumbers(lines[0])
	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()
	for range b.N {
		_ = Day07(is, false)
	}
}
