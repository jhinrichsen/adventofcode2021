package adventofcode2021

import (
	"testing"
)

func TestDay16Part1Example1(t *testing.T) {
	// "With the first example, the hex string 8A004A801A8002F478 represents an operator packet (version 4) which contains an operator packet (version 1) which contains an operator packet (version 5) which contains a literal value (version 6); this packet has a version sum of 16."
	const want = 16
	got := Day16("8A004A801A8002F478", true)
	if got != want {
		t.Fatalf("got %d, want %d", got, want)
	}
}

func TestDay16Part1Example2(t *testing.T) {
	// "The second example, 620080001611562C8802118E34, results in a version sum of 12."
	const want = 12
	got := Day16("620080001611562C8802118E34", true)
	if got != want {
		t.Fatalf("got %d, want %d", got, want)
	}
}

func TestDay16Part1Example3(t *testing.T) {
	// "The third example, C0015000016115A2E0802F182340, results in a version sum of 23."
	const want = 23
	got := Day16("C0015000016115A2E0802F182340", true)
	if got != want {
		t.Fatalf("got %d, want %d", got, want)
	}
}

func TestDay16Part1Example4(t *testing.T) {
	// "The fourth example, A0016C880162017C3686B18A3D4780, results in a version sum of 31."
	const want = 31
	got := Day16("A0016C880162017C3686B18A3D4780", true)
	if got != want {
		t.Fatalf("got %d, want %d", got, want)
	}
}

func TestDay16Part1(t *testing.T) {
	const want = 891
	lines, err := linesFromFilename(filename(16))
	if err != nil {
		t.Fatal(err)
	}
	hexStr := lines[0]
	got := Day16(hexStr, true)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func BenchmarkDay16(b *testing.B) {
	lines, err := linesFromFilename(filename(16))
	if err != nil {
		b.Fatal(err)
	}
	for range b.N {
		Day16(lines[0], true)
	}
}
