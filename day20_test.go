// day20_test.go
package adventofcode2021

import "testing"

func TestDay20Part1Example(t *testing.T) {
	// load the example input from testdata/day20_example.txt
	lines, err := linesFromFilename(exampleFilename(20))
	if err != nil {
		t.Fatalf("failed to read example file: %v", err)
	}
	tm, err := NewDay20(lines)
	if err != nil {
		t.Fatal(err)
	}
	got := Day20(tm, true) // part1
	want := 35
	if got != want {
		t.Fatalf("Part1 example: got %d, want %d", got, want)
	}
}

func TestDay20Part2Example(t *testing.T) {
	lines, err := linesFromFilename(exampleFilename(20))
	if err != nil {
		t.Fatalf("failed to read example file: %v", err)
	}
	tm, err := NewDay20(lines)
	if err != nil {
		t.Fatal(err)
	}
	got := Day20(tm, false) // part2
	want := 3351
	if got != want {
		t.Fatalf("Part2 example: got %d, want %d", got, want)
	}
}

func TestDay20Part1(t *testing.T) {
	const want = 5622
	lines, err := linesFromFilename(filename(20))
	if err != nil {
		t.Fatalf("failed to read input file: %v", err)
	}
	tm, err := NewDay20(lines)
	if err != nil {
		t.Fatal(err)
	}
	got := Day20(tm, true)
	if got != want {
		t.Fatalf("Part1: got %d, want %d", got, want)
	}
}

func TestDay20Part2(t *testing.T) {
	lines, err := linesFromFilename(filename(20))
	if err != nil {
		t.Fatalf("failed to read input file: %v", err)
	}
	tm, err := NewDay20(lines)
	if err != nil {
		t.Fatal(err)
	}
	got := Day20(tm, false)
	want := /* YOUR_PART2_ANSWER */ 0
	if got != want {
		t.Fatalf("Part2: got %d, want %d", got, want)
	}
}

func BenchmarkDay20Part1(b *testing.B) {
	lines, _ := linesFromFilename(filename(20))
	b.ResetTimer()
	for range b.N {
		tm, _ := NewDay20(lines)
		_ = Day20(tm, true)
	}
}

func BenchmarkDay20Part2(b *testing.B) {
	lines, _ := linesFromFilename(filename(20))
	b.ResetTimer()
	for range b.N {
		tm, _ := NewDay20(lines)
		_ = Day20(tm, false)
	}
}
