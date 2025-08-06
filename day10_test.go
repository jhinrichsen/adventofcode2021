package adventofcode2021

import "testing"

func day10(t *testing.T, filename string, part1 bool, want uint) {
	lines, err := linesFromFilename(filename)
	if err != nil {
		t.Fatal(err)
	}
	got := Day10(lines, part1)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay10Part1Example(t *testing.T) {
	day10(t, exampleFilename(10), true, 26397)
}

func TestDay10Part1(t *testing.T) {
	day10(t, filename(10), true, 265527)
}

func TestDay10Part2Example(t *testing.T) {
	day10(t, exampleFilename(10), false, 288957)
}

func TestDay10Part2(t *testing.T) {
	day10(t, filename(10), false, 3969823589)
}

func bench10(b *testing.B, part1 bool) {
	lines, err := linesFromFilename(filename(10))
	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Day10(lines, part1)
	}
}

func BenchmarkDay10Part1(b *testing.B) {
	bench10(b, true)
}

func BenchmarkDay10Part2(b *testing.B) {
	bench10(b, false)
}
