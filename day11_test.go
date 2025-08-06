package adventofcode2021

import "testing"

func day11(t *testing.T, filename string, part1 bool, want uint) {
	lines, err := linesFromFilename(filename)
	if err != nil {
		t.Fatal(err)
	}
	data := NewDay11(lines)
	got := Day11(data, part1)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func bench11(b *testing.B, part1 bool) {
	b.Helper()
	lines, err := linesFromFilename(filename(11))
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		data := NewDay11(lines)
		Day11(data, part1)
	}
}

func BenchmarkDay11Part1(b *testing.B) {
	bench11(b, true)
}

func BenchmarkDay11Part2(b *testing.B) {
	bench11(b, false)
}

func TestDay11Part1Example(t *testing.T) {
	day11(t, exampleFilename(11), true, 1656)
}

func TestDay11Part1(t *testing.T) {
	day11(t, filename(11), true, 1749)
}

func TestDay11Part2Example(t *testing.T) {
	day11(t, exampleFilename(11), false, 195)
}

func TestDay11Part2(t *testing.T) {
	day11(t, filename(11), false, 285)
}
