package adventofcode2021

import "testing"

func day06(t *testing.T, filename string, days uint, want uint) {
	lines, err := linesFromFilename(filename)
	if err != nil {
		t.Fatal(err)
	}
	fishes, err := ParseCommaSeparatedNumbers(lines[0])
	if err != nil {
		t.Fatal(err)
	}
	got, err := Day06(asUint(fishes), days)
	if err != nil {
		t.Fatal(err)
	}
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func asUint(is []int) []uint {
	us := make([]uint, len(is))
	for i := range is {
		us[i] = uint(is[i])
	}
	return us
}

func TestDay06Part1Example(t *testing.T) {
	// "In this example, after 18 days, there are a total of 26 fish."
	day06(t, exampleFilename(6), 18, 26)
	// " After 80 days, there would be a total of 5934."
	day06(t, exampleFilename(6), 80, 5934)
}

func TestDay06Part1(t *testing.T) {
	day06(t, filename(6), 80, 362639)
}

func TestDay06Part2Example(t *testing.T) {
	day06(t, exampleFilename(6), 256, 26984457539)
}

func TestDay06Part2(t *testing.T) {
	day06(t, filename(6), 256, 1639854996917)
}

func BenchmarkDay06Part1(b *testing.B) {
	lines, err := linesFromFilename(filename(6))
	if err != nil {
		b.Fatal(err)
	}
	fishes, err := ParseCommaSeparatedNumbers(lines[0])
	if err != nil {
		b.Fatal(err)
	}
	ufishes := asUint(fishes)
	b.ResetTimer()
	for range b.N {
		_, _ = Day06(ufishes, 80)
	}
}

func BenchmarkDay06Part2(b *testing.B) {
	lines, err := linesFromFilename(filename(6))
	if err != nil {
		b.Fatal(err)
	}
	fishes, err := ParseCommaSeparatedNumbers(lines[0])
	if err != nil {
		b.Fatal(err)
	}
	ufishes := asUint(fishes)
	b.ResetTimer()
	for range b.N {
		_, _ = Day06(ufishes, 256)
	}
}
