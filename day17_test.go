package adventofcode2021

import (
	"testing"
)

func TestDay17Part1Example(t *testing.T) {
    const want = 45
    lines, err := linesFromFilename(exampleFilename(17))
    if err != nil {
        t.Fatal(err)
    }

    data := NewDay17(lines)
    got := Day17(data, true)
    if want != got {
        t.Fatalf("want %d but got %d", want, got)
    }
}

func TestDay17Part2Example(t *testing.T) {
    // "In this example, there are 112 distinct initial velocity values that cause the probe to be within the target area after any step."
    const want = 112
    lines, err := linesFromFilename(exampleFilename(17))
    if err != nil {
        t.Fatal(err)
    }
    data := NewDay17(lines)
    got := Day17(data, false)
    if want != got {
        t.Fatalf("want %d but got %d", want, got)
    }
}

func TestDay17Part1(t *testing.T) {
	const want = 7750
	lines, err := linesFromFilename(filename(17))
	if err != nil {
		t.Fatal(err)
	}
	data := NewDay17(lines)
	got := Day17(data, true)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay17Part2(t *testing.T) {
	const want = 4120
	lines, err := linesFromFilename(filename(17))
	if err != nil {
		t.Fatal(err)
	}
	data := NewDay17(lines)
	got := Day17(data, false)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func BenchmarkDay17Part1(b *testing.B) {
	lines, err := linesFromFilename(filename(17))
	if err != nil {
		b.Fatal(err)
	}
	data := NewDay17(lines)
	b.ResetTimer()
	for range b.N {
		_ = Day17(data, true)
	}
}

func BenchmarkDay17Part2(b *testing.B) {
	lines, err := linesFromFilename(filename(17))
	if err != nil {
		b.Fatal(err)
	}
	data := NewDay17(lines)
	b.ResetTimer()
	for range b.N {
		_ = Day17(data, false)
	}
}
