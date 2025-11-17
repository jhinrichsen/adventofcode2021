package adventofcode2021

import (
	"testing"
)

func TestDay19Part1Example(t *testing.T) {
	const want = 79
	lines, err := linesFromFilename(example1Filename(19))
	if err != nil {
		t.Fatal(err)
	}

	got := Day19(lines, true)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay19Part1(t *testing.T) {
	const want = 425
	lines, err := linesFromFilename(filename(19))
	if err != nil {
		t.Fatal(err)
	}

	got := Day19(lines, true)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay19Part2Example(t *testing.T) {
	const want = 3621
	lines, err := linesFromFilename(example1Filename(19))
	if err != nil {
		t.Fatal(err)
	}

	got := Day19(lines, false)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay19Part2(t *testing.T) {
	const want = 13354
	lines, err := linesFromFilename(filename(19))
	if err != nil {
		t.Fatal(err)
	}

	got := Day19(lines, false)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func BenchmarkDay19Part1(b *testing.B) {
	lines, err := linesFromFilename(filename(19))
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for range b.N {
		_ = Day19(lines, true)
	}
}

func BenchmarkDay19Part2(b *testing.B) {
	lines, err := linesFromFilename(filename(19))
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for range b.N {
		_ = Day19(lines, false)
	}
}
