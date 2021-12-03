package adventofcode2021

import "testing"

func TestDay01Part1Example(t *testing.T) {
	const want = 7
	ns, err := numbersFromFilename(exampleFilename(1))
	if err != nil {
		t.Fatal(err)
	}
	got := Day01Part1(ns)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay01(t *testing.T) {
	const want = 1266
	ns, err := numbersFromFilename(filename(1))
	if err != nil {
		t.Fatal(err)
	}
	got := Day01Part1(ns)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay01Part2Example(t *testing.T) {
	const want = 5
	ns, err := numbersFromFilename(exampleFilename(1))
	if err != nil {
		t.Fatal(err)
	}
	got := Day01Part2(ns)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay01Part2(t *testing.T) {
	const want = 1217
	ns, err := numbersFromFilename(filename(1))
	if err != nil {
		t.Fatal(err)
	}
	got := Day01Part2(ns)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func BenchmarkDay01Part2(b *testing.B) {
	ns, err := numbersFromFilename(filename(1))
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Day01Part2(ns)
	}
}
