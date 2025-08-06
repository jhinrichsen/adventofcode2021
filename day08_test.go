package adventofcode2021

import "testing"

func day08(t *testing.T, filename string, part1 bool, want uint) {
	lines, err := linesFromFilename(filename)
	if err != nil {
		t.Fatal(err)
	}
	got := Day08(lines, part1)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay08Part1Example(t *testing.T) {
	// "In the above example, there are 26 instances of digits that use a unique number of segments"
	day08(t, exampleFilename(8), true, 26)
}

func TestDay08Part1(t *testing.T) {
	day08(t, filename(8), true, 255)
}

func TestDay08Part2Example(t *testing.T) {
	day08(t, exampleFilename(8), false, 61229)
}

func TestDay08Part2(t *testing.T) {
	day08(t, filename(8), false, 982158)
}

func BenchmarkDay08Part1(b *testing.B) {
	lines, err := linesFromFilename(filename(8))
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for range b.N {
		_ = Day08(lines, true)
	}
}

func BenchmarkDay08Part2(b *testing.B) {
	lines, err := linesFromFilename(filename(8))
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for range b.N {
		_ = Day08(lines, false)
	}
}
