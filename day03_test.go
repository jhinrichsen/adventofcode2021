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
	// "Multiplying the gamma rate (22) by the epsilon rate (9) produces the power consumption, 198."
	day03(t, exampleFilename(3), true, 198)
}

func TestDay03Part1(t *testing.T) {
	day03(t, filename(3), true, 4138664)
}

func TestDay03Part2Example(t *testing.T) {
	day03(t, exampleFilename(3), false, 230)
}

func TestDay03Part2(t *testing.T) {
	day03(t, filename(3), false, 4273224)
}

func BenchmarkDay03Part1(b *testing.B) {
	lines, err := linesFromFilename(filename(3))
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for range b.N {
		_, _ = Day03(lines, true)
	}
}

func BenchmarkDay03Part2(b *testing.B) {
	lines, err := linesFromFilename(filename(3))
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for range b.N {
		_, _ = Day03(lines, false)
	}
}
