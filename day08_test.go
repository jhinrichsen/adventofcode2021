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
	day08(t, exampleFilename(8), true, 26)
}

func TestDay08Part1(t *testing.T) {
	day08(t, filename(8), true, 255)
}

func TestDay08Part2Example(t *testing.T) {
	day08(t, exampleFilename(8), false, 61229)
}

/*
func TestDay08Part2(t *testing.T) {
	day08(t, filename(8), false, 0)
}

func BenchmarkDay08Part2(b *testing.B) {
	lines, err := linesFromFilename(filename(8))
	if err != nil {
		b.Fatal(err)
	}

	is, err := ParseCommaSeparatedNumbers(lines[0])
	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Day08(is, false)
	}
}
*/
