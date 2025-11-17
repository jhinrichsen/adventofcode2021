package adventofcode2021

import (
	"strconv"
	"strings"
)

type Point3D struct {
	x, y, z int
}

type Scanner struct {
	id      int
	beacons []Point3D
}

// parseDay19 parses scanner input
func parseDay19(lines []string) []Scanner {
	var scanners []Scanner
	var current *Scanner

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		if strings.HasPrefix(line, "---") {
			if current != nil {
				scanners = append(scanners, *current)
			}
			// Extract scanner ID
			parts := strings.Fields(line)
			if len(parts) >= 3 {
				id := 0
				if val, err := strconv.Atoi(parts[2]); err == nil {
					id = val
				}
				current = &Scanner{id: id}
			}
			continue
		}

		// Parse beacon coordinates
		parts := strings.Split(line, ",")
		if len(parts) == 3 {
			x, errX := strconv.Atoi(parts[0])
			y, errY := strconv.Atoi(parts[1])
			z, errZ := strconv.Atoi(parts[2])
			if errX == nil && errY == nil && errZ == nil {
				if current != nil {
					current.beacons = append(current.beacons, Point3D{x, y, z})
				}
			}
		}
	}

	if current != nil {
		scanners = append(scanners, *current)
	}

	return scanners
}

// All 24 rotation functions
func allRotations() []func(Point3D) Point3D {
	return []func(Point3D) Point3D{
		// Facing +x (4 rotations)
		func(p Point3D) Point3D { return Point3D{p.x, p.y, p.z} },
		func(p Point3D) Point3D { return Point3D{p.x, -p.z, p.y} },
		func(p Point3D) Point3D { return Point3D{p.x, -p.y, -p.z} },
		func(p Point3D) Point3D { return Point3D{p.x, p.z, -p.y} },
		// Facing -x (4 rotations)
		func(p Point3D) Point3D { return Point3D{-p.x, -p.y, p.z} },
		func(p Point3D) Point3D { return Point3D{-p.x, -p.z, -p.y} },
		func(p Point3D) Point3D { return Point3D{-p.x, p.y, -p.z} },
		func(p Point3D) Point3D { return Point3D{-p.x, p.z, p.y} },
		// Facing +y (4 rotations)
		func(p Point3D) Point3D { return Point3D{-p.y, p.x, p.z} },
		func(p Point3D) Point3D { return Point3D{p.z, p.x, p.y} },
		func(p Point3D) Point3D { return Point3D{p.y, p.x, -p.z} },
		func(p Point3D) Point3D { return Point3D{-p.z, p.x, -p.y} },
		// Facing -y (4 rotations)
		func(p Point3D) Point3D { return Point3D{p.y, -p.x, p.z} },
		func(p Point3D) Point3D { return Point3D{p.z, -p.x, -p.y} },
		func(p Point3D) Point3D { return Point3D{-p.y, -p.x, -p.z} },
		func(p Point3D) Point3D { return Point3D{-p.z, -p.x, p.y} },
		// Facing +z (4 rotations)
		func(p Point3D) Point3D { return Point3D{-p.z, p.y, p.x} },
		func(p Point3D) Point3D { return Point3D{-p.y, -p.z, p.x} },
		func(p Point3D) Point3D { return Point3D{p.z, -p.y, p.x} },
		func(p Point3D) Point3D { return Point3D{p.y, p.z, p.x} },
		// Facing -z (4 rotations)
		func(p Point3D) Point3D { return Point3D{p.z, p.y, -p.x} },
		func(p Point3D) Point3D { return Point3D{-p.y, p.z, -p.x} },
		func(p Point3D) Point3D { return Point3D{-p.z, -p.y, -p.x} },
		func(p Point3D) Point3D { return Point3D{p.y, -p.z, -p.x} },
	}
}

// subtract3D returns p1 - p2
func subtract3D(p1, p2 Point3D) Point3D {
	return Point3D{p1.x - p2.x, p1.y - p2.y, p1.z - p2.z}
}

// add3D returns p1 + p2
func add3D(p1, p2 Point3D) Point3D {
	return Point3D{p1.x + p2.x, p1.y + p2.y, p1.z + p2.z}
}

// tryMatch attempts to match two scanners, returns (rotation_index, offset, success)
func tryMatch(s1, s2 Scanner) (int, Point3D, bool) {
	rotations := allRotations()

	// Try each rotation
	for rotIdx, rot := range rotations {
		// Rotate all beacons in s2
		rotated := make([]Point3D, len(s2.beacons))
		for i, b := range s2.beacons {
			rotated[i] = rot(b)
		}

		// Try all pairs of beacons as potential matches
		// If they match, the offset should be the same for at least 12 beacons
		offsetCounts := make(map[Point3D]int)

		for _, b1 := range s1.beacons {
			for _, b2 := range rotated {
				offset := subtract3D(b1, b2)
				offsetCounts[offset]++
				if offsetCounts[offset] >= 12 {
					return rotIdx, offset, true
				}
			}
		}
	}

	return 0, Point3D{}, false
}

// Day19 solves day 19 puzzle
func Day19(lines []string, part1 bool) uint {
	scanners := parseDay19(lines)
	if len(scanners) == 0 {
		return 0
	}

	// Track which scanners have been positioned
	positioned := make(map[int]bool)
	positioned[0] = true

	// Store scanner positions and rotations relative to scanner 0
	scannerPos := make(map[int]Point3D)
	scannerPos[0] = Point3D{0, 0, 0}
	scannerRot := make(map[int]int)
	scannerRot[0] = 0

	// Keep trying to match unpositioned scanners with positioned ones
	rotations := allRotations()
	for len(positioned) < len(scanners) {
		matched := false

		for i := range scanners {
			if positioned[i] {
				continue
			}

			// Try to match scanner i with any positioned scanner
			for j := range scanners {
				if !positioned[j] {
					continue
				}

				// Transform scanner j's beacons to absolute coordinates
				absBeacons := make([]Point3D, len(scanners[j].beacons))
				jRot := rotations[scannerRot[j]]
				jPos := scannerPos[j]
				for k, b := range scanners[j].beacons {
					absBeacons[k] = add3D(jRot(b), jPos)
				}
				tempScanner := Scanner{beacons: absBeacons}

				// Try to match scanner i with these absolute coordinates
				rotIdx, offset, success := tryMatch(tempScanner, scanners[i])
				if success {
					positioned[i] = true
					scannerPos[i] = offset
					scannerRot[i] = rotIdx
					matched = true
					break
				}
			}

			if matched {
				break
			}
		}

		if !matched {
			// Cannot match any more scanners
			break
		}
	}

	// Collect all unique beacons in absolute coordinates
	allBeacons := make(map[Point3D]bool)
	for i := range scanners {
		if !positioned[i] {
			continue
		}

		rot := rotations[scannerRot[i]]
		pos := scannerPos[i]

		for _, b := range scanners[i].beacons {
			abs := add3D(rot(b), pos)
			allBeacons[abs] = true
		}
	}

	if part1 {
		return uint(len(allBeacons))
	}

	// Part 2: Calculate max Manhattan distance between scanners
	maxDist := 0
	positions := make([]Point3D, 0, len(scannerPos))
	for _, pos := range scannerPos {
		positions = append(positions, pos)
	}

	for i := range positions {
		for j := i + 1; j < len(positions); j++ {
			p1 := positions[i]
			p2 := positions[j]
			dist := abs(p1.x-p2.x) + abs(p1.y-p2.y) + abs(p1.z-p2.z)
			if dist > maxDist {
				maxDist = dist
			}
		}
	}

	return uint(maxDist)
}
