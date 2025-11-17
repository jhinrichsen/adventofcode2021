package adventofcode2021

import (
	"testing"
)

func TestDay22Part1Example1(t *testing.T) {
	const want = 39
	lines, err := linesFromFilename(example1Filename(22))
	if err != nil {
		t.Fatal(err)
	}

	got := Day22(lines, true)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay22Part1Example2(t *testing.T) {
	const want = 590784
	lines, err := linesFromFilename(example2Filename(22))
	if err != nil {
		t.Fatal(err)
	}

	got := Day22(lines, true)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay22Part1(t *testing.T) {
	const want = 588200
	lines, err := linesFromFilename(filename(22))
	if err != nil {
		t.Fatal(err)
	}

	got := Day22(lines, true)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay22Part2Example(t *testing.T) {
	const want = 2758514936282235
	lines, err := linesFromFilename(example3Filename(22))
	if err != nil {
		t.Fatal(err)
	}

	got := Day22(lines, false)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay22Part2(t *testing.T) {
	const want = 1207167990362099
	lines, err := linesFromFilename(filename(22))
	if err != nil {
		t.Fatal(err)
	}

	got := Day22(lines, false)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func BenchmarkDay22Part1(b *testing.B) {
	lines, err := linesFromFilename(filename(22))
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for range b.N {
		_ = Day22(lines, true)
	}
}

func BenchmarkDay22Part2(b *testing.B) {
	lines, err := linesFromFilename(filename(22))
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for range b.N {
		_ = Day22(lines, false)
	}
}
