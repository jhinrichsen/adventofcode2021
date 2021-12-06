package adventofcode2021

import "testing"

func day06(t *testing.T, filename string, days int, want uint) {
	lines, err := linesFromFilename(filename)
	if err != nil {
		t.Fatal(err)
	}
	got, err := Day06(lines, days)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay06Part1Example(t *testing.T) {
	day06(t, exampleFilename(6), 80, 5934)
}

func TestDay06Part1(t *testing.T) {
	day06(t, filename(6), 80, 362639)
}

func TestDay06Part2Example(t *testing.T) {
	day06(t, exampleFilename(6), 256, 26984457539)
}

func TestDay06Part2(t *testing.T) {
	day06(t, filename(6), 256, 1639854996917)
}

func BenchmarkDay06Part2(b *testing.B) {
	lines, err := linesFromFilename(filename(6))
	if err != nil {
		b.Fatal(err)
	}
	for i := 0; i < b.N; i++ {
		_, _ = Day06(lines, 256)
	}
}
