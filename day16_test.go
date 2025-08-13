package adventofcode2021

import (
	"strings"
	"testing"
)

// Shared examples table for Part 1 and Part 2
var day16Examples = []struct {
	name string
	hex  string
	want uint
}{
	// "With the first example, the hex string 8A004A801A8002F478 ... has a version sum of 16."
	{name: "part 1 example 1", hex: "8A004A801A8002F478", want: 16},
	// "The second example, 620080001611562C8802118E34, results in a version sum of 12."
	{name: "part 1 example 2", hex: "620080001611562C8802118E34", want: 12},
	// "The third example, C0015000016115A2E0802F182340, results in a version sum of 23."
	{name: "part 1 example 3", hex: "C0015000016115A2E0802F182340", want: 23},
	// "The fourth example, A0016C880162017C3686B18A3D4780, results in a version sum of 31."
	{name: "part 1 example 4", hex: "A0016C880162017C3686B18A3D4780", want: 31},

	// "C200B40A82 finds the sum of 1 and 2, resulting in the value 3."
	{name: "part 2 example 1", hex: "C200B40A82", want: 3},
	// "04005AC33890 finds the product of 6 and 9, resulting in the value 54."
	{name: "part 2 example 2", hex: "04005AC33890", want: 54},
	// "880086C3E88112 finds the minimum of 7, 8, and 9, resulting in the value 7."
	{name: "part 2 example 3", hex: "880086C3E88112", want: 7},
	// "CE00C43D881120 finds the maximum of 7, 8, and 9, resulting in the value 9."
	{name: "part 2 example 4", hex: "CE00C43D881120", want: 9},
	// "D8005AC2A8F0 produces 1, because 5 is less than 15."
	{name: "part 2 example 5", hex: "D8005AC2A8F0", want: 1},
	// "F600BC2D8F produces 0, because 5 is not greater than 15."
	{name: "part 2 example 6", hex: "F600BC2D8F", want: 0},
	// "9C005AC2F8F0 produces 0, because 5 is not equal to 15."
	{name: "part 2 example 7", hex: "9C005AC2F8F0", want: 0},
	// "9C0141080250320F1802104A08 produces 1, because 1 + 3 = 2 * 2."
	{name: "part 2 example 8", hex: "9C0141080250320F1802104A08", want: 1},
}

func TestDay16Part1Examples(t *testing.T) {
	for _, ex := range day16Examples {
		if !strings.HasPrefix(ex.name, "part 1") {
			continue
		}
		t.Run(ex.name, func(t *testing.T) {
			bits, err := NewDay16(ex.hex)
			if err != nil {
				t.Fatal(err)
			}
			got := Day16(bits, true)
			if got != ex.want {
				t.Fatalf("got %d, want %d", got, ex.want)
			}
		})
	}
}

func TestDay16Part1(t *testing.T) {
	const want = 891
	lines, err := linesFromFilename(filename(16))
	if err != nil {
		t.Fatal(err)
	}
	hexStr := lines[0]
	bits, err := NewDay16(hexStr)
	if err != nil {
		t.Fatal(err)
	}
	got := Day16(bits, true)
	if got != want {
		t.Fatalf("got %d, want %d", got, want)
	}
}

func TestDay16Part2Examples(t *testing.T) {
	for _, ex := range day16Examples {
		if !strings.HasPrefix(ex.name, "part 2") {
			continue
		}
		t.Run(ex.name, func(t *testing.T) {
			bits, err := NewDay16(ex.hex)
			if err != nil {
				t.Fatal(err)
			}
			got := Day16(bits, false)
			if got != ex.want {
				t.Fatalf("got %d, want %d", got, ex.want)
			}
		})
	}
}

func TestDay16Part2(t *testing.T) {
	const want = 673042777597
	lines, err := linesFromFilename(filename(16))
	if err != nil {
		t.Fatal(err)
	}
	bits, err := NewDay16(lines[0])
	if err != nil {
		t.Fatal(err)
	}
	got := Day16(bits, false)
	if got != want {
		t.Fatalf("got %d, want %d", got, want)
	}
}

func BenchmarkDay16Part1(b *testing.B) {
	lines, err := linesFromFilename(filename(16))
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for range b.N {
		bits, err := NewDay16(lines[0])
		if err != nil {
			b.Fatal(err)
		}
		_ = Day16(bits, true)
	}
}

func BenchmarkDay16Part2(b *testing.B) {
	lines, err := linesFromFilename(filename(16))
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for range b.N {
		bits, err := NewDay16(lines[0])
		if err != nil {
			b.Fatal(err)
		}
		_ = Day16(bits, false)
	}
}
