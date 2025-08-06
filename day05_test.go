package adventofcode2021

import "testing"

func day05(t *testing.T, f func([]string, bool) (int, error), filename string, part1 bool, want int) {
	lines, err := linesFromFilename(filename)
	if err != nil {
		t.Fatal(err)
	}
	got, err := f(lines, part1)
	if err != nil {
		t.Fatal(err)
	}
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay05Part1Example(t *testing.T) {
	// "In the above example, this is anywhere in the diagram with a 2 or larger - a total of 5 points."
	day05(t, Day05Int, exampleFilename(5), true, 5)
	day05(t, Day05Cmplx, exampleFilename(5), true, 5)
}

func TestDay05Part1(t *testing.T) {
	day05(t, Day05Int, filename(5), true, 5632)
	day05(t, Day05Cmplx, filename(5), true, 5632)
}

func TestDay05Part2Example(t *testing.T) {
	day05(t, Day05Int, exampleFilename(5), false, 12)
	day05(t, Day05Cmplx, exampleFilename(5), false, 12)
}

func TestDay05Part2(t *testing.T) {
	day05(t, Day05Int, filename(5), false, 22213)
	day05(t, Day05Cmplx, filename(5), false, 22213)
}

func benchDay05(b *testing.B, f func([]string, bool) (int, error)) {
	lines, err := linesFromFilename(filename(5))
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = f(lines, false)
	}
}

func BenchmarkDay05Part2Int(b *testing.B) {
	benchDay05(b, Day05Int)
}

func BenchmarkDay05Part2Cmplx(b *testing.B) {
	benchDay05(b, Day05Cmplx)
}
