package adventofcode2021

import (
	"testing"
)

func TestDay23Part1Example(t *testing.T) {
	const want = 12521
	lines, err := linesFromFilename(example1Filename(23))
	if err != nil {
		t.Fatal(err)
	}

	got := Day23(lines, true)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay23Part1(t *testing.T) {
	const want = 18051
	lines, err := linesFromFilename(filename(23))
	if err != nil {
		t.Fatal(err)
	}

	got := Day23(lines, true)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay23Part2Example(t *testing.T) {
	const want = 44169
	lines, err := linesFromFilename(example1Filename(23))
	if err != nil {
		t.Fatal(err)
	}

	got := Day23(lines, false)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay23Part2(t *testing.T) {
	const want = 50245
	lines, err := linesFromFilename(filename(23))
	if err != nil {
		t.Fatal(err)
	}

	got := Day23(lines, false)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func BenchmarkDay23Part1(b *testing.B) {
	lines, err := linesFromFilename(filename(23))
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for range b.N {
		_ = Day23(lines, true)
	}
}

func BenchmarkDay23Part2(b *testing.B) {
	lines, err := linesFromFilename(filename(23))
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for range b.N {
		_ = Day23(lines, false)
	}
}
