package adventofcode2021

import (
	"strconv"
	"strings"
)

type Cuboid struct {
	x1, x2 int
	y1, y2 int
	z1, z2 int
	on     bool
}

type WeightedCuboid struct {
	x1, x2 int
	y1, y2 int
	z1, z2 int
	weight int
}

func parseDay22(lines []string) []Cuboid {
	var cuboids []Cuboid

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		on := strings.HasPrefix(line, "on")
		line = strings.TrimPrefix(line, "on ")
		line = strings.TrimPrefix(line, "off ")

		// Parse x=a..b,y=c..d,z=e..f
		parts := strings.Split(line, ",")
		if len(parts) != 3 {
			continue
		}

		var x1, x2, y1, y2, z1, z2 int

		for _, part := range parts {
			kv := strings.Split(part, "=")
			if len(kv) != 2 {
				continue
			}

			axis := kv[0]
			rangeStr := kv[1]
			rangeParts := strings.Split(rangeStr, "..")
			if len(rangeParts) != 2 {
				continue
			}

			start, err1 := strconv.Atoi(rangeParts[0])
			end, err2 := strconv.Atoi(rangeParts[1])
			if err1 != nil || err2 != nil {
				continue
			}

			switch axis {
			case "x":
				x1, x2 = start, end
			case "y":
				y1, y2 = start, end
			case "z":
				z1, z2 = start, end
			}
		}

		cuboids = append(cuboids, Cuboid{x1, x2, y1, y2, z1, z2, on})
	}

	return cuboids
}

// intersects checks if two cuboids overlap
func intersects(a, b WeightedCuboid) bool {
	return a.x1 <= b.x2 && a.x2 >= b.x1 &&
		a.y1 <= b.y2 && a.y2 >= b.y1 &&
		a.z1 <= b.z2 && a.z2 >= b.z1
}

// intersection returns the intersection of two cuboids with the given weight
func intersection(a, b WeightedCuboid, weight int) WeightedCuboid {
	return WeightedCuboid{
		x1:     max(a.x1, b.x1),
		x2:     min(a.x2, b.x2),
		y1:     max(a.y1, b.y1),
		y2:     min(a.y2, b.y2),
		z1:     max(a.z1, b.z1),
		z2:     min(a.z2, b.z2),
		weight: weight,
	}
}

// volume calculates the volume of a cuboid
func volume(c WeightedCuboid) int64 {
	return int64(c.x2-c.x1+1) * int64(c.y2-c.y1+1) * int64(c.z2-c.z1+1)
}

// Day22 solves day 22 puzzle
func Day22(lines []string, part1 bool) uint {
	cuboids := parseDay22(lines)

	if part1 {
		// For Part 1, only consider cubes in -50..50 range
		// Use a map to track which cubes are on
		onCubes := make(map[Point3D]bool)

		for _, c := range cuboids {
			// Clamp to -50..50 range
			x1 := max(-50, c.x1)
			x2 := min(50, c.x2)
			y1 := max(-50, c.y1)
			y2 := min(50, c.y2)
			z1 := max(-50, c.z1)
			z2 := min(50, c.z2)

			// Skip if completely outside range
			if x1 > 50 || x2 < -50 || y1 > 50 || y2 < -50 || z1 > 50 || z2 < -50 {
				continue
			}

			// Apply the on/off command to each cube in the range
			for x := x1; x <= x2; x++ {
				for y := y1; y <= y2; y++ {
					for z := z1; z <= z2; z++ {
						if c.on {
							onCubes[Point3D{x, y, z}] = true
						} else {
							delete(onCubes, Point3D{x, y, z})
						}
					}
				}
			}
		}

		return uint(len(onCubes))
	}

	// Part 2: Use inclusion-exclusion with weighted cuboids
	var weightedCuboids []WeightedCuboid

	for _, c := range cuboids {
		newCuboid := WeightedCuboid{c.x1, c.x2, c.y1, c.y2, c.z1, c.z2, 1}

		// For all existing cuboids that overlap with the new one,
		// add their intersection with opposite weight to fix double-counting
		var toAdd []WeightedCuboid
		for _, existing := range weightedCuboids {
			if intersects(existing, newCuboid) {
				overlap := intersection(existing, newCuboid, -existing.weight)
				toAdd = append(toAdd, overlap)
			}
		}

		// Add all the overlap cuboids
		weightedCuboids = append(weightedCuboids, toAdd...)

		// If this is an "on" instruction, also add the original cuboid
		if c.on {
			weightedCuboids = append(weightedCuboids, newCuboid)
		}
	}

	// Calculate total volume
	var total int64
	for _, wc := range weightedCuboids {
		total += volume(wc) * int64(wc.weight)
	}

	return uint(total)
}
