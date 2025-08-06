package adventofcode2021

import (
	"testing"
)

// TestDay12Part1 tests the actual puzzle input for Part 1
func TestDay12Part1(t *testing.T) {
	lines, err := linesFromFilename(filename(12))
	if err != nil {
		t.Fatal(err)
	}

	connections := NewDay12(lines)

	got := Day12(connections, true)
	// Expected result for Part 1 with the given input
	want := uint(4775)
	if got != want {
		t.Errorf("Day12() part1 = %d, want %d", got, want)
	}
}
func TestDay12Part1Example1(t *testing.T) {
	lines, err := linesFromFilename(example1Filename(12))
	if err != nil {
		t.Fatal(err)
	}

	connections := NewDay12(lines)

	got := Day12(connections, true)
	want := uint(10)
	if got != want {
		t.Errorf("Day12Part1() = %d, want %d", got, want)
	}
}

func TestDay12Part1Example2(t *testing.T) {
	lines, err := linesFromFilename(example2Filename(12))
	if err != nil {
		t.Fatal(err)
	}

	connections := NewDay12(lines)

	got := Day12(connections, true)
	want := uint(19)
	if got != want {
		t.Errorf("Day12Part1() = %d, want %d", got, want)
	}
}

func TestDay12Part1Example3(t *testing.T) {
	lines, err := linesFromFilename(example3Filename(12))
	if err != nil {
		t.Fatal(err)
	}

	connections := NewDay12(lines)

	got := Day12(connections, true)
	want := uint(226)
	if got != want {
		t.Errorf("Day12Part1() = %d, want %d", got, want)
	}
}

func TestDay12Part2Example1(t *testing.T) {
	lines, err := linesFromFilename(example1Filename(12))
	if err != nil {
		t.Fatal(err)
	}

	connections := NewDay12(lines)

	got := Day12(connections, false)
	want := uint(36)
	if got != want {
		t.Errorf("Day12Part2() = %d, want %d", got, want)
	}
}

func TestDay12Part2Example2(t *testing.T) {
	lines, err := linesFromFilename(example2Filename(12))
	if err != nil {
		t.Fatal(err)
	}

	connections := NewDay12(lines)

	got := Day12(connections, false)
	want := uint(103)
	if got != want {
		t.Errorf("Day12Part2() = %d, want %d", got, want)
	}
}

func TestDay12Part2Example3(t *testing.T) {
}

func TestDay12Part2(t *testing.T) {
	lines, err := linesFromFilename(filename(12))
	if err != nil {
		t.Fatal(err)
	}

	connections := NewDay12(lines)

	got := Day12(connections, false)
	// Expected result for Part 2 with the given input
	want := uint(152480)
	if got != want {
		t.Errorf("Day12() part2 = %d, want %d", got, want)
	}
}

func BenchmarkDay12Part1(b *testing.B) {
	lines, err := linesFromFilename(filename(12))
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for range b.N {
		connections := NewDay12(lines)
		_ = Day12(connections, true)
	}
}

func BenchmarkDay12Part2(b *testing.B) {
	lines, err := linesFromFilename(filename(12))
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for range b.N {
		connections := NewDay12(lines)
		_ = Day12(connections, false)
	}
}
