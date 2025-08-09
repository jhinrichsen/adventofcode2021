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
func hexToBin(hexStr string) string {
	hexToBinMap := map[rune]string{
		'0': "0000",
		'1': "0001",
		'2': "0010",
		'3': "0011",
		'4': "0100",
		'5': "0101",
		'6': "0110",
		'7': "0111",
		'8': "1000",
		'9': "1001",
		'A': "1010",
		'B': "1011",
		'C': "1100",
		'D': "1101",
		'E': "1110",
		'F': "1111",
	}

	binStr := ""
	for _, h := range hexStr {
		binStr += hexToBinMap[h]
	}
	return binStr
}

// parsePacket parses a binary string into a Packet
// This is an internal helper function used by NewDay16
// Returns the packet and the number of bits consumed
func parsePacket(binStr string) (Packet, int) {
	if len(binStr) < 6 { // Minimum packet size is 6 bits (version + type)
		return Packet{}, 0
	}

	// Parse version (first 3 bits)
	version := binToUint(binStr[0:3])
	typeID := binToUint(binStr[3:6])

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
	var valueStr string
	var i int
	for i = 0; i < len(binStr); i += 5 {
		if i+5 > len(binStr) {
			break
		}
		group := binStr[i : i+5]
		valueStr += group[1:]
		if group[0] == '0' {
			i += 5
			break
		}
	}
	return binToUint(valueStr), i
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

// sumVersions recursively sums the version numbers of all packets
func sumVersions(p Packet) uint {
	sum := p.Version
	for _, subPkt := range p.SubPackets {
		sum += sumVersions(subPkt)
	}
	return sum
}

// evaluatePacket evaluates the packet expression based on its type ID
func evaluatePacket(p Packet) uint {
	switch p.TypeID {
	case 0: // sum
		sum := uint(0)
		for _, subPkt := range p.SubPackets {
			sum += evaluatePacket(subPkt)
		}
		return sum

	case 1: // product
		if len(p.SubPackets) == 0 {
			return 0
		}
		product := uint(1)
		for _, subPkt := range p.SubPackets {
			product *= evaluatePacket(subPkt)
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
