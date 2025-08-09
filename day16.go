package adventofcode2021

import (
	"strconv"
)

// Packet represents a BITS packet
// The puzzle description indicates packets have a version, type ID, and content
// that varies based on the type ID.
type Packet struct {
	Version uint
	TypeID  uint
	// For literal packets
	Value uint
	// For operator packets
	SubPackets []Packet
}

// NewDay16 parses the input data for Day 16 and returns a Packet.
// The input is a single line containing a hexadecimal string.
func NewDay16(input string) (Packet, error) {
	// Convert hex string to binary
	binStr := hexToBin(input)

	// Parse the packet
	packet, _ := parsePacket(binStr)
	return packet, nil
}

// Day16 solves the Day 16 puzzle.
// part1 flag indicates whether to solve part 1 (true) or part 2 (false).
func Day16(hexStr string, part1 bool) uint {
	packet, _ := NewDay16(hexStr)
	if part1 {
		return sumVersions(packet)
	}
	return evaluatePacket(packet)
}

// hexToBin converts a hexadecimal string to a binary string
var hexToBinMap = [256]string{
	'0': "0000", '1': "0001", '2': "0010", '3': "0011",
	'4': "0100", '5': "0101", '6': "0110", '7': "0111",
	'8': "1000", '9': "1001", 'A': "1010", 'B': "1011",
	'C': "1100", 'D': "1101", 'E': "1110", 'F': "1111",
	'a': "1010", 'b': "1011", 'c': "1100", 'd': "1101",
	'e': "1110", 'f': "1111",
}

func hexToBin(hexStr string) string {
	size := len(hexStr) * 4
	binStr := make([]byte, 0, size)
	for i := 0; i < len(hexStr); i++ {
		binStr = append(binStr, hexToBinMap[hexStr[i]]...)
	}
	return string(binStr)
}

// parsePacket parses a binary string into a Packet
// This is an internal helper function used by NewDay16
// Returns the packet and the number of bits consumed
func parsePacket(binStr string) (Packet, int) {
	if len(binStr) < 6 { // Minimum packet size is 6 bits (version + type)
		return Packet{}, 0
	}

	// Parse version (first 3 bits) and type (next 3 bits) using bit operations
	version := uint(binStr[0]-'0')<<2 | uint(binStr[1]-'0')<<1 | uint(binStr[2]-'0')
	typeID := uint(binStr[3]-'0')<<2 | uint(binStr[4]-'0')<<1 | uint(binStr[5]-'0')

	var p Packet
	p.Version = version
	p.TypeID = typeID

	// Literal value packet
	if typeID == 4 {
		value, bitsRead := parseLiteralValue(binStr[6:])
		p.Value = value
		return p, 6 + bitsRead
	}

	// Operator packet
	lengthTypeID := binStr[6] - '0'
	var subPackets []Packet
	var bitsRead int

	if lengthTypeID == 0 {
		// Next 15 bits are total length in bits of sub-packets
		length := binToUint(binStr[7:22])
		subPacketsStr := binStr[22 : 22+int(length)]
		subPackets, _ = parseSubPacketsByLength(subPacketsStr)
		bitsRead = 22 + int(length)
	} else {
		// Next 11 bits are number of sub-packets
		subPacketCount := binToUint(binStr[7:18])
		subPackets, bitsRead = parseSubPacketsByCount(binStr[18:], int(subPacketCount))
		bitsRead += 18
	}

	p.SubPackets = subPackets
	return p, bitsRead
}

// parseLiteralValue parses a literal value from a binary string
// Returns the value and the number of bits consumed
func parseLiteralValue(binStr string) (uint, int) {
	var value uint
	var bitsRead int

	for i := 0; i < len(binStr); i += 5 {
		if i+5 > len(binStr) {
			break
		}

		// Get the current 5-bit group
		group := binStr[i : i+5]

		// Parse the 4 value bits and add them to the result
		for j := 1; j < 5; j++ {
			value = (value << 1) | uint(group[j]-'0')
		}

		bitsRead += 5

		// Last group has 0 as the first bit
		if group[0] == '0' {
			break
		}
	}

	return value, bitsRead
}

// parseSubPacketsByLength parses sub-packets from a binary string with a given total bit length
func parseSubPacketsByLength(binStr string) ([]Packet, int) {
	var packets []Packet
	totalBitsRead := 0
	for totalBitsRead < len(binStr) {
		packet, bitsRead := parsePacket(binStr[totalBitsRead:])
		if bitsRead == 0 {
			break
		}
		packets = append(packets, packet)
		totalBitsRead += bitsRead
	}
	return packets, totalBitsRead
}

// parseSubPacketsByCount parses a specific number of sub-packets from a binary string
func parseSubPacketsByCount(binStr string, count int) ([]Packet, int) {
	var packets []Packet
	totalBitsRead := 0
	for i := 0; i < count && totalBitsRead < len(binStr); i++ {
		packet, bitsRead := parsePacket(binStr[totalBitsRead:])
		if bitsRead == 0 {
			break
		}
		packets = append(packets, packet)
		totalBitsRead += bitsRead
	}
	return packets, totalBitsRead
}

// binToUint converts a binary string to uint
func binToUint(binStr string) uint {
	val, _ := strconv.ParseUint(binStr, 2, 64)
	return uint(val)
}

// sumVersions sums the version numbers of all packets using iterative DFS
func sumVersions(root Packet) uint {
	sum := uint(0)
	stack := []Packet{root}

	for len(stack) > 0 {
		// Pop the last element
		n := len(stack) - 1
		p := stack[n]
		stack = stack[:n]

		// Add this packet's version
		sum += p.Version

		// Push all sub-packets onto the stack
		for i := len(p.SubPackets) - 1; i >= 0; i-- {
			stack = append(stack, p.SubPackets[i])
		}
	}

	return sum
}

// evaluatePacket evaluates the packet expression based on its type ID
func evaluatePacket(p Packet) uint {
	switch p.TypeID {
	case 0: // sum
		sum := uint(0)
		for i := range p.SubPackets {
			sum += evaluatePacket(p.SubPackets[i])
		}
		return sum

	case 1: // product
		if len(p.SubPackets) == 0 {
			return 0
		}
		product := uint(1)
		for i := range p.SubPackets {
			product *= evaluatePacket(p.SubPackets[i])
		}
		return product

	case 2: // minimum
		if len(p.SubPackets) == 0 {
			return 0
		}
		min := evaluatePacket(p.SubPackets[0])
		for _, subPkt := range p.SubPackets[1:] {
			if val := evaluatePacket(subPkt); val < min {
				min = val
			}
		}
		return min

	case 3: // maximum
		if len(p.SubPackets) == 0 {
			return 0
		}
		max := evaluatePacket(p.SubPackets[0])
		for _, subPkt := range p.SubPackets[1:] {
			if val := evaluatePacket(subPkt); val > max {
				max = val
			}
		}
		return max

	case 4: // literal
		return p.Value

	case 5: // greater than
		if len(p.SubPackets) != 2 {
			return 0
		}
		a := evaluatePacket(p.SubPackets[0])
		b := evaluatePacket(p.SubPackets[1])
		if a > b {
			return 1
		}
		return 0

	case 6: // less than
		if len(p.SubPackets) != 2 {
			return 0
		}
		a := evaluatePacket(p.SubPackets[0])
		b := evaluatePacket(p.SubPackets[1])
		if a < b {
			return 1
		}
		return 0

	case 7: // equal
		if len(p.SubPackets) != 2 {
			return 0
		}
		a := evaluatePacket(p.SubPackets[0])
		b := evaluatePacket(p.SubPackets[1])
		if a == b {
			return 1
		}
		return 0

	default:
		return 0
	}
}
