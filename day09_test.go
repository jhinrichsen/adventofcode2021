package adventofcode2021

import "testing"

func day09(t *testing.T, filename string, part1 bool, want uint) {
	lines, err := linesFromFilename(filename)
	if err != nil {
		t.Fatal(err)
	}
	got, err := Day09(lines, part1)
	if err != nil {
		t.Fatal(err)
	}
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay09Part1Example(t *testing.T) {
	// "The sum of the risk levels of all low points in the heightmap is therefore 15"
	day09(t, exampleFilename(9), true, 15)
}

func TestDay09Part1(t *testing.T) {
	day09(t, filename(9), true, 514)
}

func TestDay09Part2Example(t *testing.T) {
	// "The product of the sizes of the three largest basins is 1134"
	day09(t, exampleFilename(9), false, 1134)
}

func TestDay09Part2(t *testing.T) {
	day09(t, filename(9), false, 1103130)
}

func BenchmarkDay09Part1(b *testing.B) {
	lines, err := linesFromFilename(filename(9))
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for range b.N {
		_, _ = Day09(lines, true)
	}
}

func BenchmarkDay09Part2(b *testing.B) {
	lines, err := linesFromFilename(filename(9))
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for range b.N {
		_, _ = Day09(lines, false)
	}
}
