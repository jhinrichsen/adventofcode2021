package adventofcode2021

import (
	"testing"
)

func TestDay18Explode(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{"[[[[[9,8],1],2],3],4]", "[[[[0,9],2],3],4]"},
		{"[7,[6,[5,[4,[3,2]]]]]", "[7,[6,[5,[7,0]]]]"},
		{"[[6,[5,[4,[3,2]]]],1]", "[[6,[5,[7,0]]],3]"},
		{"[[3,[2,[1,[7,3]]]],[6,[5,[4,[3,2]]]]]", "[[3,[2,[8,0]]],[9,[5,[4,[3,2]]]]]"},
		{"[[3,[2,[8,0]]],[9,[5,[4,[3,2]]]]]", "[[3,[2,[8,0]]],[9,[5,[7,0]]]]"},
	}

	for _, tt := range tests {
		num, _ := parseSnailfish(tt.input)
		if tryExplode(num) {
			got := num.String()
			if got != tt.want {
				t.Errorf("explode(%s) = %s, want %s", tt.input, got, tt.want)
			}
		} else {
			t.Errorf("explode(%s) returned false, expected explosion", tt.input)
		}
	}
}

func TestDay18Split(t *testing.T) {
	tests := []struct {
		value int
		left  int
		right int
	}{
		{10, 5, 5},
		{11, 5, 6},
		{12, 6, 6},
	}

	for _, tt := range tests {
		num := &SnailfishNumber{value: tt.value, isRegular: true}
		parent := &SnailfishNumber{left: num, isRegular: false}

		if splitAny(parent) {
			if !parent.left.isRegular {
				leftVal := parent.left.left.value
				rightVal := parent.left.right.value
				if leftVal != tt.left || rightVal != tt.right {
					t.Errorf("split(%d) = [%d,%d], want [%d,%d]", tt.value, leftVal, rightVal, tt.left, tt.right)
				}
			}
		}
	}
}

func TestDay18Magnitude(t *testing.T) {
	tests := []struct {
		input string
		want  uint
	}{
		{"[9,1]", 29},
		{"[1,9]", 21},
		{"[[9,1],[1,9]]", 129},
		{"[[1,2],[[3,4],5]]", 143},
		{"[[[[0,7],4],[[7,8],[6,0]]],[8,1]]", 1384},
		{"[[[[1,1],[2,2]],[3,3]],[4,4]]", 445},
		{"[[[[3,0],[5,3]],[4,4]],[5,5]]", 791},
		{"[[[[5,0],[7,4]],[5,5]],[6,6]]", 1137},
		{"[[[[8,7],[7,7]],[[8,6],[7,7]]],[[[0,7],[6,6]],[8,7]]]", 3488},
	}

	for _, tt := range tests {
		num, _ := parseSnailfish(tt.input)
		got := magnitude(num)
		if got != tt.want {
			t.Errorf("magnitude(%s) = %d, want %d", tt.input, got, tt.want)
		}
	}
}

func TestDay18Addition(t *testing.T) {
	// Test the example from the puzzle description
	a, _ := parseSnailfish("[[[[4,3],4],4],[7,[[8,4],9]]]")
	b, _ := parseSnailfish("[1,1]")
	result := add(a, b)
	want := "[[[[0,7],4],[[7,8],[6,0]]],[8,1]]"
	got := result.String()
	if got != want {
		t.Errorf("add result = %s, want %s", got, want)
	}
}

func TestDay18Part1Example(t *testing.T) {
	const want = 4140
	lines, err := linesFromFilename(example1Filename(18))
	if err != nil {
		t.Fatal(err)
	}

	got := Day18(lines, true)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay18Part1(t *testing.T) {
	const want = 4145
	lines, err := linesFromFilename(filename(18))
	if err != nil {
		t.Fatal(err)
	}

	got := Day18(lines, true)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay18Part2Example(t *testing.T) {
	const want = 3993
	lines, err := linesFromFilename(example1Filename(18))
	if err != nil {
		t.Fatal(err)
	}

	got := Day18(lines, false)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay18Part2(t *testing.T) {
	const want = 4855
	lines, err := linesFromFilename(filename(18))
	if err != nil {
		t.Fatal(err)
	}

	got := Day18(lines, false)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func BenchmarkDay18Part1(b *testing.B) {
	lines, err := linesFromFilename(filename(18))
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for range b.N {
		_ = Day18(lines, true)
	}
}

func BenchmarkDay18Part2(b *testing.B) {
	lines, err := linesFromFilename(filename(18))
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for range b.N {
		_ = Day18(lines, false)
	}
}
