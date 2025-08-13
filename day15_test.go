package adventofcode2021

import "testing"

func day15(t *testing.T, filename string, part1 bool, want uint) {
	lines, err := linesFromFilename(filename)
	if err != nil {
		t.Fatal(err)
	}
	got, err := Day15(lines, part1)
	if err != nil {
		t.Fatal(err)
	}
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay15Part1Example(t *testing.T) {
	// "In the above example, the lowest total risk is 40."
	day15(t, exampleFilename(15), true, 40)
}

func TestDay15Part1(t *testing.T) {
	// Expected value should be derived from running the solution on the actual input once.
	// We'll compute it and then lock it in here.
	day15(t, filename(15), true, 435)
}

func BenchmarkDay15Part1(b *testing.B) {
	lines, err := linesFromFilename(filename(15))
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for range b.N {
		_, _ = Day15(lines, true)
	}
}

func TestDay15Part2Example(t *testing.T) {
	// "The total risk of this path is 315 (the starting position is still never entered, so its risk is not counted)."
	day15(t, exampleFilename(15), false, 315)
}

func TestDay15Part2(t *testing.T) {
	// Expected value should be derived from running the solution on the actual input once.
	// We'll compute it and then lock it in here.
	day15(t, filename(15), false, 2842)
}

func BenchmarkDay15Part2(b *testing.B) {
	lines, err := linesFromFilename(filename(15))
	if err != nil {
		b.Fatal(err)
	}
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if _, err := Day15(lines, false); err != nil {
			b.Fatal(err)
		}
	}
}
