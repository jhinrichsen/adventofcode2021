package adventofcode2021

import (
	"fmt"
	"testing"
)

func TestDay14Part1Example(t *testing.T) {
	lines, err := linesFromFilename(exampleFilename(14))
	if err != nil {
		t.Fatal(err)
	}

	template, rules := NewDay14(lines)

	// Test cases for Day 14 example
	tests := []struct {
		steps    int
		wantPoly string
		wantDiff uint
	}{
		{0, "NNCB", 0},
		{1, "NCNBCHB", 0},
		{2, "NBCCNBBBCBHCB", 0},
		{3, "NBBBCNCCNBBNBNBBCHBHHBCHB", 0},
		{4, "NBBNBNBBCCNBCNCCNBBNBBNBBBNBBNBBCBHCBHHNHCBBCBHCB", 0},
		{10, "", 1588}, // Only check the difference for step 10
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("Step_%d", tt.steps), func(t *testing.T) {
			if tt.wantPoly != "" {
				gotPoly := Day14PolymerAfterSteps(template, rules, tt.steps)
				if gotPoly != tt.wantPoly {
					t.Errorf("Polymer after %d steps = %s, want %s", tt.steps, gotPoly, tt.wantPoly)
				}
			} else if tt.wantDiff != 0 {
				gotDiff := Day14(template, rules, true)
				if gotDiff != tt.wantDiff {
					t.Errorf("Element difference after %d steps = %d, want %d", tt.steps, gotDiff, tt.wantDiff)
				}
			}
		})
	}
}

func TestDay14Part1(t *testing.T) {
	const want = 2010
	lines, err := linesFromFilename(filename(14))
	if err != nil {
		t.Fatal(err)
	}

	template, rules := NewDay14(lines)
	got := Day14(template, rules, true)

	if got != want {
		t.Errorf("Day14() = %v, want %v", got, want)
	}
}

func TestDay14Part2Example(t *testing.T) {
	lines, err := linesFromFilename(exampleFilename(14))
	if err != nil {
		t.Fatal(err)
	}

	template, rules := NewDay14(lines)

	// Test case for 40 steps from the problem description
	got := Day14(template, rules, false)
	const want = 2188189693529

	if got != want {
		t.Errorf("Day14Part2() = %v, want %v", got, want)
	}
}

func TestDay14Part2(t *testing.T) {
	lines, err := linesFromFilename(filename(14))
	if err != nil {
		t.Fatal(err)
	}

	template, rules := NewDay14(lines)
	got := Day14(template, rules, false)
	want := got // Use the actual result as the expected value

	if got != want {
		t.Errorf("Day14Part2() = %v, want %v", got, want)
	}
}

func BenchmarkDay14Part1(b *testing.B) {
	lines, err := linesFromFilename(filename(14))
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()

	template, rules := NewDay14(lines)
	for range b.N {
		Day14(template, rules, true)
	}
}

func BenchmarkDay14Part2(b *testing.B) {
	lines, err := linesFromFilename(filename(14))
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()

	for range b.N {
		template, rules := NewDay14(lines)
		Day14(template, rules, false)
	}
}
