package adventofcode2021

import (
	"strconv"
	"strings"
)

type Point3D struct {
	x, y, z int
}

type Scanner struct {
	id          int
	beacons     []Point3D
	// distances contains all pairwise squared distances between beacons (rotation-invariant)
	distances   map[int]int
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

	// Compute distance fingerprints for each scanner
	for i := range scanners {
		scanners[i].distances = computeDistances(scanners[i].beacons)
	}

	return scanners
}

// computeDistances computes all pairwise squared distances between beacons
func computeDistances(beacons []Point3D) map[int]int {
	distances := make(map[int]int)
	for i := range beacons {
		for j := i + 1; j < len(beacons); j++ {
			dx := beacons[i].x - beacons[j].x
			dy := beacons[i].y - beacons[j].y
			dz := beacons[i].z - beacons[j].z
			distSq := dx*dx + dy*dy + dz*dz
			distances[distSq]++
		}
	}
	return distances
}

// rotate applies one of 24 rotations to a point
func rotate(p Point3D, rotIdx int) Point3D {
	switch rotIdx {
	// Facing +x (4 rotations)
	case 0:
		return Point3D{p.x, p.y, p.z}
	case 1:
		return Point3D{p.x, -p.z, p.y}
	case 2:
		return Point3D{p.x, -p.y, -p.z}
	case 3:
		return Point3D{p.x, p.z, -p.y}
	// Facing -x (4 rotations)
	case 4:
		return Point3D{-p.x, -p.y, p.z}
	case 5:
		return Point3D{-p.x, -p.z, -p.y}
	case 6:
		return Point3D{-p.x, p.y, -p.z}
	case 7:
		return Point3D{-p.x, p.z, p.y}
	// Facing +y (4 rotations)
	case 8:
		return Point3D{-p.y, p.x, p.z}
	case 9:
		return Point3D{p.z, p.x, p.y}
	case 10:
		return Point3D{p.y, p.x, -p.z}
	case 11:
		return Point3D{-p.z, p.x, -p.y}
	// Facing -y (4 rotations)
	case 12:
		return Point3D{p.y, -p.x, p.z}
	case 13:
		return Point3D{p.z, -p.x, -p.y}
	case 14:
		return Point3D{-p.y, -p.x, -p.z}
	case 15:
		return Point3D{-p.z, -p.x, p.y}
	// Facing +z (4 rotations)
	case 16:
		return Point3D{-p.z, p.y, p.x}
	case 17:
		return Point3D{-p.y, -p.z, p.x}
	case 18:
		return Point3D{p.z, -p.y, p.x}
	case 19:
		return Point3D{p.y, p.z, p.x}
	// Facing -z (4 rotations)
	case 20:
		return Point3D{p.z, p.y, -p.x}
	case 21:
		return Point3D{-p.y, p.z, -p.x}
	case 22:
		return Point3D{-p.z, -p.y, -p.x}
	case 23:
		return Point3D{p.y, -p.z, -p.x}
	default:
		return p
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

// mightOverlap checks if two scanners might have overlapping beacons based on distance fingerprints
func mightOverlap(s1, s2 Scanner) bool {
	// For 12 overlapping beacons, we need at least C(12,2) = 66 common distances
	const minCommonDistances = 66

	commonCount := 0
	for dist := range s1.distances {
		if s2.distances[dist] > 0 {
			// Count the minimum of the two counts for this distance
			count := s1.distances[dist]
			if s2.distances[dist] < count {
				count = s2.distances[dist]
			}
			commonCount += count
			if commonCount >= minCommonDistances {
				return true
			}
		}
	}
	return false
}

// tryMatch attempts to match two scanners, returns (rotation_index, offset, success)
func tryMatch(s1, s2 Scanner, rotated []Point3D, offsetCounts map[Point3D]int) (int, Point3D, bool) {
	// Try each rotation
	for rotIdx := range 24 {
		// Rotate all beacons in s2
		for i, b := range s2.beacons {
			rotated[i] = rotate(b, rotIdx)
		}

		// Try all pairs of beacons as potential matches
		// If they match, the offset should be the same for at least 12 beacons
		clear(offsetCounts)

		for _, b1 := range s1.beacons {
			for i := range len(s2.beacons) {
				offset := subtract3D(b1, rotated[i])
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

	// Cache absolute beacon coordinates for positioned scanners
	absCache := make(map[int][]Point3D)
	absCache[0] = make([]Point3D, len(scanners[0].beacons))
	copy(absCache[0], scanners[0].beacons)

	// Pre-allocate working buffers for tryMatch
	maxBeacons := 0
	for i := range scanners {
		if len(scanners[i].beacons) > maxBeacons {
			maxBeacons = len(scanners[i].beacons)
		}
	}
	rotated := make([]Point3D, maxBeacons)
	offsetCounts := make(map[Point3D]int, maxBeacons*maxBeacons)

	// Keep trying to match unpositioned scanners with positioned ones
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

				// Quick check: do these scanners might overlap based on distance fingerprints?
				if !mightOverlap(scanners[j], scanners[i]) {
					continue
				}

				// Use cached absolute coordinates
				tempScanner := Scanner{beacons: absCache[j]}

				// Try to match scanner i with these absolute coordinates
				rotIdx, offset, success := tryMatch(tempScanner, scanners[i], rotated, offsetCounts)
				if success {
					positioned[i] = true
					scannerPos[i] = offset
					scannerRot[i] = rotIdx

					// Cache absolute coordinates for scanner i
					absBeacons := make([]Point3D, len(scanners[i].beacons))
					for k, b := range scanners[i].beacons {
						absBeacons[k] = add3D(rotate(b, rotIdx), offset)
					}
					absCache[i] = absBeacons

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

		for _, b := range absCache[i] {
			allBeacons[b] = true
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
