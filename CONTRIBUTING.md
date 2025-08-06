# Advent of Code Day NN (YEAR YYYY) Go Implementation Cookbook

The Advent Of Code is a yearly programming contest where participants solve puzzles in a competitive environment. 
It starts Dec 1st and ends Dec 25th.
For each day, participants are given a puzzle that they need to solve.
Puzzles are split into two parts, Part 1 and Part 2.

1. Puzzle description explains the problem and provides one or more examples.
2. It is your job to implement test cases for all examples. Use quotes in comments to back reference expected values.
3. Once all test cases for examples pass, create a test case for the actual puzzle input. Inputs are user specific, so don't come up with some random guesswork input data. If the puzzle input cannot be located, stop and ask.
4. Let the test case for part 1 run on the user specific input, and use the calculated value as the expected value for the test case.
5. Stop here and wait for the value to be verified by the AOC website.
6. If the value is correct, Part 2 is unlocked.
7. Repeat Step 2 - 5 for Part 2.

Do not implement code for puzzle parts that haven't been explicitly unlocked by the user.
Wait for the user to provide specific requirements and test cases for each part
Focus exclusively on the currently requested puzzle part rather than anticipating future requirements.
Do not implement part 1 when asked to implement the examples.
Do not implement part 2 when asked to implement part 1.

## Project Structre & Naming conventions

1. Single Package, Two-Digit Day IDs: All solution code resides in the root package (no
subpackages) and uses zero-padded day numbers in names 1 . For example, Day08 (not Day8)
would correspond to files day08.go and day08_test.go , containing NewDay08 (parser),
Day08 (solver), test functions TestDay08... , and benchmarks BenchmarkDay08...

This consistent naming makes it easy to locate code and ensures tests/benchmarks are
automatically discovered.

2. File Layout: The project typically provides input files in a testdata/ directory. Use the
provided helper functions (e.g. filename(day int) to get "testdata/dayNN.txt" and
possibly "testdata/dayNN_example.txt" ) and linesFromFilename (from input.go) to
read input lines conveniently. Keep all code for a given day in the single dayNN.go file (except
any example input files in testdata/ ). No additional utility packages or shared libraries should
be created – solve each day’s problem in isolation

3. Functions: Implement two primary functions per day:
- NewDayNN : Parsing function, responsible for reading or accepting raw input and converting it
into a suitable in-memory representation (data structure).
- DayNN : Solver function, which takes the parsed data from NewDayNN and a boolean part1
flag to distinguish Part 1 vs Part 2 logic 3 . By using a boolean parameter, you avoid duplicate
code and keep a single solve function that can handle both puzzle parts 3 . For example:
func Day08(data SomeType, part1 bool) uint .

## Input Parsing Efficiently

1. Use Custom Parsers: Advent of Code inputs are often plain text (lines of numbers, comma-
separated values, grids of characters, etc.). Write a custom parser in NewDayNN tailored to the
input format, rather than using heavy utilities like strings.Fields or regex, which can add
overhead. A purpose-built parser can read the input string/lines in one pass, converting data
on the fly and minimizing temporary allocations. For example, instead of splitting a comma-
separated line into a slice of substrings only to convert each to int, iterate over the line character
by character and build numbers directly (skipping delimiters). This approach avoids creating
throwaway slices or strings and is more GC-friendly.

2. Avoid Unnecessary Allocation: Parse in-place where possible. If you know the input size or can
estimate it, pre-allocate slices/arrays to that capacity. For instance, if the input has N lines or
comma-separated values, use make([]T, 0, N) to allocate once and append. A real example
from an AoC solution shows that replacing strings.Split with a simple loop to parse
characters significantly reduced allocations. Similarly, skip using fmt.Sscanf or regex for
parsing structured text; manual parsing with byte/ rune inspection is often faster and uses less
memory.

3. Minimal Data Structures: Choose the simplest data representation that fully encodes the
problem state. For example, if the input is a grid of known dimensions, consider using a 2D slice
or a flat 1D slice of length width*height instead of a map keyed by coordinates. If the puzzle
involves a fixed range of keys (e.g., 0–9 digits, or known small set of categories), prefer an array
or slice indexed by the key rather than a map. This avoids map overhead and takes advantage of
contiguous memory. In one performance analysis, changing a map[uint8]int (keys 0–8) to a
slice []int of length 9 eliminated all allocations and made the solution ~96% faster. The
contiguous slice is more cache-friendly and incurs no hashing cost, illustrating why a predictable
layout should use arrays/slices over maps.

4. Reuse Buffers and Slices: Where possible, reuse memory instead of allocating new slices or
objects in loops. For example, if you need a temporary slice inside a loop, allocate it once outside
and then reuse it by resetting its length (e.g. temp = temp[:0] ) each iteration. This pattern
avoids constant allocate/free cycles. In one AoC puzzle, reusing a slice for accumulating items
each cycle (by clearing it with [:0] ) drastically reduced allocations and GC pressure 6 . The
key is to identify memory that can be safely reset and reused because its content is short-lived.
Using a single bytes.Buffer for string building or a pooled object for repeated computations
can similarly help if the problem needs it.

5. Error Handling: While AoC inputs are well-formed (you can assume no malicious data), do
handle potential errors in parsing and file I/O. For example, check os.ReadFile / strconv.Atoi
errors and propagate or handle them (e.g., return an error from
NewDayNN or even panic in unreachable conditions). The code should prefer returning errors
over calling log.Fatal or os.Exit , especially in library code, to keep tests and usage clean.
During development you might log or panic for quick feedback, but remove all logging and
debugging output before final submission. The final code should be clean of print
statements and assume the input is correct (no need for excessive defensive checks beyond
basic error handling on file reads, etc.).

## Solution Implementation (Examples, Part 1, Part 2)
The puzzle description will contain a clear description of the problem and the expected output.
If the puzzle description contains more than one example, all examples MUST be used in test cases
to make sure the real input is solved correctly.

1. Single Solver Function: Implement func DayNN(data ParsedType, part1 bool) uint
to produce the answer for Part 1 or Part 2 depending on the flag. This function should contain
the core puzzle logic. Keeping part1 and part2 together (with an if part1 { ... } else
{ ... } internally) avoids duplicating parse or iterate logic, and ensures consistency. Often
Part 2 builds on Part 1; you might handle this by writing a common solving routine
parameterized by part, or by computing Part 1’s result and then extending it for Part 2. But do
not over-generalize – implement exactly what is needed for the given puzzle, no more 9 . For
example, if Part 2 requires an extra loop or a modification of Part 1’s process, handle that in the
conditional rather than trying to design an abstract framework.

2. Return Type – use uint : By project convention, any function that returns a count, sum, or
generally non-negative numeric result should use an unsigned integer type 10 . Using uint (or
a specific width like uint64 if extremely large numbers are possible) communicates that the
result cannot be negative and may prevent inadvertent negative values. For instance, if the
puzzle asks "how many pixels are lit," implement func DayNN(data ImageType, part1
bool) uint instead of returning an int . This also gives a slightly larger positive range than
int for free. Ensure you handle any arithmetic in unsigned form and be mindful of overflow if the
counts could be very large (in AoC, 64-bit is usually sufficient). If a puzzle does expect a negative
or signed result (rare in AoC), then obviously use int or the appropriate type; but for most
“count” style answers, stick to uint as required.

3. No Side Effects: DayNN should not modify global state or rely on external inputs. It should
ideally be a pure function (apart from perhaps some internal caching) that computes the result
from the given parsed data. This makes it easy to call from tests for various scenarios. If the
solution involves simulation or iterative transformation of the data, you can choose to mutate
the passed-in data structure for speed (document that in code if so), or work on a copy,
depending on performance needs. If mutating, ensure that parsing returns a fresh data
structure so that running Part 1 and Part 2 in sequence doesn't interfere unless intended. (In
some cases, you might reuse the computation from Part 1 in Part 2, but typically it’s cleaner to
recompute or have separate code paths toggled by the flag.)

4. Efficiency in Loops: Implement the core logic with attention to efficiency. Use simple loops and
conditionals – these compile down to very efficient code in Go. Avoid creating unnecessary
goroutines or channels for AoC tasks; the overhead usually outweighs any benefit since these
puzzles are mostly CPU-bound and not IO-bound. Likewise, avoid reflective or interface-heavy
constructs; stick to concrete types (slices, arrays, structs) for the hot paths. In critical loops, try to
minimize bounds checks and function calls inside the loop (the compiler often handles inlining
and bounds-check elimination if you write loops simply). For example, when iterating a slice,
using for i := 0; i < len(slice); i++ { … } can be a bit more efficient than
for _, v := range slice { … } in certain scenarios, and it gives you more control to
eliminate bounds checks if needed (though Go 1.21+ does this well in many cases). Use break /
continue when appropriate to skip unnecessary work.

## Memory & Performance Considerations

Performance is a high priority – solutions should aim to run in under a second even for the largest
inputs, and memory usage should be kept modest. Here are key considerations:

1. Minimize Allocations: Every allocation is work for the garbage collector. Strive to allocate once
and reuse. As mentioned, parse the input into efficient structures up front. During computation,
if you need temporary structures, consider initializing them once. For example, if you need a
buffer each iteration of a simulation, reuse the same buffer object or slice. A real-world AoC
optimization showed that reusing a slice (clearing it each time instead of making a new one)
reduced allocations by 95% and sped up the code significantly 6 11 . Tools like go test -
benchmem and pprof can help identify where allocations happen – often it’s in places like
append (when capacity runs out), string conversions, or map operations.

2. Prefer Arrays/Slices over Maps: Maps are powerful but come with overhead. If the problem’s
data can be indexed (by an integer or small range), an array or slice will be much faster. Iterating
an array is more cache-friendly than iterating a map with random order. For instance, using a
fixed-size array of length 8 or 9 to count items (instead of a map with keys 0-8) can remove all
overhead of map hashing and allocation 5 . Only use maps when the input size or keys are truly
dynamic or unbounded. Even then, consider the access patterns: sometimes sorting a slice and
binary searching can outperform a map lookup for moderately sized data sets.

3. In-Place Operations: Whenever you can update data in-place without breaking correctness, do
so. This might mean modifying a slice instead of creating a new one for each step. For example,
if you simulate a process over an array, try to update the array in place or toggle between two
pre-allocated arrays (double-buffering) rather than allocating a new array each time. If you must
create new slices (say, filtering a list), consider reusing the old slice's memory (set length to 0 and
append) or using copy . The standard library’s slices package (Go 1.21+) even provides
slices.Clip to drop a slice’s excess capacity when needed 12 , which can help free memory
if you significantly overshot capacity, but typically you can size things correctly from the start in
AoC tasks.

4. Algorithmic Efficiency: Choose algorithms with appropriate complexity. Advent of Code often
doesn’t require highly complex algorithms if input sizes are moderate, but occasionally an
O(n^2) approach will be too slow on a large input. Aim for linear or linearithmic solutions
when possible. Use Go’s efficient sort ( sort.Slice or slices.Sort in Go 1.21) for sorting
needs, which is highly optimized. For graph or search problems, prefer iterative approaches or
use a queue (slice) rather than recursion to avoid stack overhead. Before micro-optimizing, make
sure the overall approach is sound – no amount of micro-optimization will save an algorithm
that’s exponentially slow on large inputs.

5. Leverage Compiler Optimizations: Modern Go can inline functions, eliminate bounds checks,
and vectorize certain loops. Write code in a way that the compiler can understand and optimize.
For example, using built-in functions from standard packages can be beneficial: functions like
copy() for slices or bytes.Equal are highly optimized in Go’s runtime (often using
assembly under the hood). Using the new slices and maps generics can be both convenient
and fast – for example, summing a slice of numbers using slices.Sum (if available in your Go
version) is concise and is expected to be as efficient as a manual loop (the compiler will inline
simple generic functions). Always verify with benchmarks if in doubt, but generally, these library
routines are well-implemented.

## Utilizing Latest Go Features

1. Target Go Version: Always check the Go version in go.mod and use the full feature set of that
Go release. If the project specifies Go 1.21 or newer, you have access to the new standard
packages like slices and maps , as well as built-in functions like min , max (added in
Go 1.21) for convenience. Using these can make code shorter and clearer without sacrificing
performance. For example, the slices package provides generic slice manipulation: sorting
( slices.Sort ), searching ( slices.Index ), and possibly summing or containing (depending
on version). These functions are type-safe and often compile down to simple loops.

2. Generics and Standard Library: Go’s standard library now includes packages (like slices and
maps ) in the core or golang.org/x/exp (for experimental features) that can be used to
avoid writing boilerplate. If go.mod indicates use of an experimental package (e.g.
golang.org/x/exp/slicesfor older Go versions), it’s permitted since it’s an official
extension. But do not pull in random third-party libraries – stick to what’s provided by the Go
team. For instance, do not use a third-party collections library or parsing library; instead, use
what's in the stdlib or write it yourself. The goal is to rely on well-optimized, trusted code or
custom code you control, and avoid any external dependency burden.

3. New Language Features: If the Go version allows, use features like range over strings yielding
Unicode code points (if needed for character processing), use multiple result assignment cleverly
to swap values (for example, rotating buffers), and use defer effectively for clean-up if it doesn’t
impact performance critical loops. However, avoid using fancy features in hot loops if they add
overhead – e.g., reflection or interface{} should be avoided in tight computations. Stick to
features that make code faster or clearer without cost.

4. Example – Slices Package Usage: Suppose you have a slice of integers and you need the sum.
Instead of manually looping, if using Go 1.21+, you might write total :=
slices.Sum(intSlice) (if available) or use maps.Values and then slices.Sort on a
map’s values if you need sorted values, etc. These abstractions are fine because they compile to
efficient code. Writing clear code using such helpers is encouraged when performance is not
compromised. Always double-check if a particular helper (like slices.Delete or
slices.Insert ) is efficient enough or if a custom approach would be better; in many cases
the library is optimal, but for something trivial like summing, both approaches are O(n) anyway.

## Testing Strategy

Writing tests is crucial to ensure your solution works for both the example scenarios and your specific
puzzle input:

1. Example Tests: Start by writing tests for the examples given in the puzzle description. Typically,
each AoC problem description includes one or more small examples with expected outputs for
Part 1 (and similarly for Part 2 once unlocked). Create a test function named
TestDayNNPart1Example in dayNN_test.go that constructs the example input and verifies
the output. You can either embed the example input directly as a literal (multi-line string or slice
of strings) or place it in testdata/dayNN_example.txt and load it via
linesFromFilename . For clarity and consistency, many prefer using testdata files for
examples as well, in which case you might have dayNN_example.txt containing the exact
example input. Use the parsing and solving functions in the test: e.g.

func TestDay08Part1Example(t *testing.T) {
    lines := linesFromFilename("testdata/day08_example.txt")
    data := NewDay08(lines)
    got := Day08(data, true) // part1 = true
    want := uint(26)
    // expected output for the example
    if got != want {
        t.Errorf("Day08 Part1 example: got %d, want %d", got, want)
    }
}

If there are multiple example cases, you can either write separate test functions for each (e.g.,
TestDay08Part1Example2 ), or loop over a list of sub-cases in one test. The key is to cover all
provided examples 14 . Do the same for Part 2 with TestDayNNPart2Example once you implement
Part 2.

1. Puzzle Input Tests: Once the example tests pass, write tests for the actual puzzle input.
According to the guidelines, you should have TestDayNNPart1 and TestDayNNPart2
functions that read the full puzzle input (from the testdata/dayNN.txt file) and verify the
correct answers for Part 1 and Part 2, respectively 15 16 . These expected answers would be the
ones you obtained when running your solution on the real input (usually by submitting to AoC,
you get the correct answer). It's important to include these tests so that any future changes or
refactoring can be validated against known correct results. For these tests, use the same pattern:
load input lines, call NewDayNN , then DayNN(..., true) or false , and compare to the
known result. For example:

func TestDay08Part1(t *testing.T) {
    lines := linesFromFilename(filename(8)) // using helper to get "testdata/day08.txt"
    data := NewDay08(lines)
    got := Day08(data, true)
    want := uint(1673) // (for instance, the known answer for part1)
    if got != want {
        t.Errorf("Day08 Part1: got %d, want %d", got, want)
    }
}

And similarly TestDay08Part2 for Part 2. The project helpers like filename(day) and
exampleFilename(day) give the
proper path. Remember to use uint(...) for literal expected values if they are not already typed
as uint, to match the function’s return type.

2. Test Behavior: Tests should not print anything on success; rely on the testing framework output.
Use t.Fatalf or t.Errorf to report mismatches. Also, consider edge cases if any (though
AoC problems typically ensure inputs meet certain conditions, you might still test minimal inputs
or boundary behavior if applicable). Make sure to run go test ./... to confirm all tests
pass.

3. Example Continuity: Some AoC part 2 descriptions build on the part 1 example (the same input
with a new twist). Ensure your Part 2 example test uses the appropriate expected result given the
example input. It might reuse the same dayNN_example.txt if the example input is the same
for both parts, just checking a different expected value with part1=false .

4. No Flaky Tests: Because performance is a concern, you might want to avoid tests that take too
long. But AoC puzzles are usually fine to run in tests (they are not that large to exceed typical test
timeouts). If your solution for the actual input is a bit slow, it's still good to test it; you can tag
very slow tests with a build tag or use t.Skip() based on an environment variable, but
usually this is not needed if your solution is optimized as intended.

## Benchmarking

Benchmarking your solution ensures it meets the performance goals and helps catch regressions or
high allocation rates:

1. Benchmark Functions: Include BenchmarkDayNNPart1 and BenchmarkDayNNPart2 in
dayNN_test.go. These should measure the end-to-end execution time of parsing +
solving for each part. The convention is to include parsing in the benchmark, since input parsing
can be a significant part of AoC runtime and we want to optimize it as well. For example:

func BenchmarkDay08Part1(b *testing.B) {
    raw := linesFromFilename(filename(8)) // load the input once
    b.ResetTimer()
    for range b.N {
        data := NewDay08(raw)
        _ = Day08(data, true)
    }
}

This will run the full parse and solve b.N times. You might also call b.ReportAllocs() to have Go
count allocations, helping ensure you meet the “low GC” objective. If parsing is expensive and you want
to measure solve alone, you could separate those, but the project explicitly asks to include parsing in
benchmarks 18 , so the above approach is correct. For Part 2, similarly call DayNN(data, false)
inside the loop.

2. Interpreting Benchmark Results: After running go test -bench . -benchmem , you should
look at ns/op (nanoseconds per operation), B/op (bytes allocated per op), and allocs/op. Aim for
as low B/op and allocs/op as possible, ideally in the tens of kilobytes or less and a small number
of allocations. If you see a high number of allocations, use tools to identify where (for example,
pprof or just searching for make / append in your code). The example of Day 6 optimization
reduced allocations from 1025 per op to essentially 0 by using better data structures and reuse.
While not every puzzle can be zero-allocation, you should strive to get rid of any obvious
unnecessary allocations.

3. Benchmarking Iterations: Ensure your benchmark doesn’t carry state between iterations. In
the example above, we parse fresh each time inside the loop. If instead you wanted to factor
parsing out (say, to measure pure solve time), be careful to clone any data structures if the solve
mutates them. In our case, including parsing naturally avoids any cross-iteration issues since we
create a new data each time.

## Final Checks and Best Practices

1. Code Cleanliness: Before considering the solution final, remove all debug prints, log
statements, or stray comments that were used during development. The code should be
production-quality: clear, formatted ( go fmt ), with no lint or vet issues. The project demands
that the code “vet clean” – run go vet ./... to ensure there are no suspicious constructs.

2. No Panics or Exits: The solution should not call panic() or os.Exit(1) on normal
operation. Panics should be reserved for truly unrecoverable situations (which ideally you won't
encounter with valid AoC input). Instead, handle errors gracefully by returning an error from
NewDayNN or by failing tests via t.Fatal .

3. Adhere to Spec (No hardcoding Answers): It should go without saying, but never hardcode the
known answer for your specific input. Always implement the logic to derive it. AoC inputs differ
per user, and even though in tests we use our input’s known answer as the expected value, the
solution must compute it. Also, do not include any code or solutions that aren't your own work
(no copying from others’ AoC solutions). The goal is a clean-room implementation following the
puzzle description.

4. Focus on Current Puzzle Only: Don’t generalize your solution in anticipation of other days or
reuse. Each day’s code should solve that day’s puzzle and nothing more 9 . For example, if
multiple days have similar grid processing, you might be tempted to abstract a helper – resist
that for this project. Keep the code self-contained for Day NN. This also means if Part 2 can be
handled with a quick tweak to Part 1’s logic, just implement that tweak inline (controlled by the
part1 flag) rather than building an elaborate framework.

5. No Global State: avoid global variables or package-level state to hold input or results; this makes testing
easier and prevents cross-test interference. Each test or run should explicitly parse and then
solve using these functions.


## Dont's
Here are some additional rules to avoid when implementing a solution:

- looking up solutions online
- Using a structure that is used to pass parsing results to the solver that has less than three fields. Use fields directly instead.
- Making the solver function a method. Use a function instead (https://go.dev/doc/effective_go#methods).
- Use recursion. Never ever use recursion. There's two kind of people in this world: those who use 
recursion, and those that have been woken up by their pager at 3am because of a stack overflow.
- guess some puzzle input because the corresponding file in `testdata/day{{DAY}}.txt` doesn't exist. Instead, stop and wait for the input to be available. 