package adventofcode2021

import "fmt"

// BITS represents the binary data for the BITS transmission
type BITS []byte

// readBits reads n bits from the BITS data starting at offset
// Returns the value and new offset
func readBits(data BITS, offset, n int) (uint, int) {
	var result uint
	for i := 0; i < n; i++ {
		if offset >= len(data)*8 {
			break
		}
		bytePos := offset / 8
		bitPos := 7 - (offset % 8) // MSB first
		bit := (data[bytePos] >> bitPos) & 1
		result = (result << 1) | uint(bit)
		offset++
	}
	return result, offset
}

// NewDay16 parses the input data for Day 16 and returns the BITS representation.
// The input is a hexadecimal string where each character represents 4 bits.
// Returns an error if any character in the input is not a valid hex digit (0-9, A-F).
func NewDay16(s string) (BITS, error) {
	// Allocate enough bytes to hold all nibbles (2 per byte)
	bits := make(BITS, (len(s)+1)/2)

	for i := 0; i < len(s); i++ {
		c := s[i]
		var val byte

		switch {
		case '0' <= c && c <= '9':
			val = c - '0'
		case 'A' <= c && c <= 'F':
			val = c - 'A' + 10
		default:
			return nil, fmt.Errorf("invalid hex digit: %c (must be 0-9 or A-F)", c)
		}

		// Pack two nibbles into each byte
		if i%2 == 0 {
			bits[i/2] = val << 4 // High nibble
		} else {
			bits[i/2] |= val // Low nibble
		}
	}

	return bits, nil
}

// parseLiteral parses a literal value from the BITS data starting at offset
// Returns the value and new offset
func parseLiteral(data BITS, offset int) (uint, int) {
	var value uint

	for {
		if offset+5 > len(data)*8 {
			break
		}

		// Read 5-bit group
		group, newOffset := readBits(data, offset, 5)
		offset = newOffset

		// Add the 4 value bits to the result
		value = (value << 4) | (group & 0x0F)

		// Check if this is the last group
		if (group & 0x10) == 0 {
			break
		}
	}

	return value, offset
}

// stackFrame represents the state of a packet being processed
type stackFrame struct {
	offset     int    // current bit offset in the BITS data
	count      int    // remaining sub-packets to process (when isCount is true)
	isCount    bool   // true if using count, false if using endOffset
	endOffset  int    // used when isCount is false
	versionSum uint   // for part 1
	values     []uint // for part 2
	typeID     uint   // operator type ID
	hasValue   bool   // true if we've processed this frame's value
}

// sumVersions iteratively sums all version numbers in the BITS transmission
func sumVersions(bits BITS, startOffset int) (uint, int) {
	var total uint
	offset := startOffset

	for {
		// Check for end of input
		if offset+6 > len(bits)*8 {
			break
		}

		// Read version and type ID
		version, newOffset := readBits(bits, offset, 3)
		typeID, newOffset := readBits(bits, newOffset, 3)

		// Add version to total for part 1
		total += version
		offset = newOffset

		// Handle literal value
		if typeID == 4 {
			// Skip over the literal value
			for {
				done, newOffset := readBits(bits, offset, 1)
				offset = newOffset
				_, offset = readBits(bits, offset, 4)
				if done == 0 {
					break
				}
			}
		} else {
			// Handle operator - just skip over the operator payload
			lengthTypeID, newOffset := readBits(bits, offset, 1)
			offset = newOffset

			if lengthTypeID == 0 {
				// Next 15 bits are total length in bits
				_, offset = readBits(bits, offset, 15)
			} else {
				// Next 11 bits are number of sub-packets
				_, offset = readBits(bits, offset, 11)
			}
		}

		// If we've reached the end of the input, break
		if offset >= len(bits)*8 {
			break
		}
	}

	return total, offset
}

// evaluatePacket evaluates a packet and returns its value and the new offset
func evaluatePacket(bits BITS, offset int) (uint, int) {
	// Read version and type ID
	_, offset = readBits(bits, offset, 3) // Skip version for part 2
	typeID, offset := readBits(bits, offset, 3)

	// Handle literal value
	if typeID == 4 {
		return parseLiteral(bits, offset)
	}

	// Handle operator
	lengthTypeID, offset := readBits(bits, offset, 1)
	var values []uint

	if lengthTypeID == 0 {
		// Next 15 bits are total length in bits
		length, newOffset := readBits(bits, offset, 15)
		offset = newOffset
		endOffset := offset + int(length)

		// Parse sub-packets until we reach the end offset
		for offset < endOffset {
			value, newOffset := evaluatePacket(bits, offset)
			values = append(values, value)
			offset = newOffset
		}
	} else {
		// Next 11 bits are number of sub-packets
		numSubPackets, newOffset := readBits(bits, offset, 11)
		offset = newOffset

		// Parse exactly numSubPackets sub-packets
		for i := 0; i < int(numSubPackets); i++ {
			value, newOffset := evaluatePacket(bits, offset)
			values = append(values, value)
			offset = newOffset
		}
	}

	// Apply operation based on typeID
	var result uint
	switch typeID {
	case 0: // sum
		for _, v := range values {
			result += v
		}
	case 1: // product
		result = 1
		for _, v := range values {
			result *= v
		}
	case 2: // minimum
		if len(values) > 0 {
			result = values[0]
			for _, v := range values[1:] {
				if v < result {
					result = v
				}
			}
		}
	case 3: // maximum
		if len(values) > 0 {
			result = values[0]
			for _, v := range values[1:] {
				if v > result {
					result = v
				}
			}
		}
	case 5: // greater than
		if len(values) >= 2 && values[0] > values[1] {
			result = 1
		}
	case 6: // less than
		if len(values) >= 2 && values[0] < values[1] {
			result = 1
		}
	case 7: // equal to
		if len(values) >= 2 && values[0] == values[1] {
			result = 1
		}
	}

	return result, offset
}

// Day16 solves the Day 16 puzzle.
// It takes the parsed BITS input (from NewDay16) and a part1 flag to determine which part to solve.
// part1=true returns the sum of all version numbers (part 1).
// part1=false returns the evaluated expression value (part 2).
func Day16(bits BITS, part1 bool) uint {
	if len(bits) == 0 {
		return 0
	}

	if part1 {
		total, _ := sumVersions(bits, 0)
		return total
	}

	// For part 2, evaluate the expression
	result, _ := evaluatePacket(bits, 0)
	return result
}
