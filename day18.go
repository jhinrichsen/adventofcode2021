package adventofcode2021

import (
	"fmt"
)

// SnailfishNumber represents a snailfish number - either a regular number or a pair
type SnailfishNumber struct {
	value      int
	left       *SnailfishNumber
	right      *SnailfishNumber
	isRegular  bool
}

// parseSnailfish parses a snailfish number from a string
func parseSnailfish(s string) (*SnailfishNumber, int) {
	if len(s) == 0 {
		return nil, 0
	}

	// Regular number
	if s[0] >= '0' && s[0] <= '9' {
		num := 0
		i := 0
		for i < len(s) && s[i] >= '0' && s[i] <= '9' {
			num = num*10 + int(s[i]-'0')
			i++
		}
		return &SnailfishNumber{value: num, isRegular: true}, i
	}

	// Pair: [left,right]
	if s[0] == '[' {
		left, leftLen := parseSnailfish(s[1:])
		// Skip comma
		comma := 1 + leftLen
		right, rightLen := parseSnailfish(s[comma+1:])
		// Skip closing bracket
		total := comma + 1 + rightLen + 1
		return &SnailfishNumber{left: left, right: right, isRegular: false}, total
	}

	return nil, 0
}

// clone creates a deep copy of a snailfish number
func (n *SnailfishNumber) clone() *SnailfishNumber {
	if n == nil {
		return nil
	}
	if n.isRegular {
		return &SnailfishNumber{value: n.value, isRegular: true}
	}
	return &SnailfishNumber{
		left:      n.left.clone(),
		right:     n.right.clone(),
		isRegular: false,
	}
}

// add creates a new snailfish number by adding two numbers
func add(a, b *SnailfishNumber) *SnailfishNumber {
	result := &SnailfishNumber{
		left:      a.clone(),
		right:     b.clone(),
		isRegular: false,
	}
	reduce(result)
	return result
}

// reduce applies reduction rules until no more reductions are possible
func reduce(n *SnailfishNumber) {
	for {
		// Try to explode
		if tryExplode(n) {
			continue
		}
		// Try to split
		if splitAny(n) {
			continue
		}
		// No more reductions possible
		break
	}
}

// tryExplode attempts to explode the leftmost pair nested 4 levels deep
func tryExplode(root *SnailfishNumber) bool {
	// Collect all regular numbers in left-to-right order
	var nums []*SnailfishNumber
	var collectNums func(*SnailfishNumber)
	collectNums = func(node *SnailfishNumber) {
		if node == nil {
			return
		}
		if node.isRegular {
			nums = append(nums, node)
			return
		}
		collectNums(node.left)
		collectNums(node.right)
	}
	collectNums(root)

	// Find the leftmost pair at depth 4 or greater
	var explodePair *SnailfishNumber
	var explodeParent *SnailfishNumber
	var explodeIsLeft bool

	var findPair func(*SnailfishNumber, *SnailfishNumber, bool, int) bool
	findPair = func(node, parent *SnailfishNumber, isLeft bool, depth int) bool {
		if node == nil || node.isRegular {
			return false
		}

		// Check if this pair should explode
		if depth >= 4 && node.left != nil && node.left.isRegular && node.right != nil && node.right.isRegular {
			explodePair = node
			explodeParent = parent
			explodeIsLeft = isLeft
			return true
		}

		// Recurse left first (for leftmost)
		if findPair(node.left, node, true, depth+1) {
			return true
		}
		if findPair(node.right, node, false, depth+1) {
			return true
		}
		return false
	}

	if !findPair(root, nil, false, 0) {
		return false
	}

	// Found a pair to explode
	leftVal := explodePair.left.value
	rightVal := explodePair.right.value

	// Find the exploding pair's left and right values in the nums array
	leftIdx := -1
	rightIdx := -1
	for i, num := range nums {
		if num == explodePair.left {
			leftIdx = i
		}
		if num == explodePair.right {
			rightIdx = i
		}
	}

	// Add left value to the first regular number to the left
	if leftIdx > 0 {
		nums[leftIdx-1].value += leftVal
	}

	// Add right value to the first regular number to the right
	if rightIdx >= 0 && rightIdx < len(nums)-1 {
		nums[rightIdx+1].value += rightVal
	}

	// Replace the exploding pair with 0
	replacement := &SnailfishNumber{value: 0, isRegular: true}
	if explodeParent == nil {
		// This shouldn't happen in practice for valid inputs
		return false
	}
	if explodeIsLeft {
		explodeParent.left = replacement
	} else {
		explodeParent.right = replacement
	}

	return true
}

// splitAny tries to split the leftmost regular number >= 10
func splitAny(n *SnailfishNumber) bool {
	if n == nil {
		return false
	}

	if n.isRegular {
		if n.value >= 10 {
			// Can't split in place, need parent
			return false
		}
		return false
	}

	// Check if left child is a regular number >= 10
	if n.left != nil && n.left.isRegular && n.left.value >= 10 {
		val := n.left.value
		n.left = &SnailfishNumber{
			left:      &SnailfishNumber{value: val / 2, isRegular: true},
			right:     &SnailfishNumber{value: (val + 1) / 2, isRegular: true},
			isRegular: false,
		}
		return true
	}

	// Recurse into left subtree
	if !n.left.isRegular {
		if splitAny(n.left) {
			return true
		}
	}

	// Check if right child is a regular number >= 10
	if n.right != nil && n.right.isRegular && n.right.value >= 10 {
		val := n.right.value
		n.right = &SnailfishNumber{
			left:      &SnailfishNumber{value: val / 2, isRegular: true},
			right:     &SnailfishNumber{value: (val + 1) / 2, isRegular: true},
			isRegular: false,
		}
		return true
	}

	// Recurse into right subtree
	if !n.right.isRegular {
		if splitAny(n.right) {
			return true
		}
	}

	return false
}

// magnitude calculates the magnitude of a snailfish number
func magnitude(n *SnailfishNumber) uint {
	if n == nil {
		return 0
	}
	if n.isRegular {
		return uint(n.value)
	}
	return 3*magnitude(n.left) + 2*magnitude(n.right)
}

// Day18 solves day 18 puzzle
func Day18(lines []string, part1 bool) uint {
	if len(lines) == 0 {
		return 0
	}

	// Parse all numbers
	numbers := make([]*SnailfishNumber, 0, len(lines))
	for _, line := range lines {
		num, _ := parseSnailfish(line)
		numbers = append(numbers, num)
	}

	if part1 {
		// Add all numbers together
		result := numbers[0].clone()
		for i := 1; i < len(numbers); i++ {
			result = add(result, numbers[i])
		}
		return magnitude(result)
	}

	// Part 2: Find largest magnitude from adding any two different numbers
	maxMag := uint(0)
	for i := range numbers {
		for j := range numbers {
			if i == j {
				continue
			}
			result := add(numbers[i], numbers[j])
			mag := magnitude(result)
			if mag > maxMag {
				maxMag = mag
			}
		}
	}
	return maxMag
}

// String representation for debugging
func (n *SnailfishNumber) String() string {
	if n == nil {
		return ""
	}
	if n.isRegular {
		return fmt.Sprintf("%d", n.value)
	}
	return fmt.Sprintf("[%s,%s]", n.left.String(), n.right.String())
}
