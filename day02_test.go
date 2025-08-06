package adventofcode2021

import (
	"testing"
)

// 6:30 - 7:11

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
	// "After following these instructions, you would have a horizontal position of 15 and a depth of 10. (Multiplying these together produces 150.)"
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

func BenchmarkDay02Part1(b *testing.B) {
	lines, err := linesFromFilename(filename(2))
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for range b.N {
		_, _ = Day02(lines, true)
	}
}

func BenchmarkDay02Part2(b *testing.B) {
	lines, err := linesFromFilename(filename(2))
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for range b.N {
		_, _ = Day02(lines, false)
	}
}
