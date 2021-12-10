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

/*
func TestDay09Part1Example(t *testing.T) {
	day09(t, exampleFilename(9), true, 15)
}
*/

func TestDay09Part1(t *testing.T) {
	day09(t, filename(9), true, 514)
}

/*
func TestDay09Part2Example(t *testing.T) {
	day09(t, exampleFilename(9), false, 61229)
}

func TestDay09Part2(t *testing.T) {
	day09(t, filename(9), false, 0)
}

func BenchmarkDay09Part2(b *testing.B) {
	lines, err := linesFromFilename(filename(9))
	if err != nil {
		b.Fatal(err)
	}

	is, err := ParseCommaSeparatedNumbers(lines[0])
	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Day09(is, false)
	}
}
*/
