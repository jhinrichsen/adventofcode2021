package adventofcode2021

import "testing"

func day05(t *testing.T, filename string, part1 bool, want int) {
	lines, err := linesFromFilename(filename)
	if err != nil {
		t.Fatal(err)
	}
	got, err := Day05(lines, part1)
	if err != nil {
		t.Fatal(err)
	}
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay05Part1Example(t *testing.T) {
	day05(t, exampleFilename(5), true, 5)
}

func TestDay05Part1(t *testing.T) {
	day05(t, filename(5), true, 5632)
}

func TestDay05Part2Example(t *testing.T) {
	day05(t, exampleFilename(5), false, 12)
}

func TestDay05Part2(t *testing.T) {
	day05(t, filename(5), false, 22213)
}

func BenchmarkDay05Part2(b *testing.B) {
	lines, err := linesFromFilename(filename(5))
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = Day05(lines, false)
	}
}
