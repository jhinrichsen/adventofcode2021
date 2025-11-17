package adventofcode2021

import (
	"container/heap"
)

type AmphipodState struct {
	hallway   [11]byte    // hallway positions (0=empty, 'A'/'B'/'C'/'D')
	rooms     [4][4]byte  // 4 rooms, each with depth up to 4
	cost      uint        // total cost to reach this state
	roomDepth int         // depth of rooms (2 for part1, 4 for part2)
}

func (s AmphipodState) String() string {
	result := string(s.hallway[:])
	for i := range s.rooms {
		result += string(s.rooms[i][:s.roomDepth])
	}
	return result
}

func (s AmphipodState) isGoal() bool {
	for roomIdx := range s.rooms {
		expected := byte('A' + roomIdx)
		for depth := 0; depth < s.roomDepth; depth++ {
			if s.rooms[roomIdx][depth] != expected {
				return false
			}
		}
	}
	return true
}

func energyCost(amphipod byte) uint {
	switch amphipod {
	case 'A':
		return 1
	case 'B':
		return 10
	case 'C':
		return 100
	case 'D':
		return 1000
	}
	return 0
}

func roomIndex(amphipod byte) int {
	return int(amphipod - 'A')
}

func roomHallwayPos(roomIdx int) int {
	return 2 + roomIdx*2
}

// Priority queue for Dijkstra
type StateQueue []AmphipodState

func (pq StateQueue) Len() int           { return len(pq) }
func (pq StateQueue) Less(i, j int) bool { return pq[i].cost < pq[j].cost }
func (pq StateQueue) Swap(i, j int)      { pq[i], pq[j] = pq[j], pq[i] }
func (pq *StateQueue) Push(x interface{}) {
	*pq = append(*pq, x.(AmphipodState))
}
func (pq *StateQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

func parseDay23(lines []string, part1 bool) AmphipodState {
	var state AmphipodState

	// Set room depth
	if part1 {
		state.roomDepth = 2
	} else {
		state.roomDepth = 4
	}

	// Initialize empty
	for i := range state.hallway {
		state.hallway[i] = '.'
	}
	for i := range state.rooms {
		for j := range state.rooms[i] {
			state.rooms[i][j] = '.'
		}
	}

	if part1 {
		// Parse the diagram for Part 1 (depth 2)
		if len(lines) >= 5 {
			// Line 3: ###B#C#B#D###
			line3 := lines[2]
			if len(line3) >= 10 {
				state.rooms[0][0] = line3[3]
				state.rooms[1][0] = line3[5]
				state.rooms[2][0] = line3[7]
				state.rooms[3][0] = line3[9]
			}

			// Line 4:   #A#D#C#A#
			line4 := lines[3]
			if len(line4) >= 10 {
				state.rooms[0][1] = line4[3]
				state.rooms[1][1] = line4[5]
				state.rooms[2][1] = line4[7]
				state.rooms[3][1] = line4[9]
			}
		}
	} else {
		// Parse the diagram for Part 2 (depth 4)
		// Insert the fixed rows between original rows
		if len(lines) >= 5 {
			// Line 3: ###B#C#B#D### (original first row)
			line3 := lines[2]
			if len(line3) >= 10 {
				state.rooms[0][0] = line3[3]
				state.rooms[1][0] = line3[5]
				state.rooms[2][0] = line3[7]
				state.rooms[3][0] = line3[9]
			}

			// Fixed rows to insert:
			// Row 2: #D#C#B#A#
			state.rooms[0][1] = 'D'
			state.rooms[1][1] = 'C'
			state.rooms[2][1] = 'B'
			state.rooms[3][1] = 'A'

			// Row 3: #D#B#A#C#
			state.rooms[0][2] = 'D'
			state.rooms[1][2] = 'B'
			state.rooms[2][2] = 'A'
			state.rooms[3][2] = 'C'

			// Line 4:   #A#D#C#A# (original second row)
			line4 := lines[3]
			if len(line4) >= 10 {
				state.rooms[0][3] = line4[3]
				state.rooms[1][3] = line4[5]
				state.rooms[2][3] = line4[7]
				state.rooms[3][3] = line4[9]
			}
		}
	}

	return state
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// Check if path in hallway is clear
func isPathClear(hallway [11]byte, from, to int) bool {
	if from > to {
		from, to = to, from
	}
	for i := from; i <= to; i++ {
		if hallway[i] != '.' {
			return false
		}
	}
	return true
}

// Generate all valid next states
func getNextStates(state AmphipodState) []AmphipodState {
	var next []AmphipodState

	// Try moving from room to hallway
	for roomIdx := range state.rooms {
		roomPos := roomHallwayPos(roomIdx)

		// Find top amphipod in room
		var amphipod byte
		var depth int
		found := false
		for d := 0; d < state.roomDepth; d++ {
			if state.rooms[roomIdx][d] != '.' {
				amphipod = state.rooms[roomIdx][d]
				depth = d
				found = true
				break
			}
		}
		if !found {
			continue // room is empty
		}

		// Check if this amphipod should stay (it's in the right room and all below are correct)
		targetRoom := roomIndex(amphipod)
		if targetRoom == roomIdx {
			allBelowCorrect := true
			for d := depth; d < state.roomDepth; d++ {
				if state.rooms[roomIdx][d] != amphipod && state.rooms[roomIdx][d] != '.' {
					allBelowCorrect = false
					break
				}
			}
			if allBelowCorrect {
				continue // can stay
			}
		}

		// Try moving to each valid hallway position
		for hallwayPos := 0; hallwayPos < 11; hallwayPos++ {
			// Can't stop directly above rooms
			if hallwayPos == 2 || hallwayPos == 4 || hallwayPos == 6 || hallwayPos == 8 {
				continue
			}

			// Check if path is clear
			from := min(roomPos, hallwayPos)
			to := max(roomPos, hallwayPos)
			pathClear := true
			for i := from; i <= to; i++ {
				if state.hallway[i] != '.' {
					pathClear = false
					break
				}
			}
			if !pathClear {
				continue
			}

			// Create new state
			newState := state
			newState.hallway[hallwayPos] = amphipod
			newState.rooms[roomIdx][depth] = '.'
			steps := uint(abs(roomPos-hallwayPos) + depth + 1)
			newState.cost = state.cost + steps*energyCost(amphipod)
			next = append(next, newState)
		}
	}

	// Try moving from hallway to room
	for hallwayPos := range state.hallway {
		amphipod := state.hallway[hallwayPos]
		if amphipod == '.' {
			continue
		}

		targetRoom := roomIndex(amphipod)
		roomPos := roomHallwayPos(targetRoom)

		// Check if room is ready (only contains same type or empty)
		roomReady := true
		for d := 0; d < state.roomDepth; d++ {
			occupant := state.rooms[targetRoom][d]
			if occupant != '.' && occupant != amphipod {
				roomReady = false
				break
			}
		}
		if !roomReady {
			continue
		}

		// Check if path is clear (excluding current position)
		from := min(roomPos, hallwayPos)
		to := max(roomPos, hallwayPos)
		pathClear := true
		for i := from; i <= to; i++ {
			if i == hallwayPos {
				continue
			}
			if state.hallway[i] != '.' {
				pathClear = false
				break
			}
		}
		if !pathClear {
			continue
		}

		// Find deepest empty position in target room
		depth := -1
		for d := state.roomDepth - 1; d >= 0; d-- {
			if state.rooms[targetRoom][d] == '.' {
				depth = d
				break
			}
		}
		if depth == -1 {
			continue // room is full
		}

		// Create new state
		newState := state
		newState.hallway[hallwayPos] = '.'
		newState.rooms[targetRoom][depth] = amphipod
		steps := uint(abs(roomPos-hallwayPos) + depth + 1)
		newState.cost = state.cost + steps*energyCost(amphipod)
		next = append(next, newState)
	}

	return next
}

// Day23 solves day 23 puzzle
func Day23(lines []string, part1 bool) uint {
	initialState := parseDay23(lines, part1)

	// Dijkstra's algorithm
	pq := &StateQueue{initialState}
	heap.Init(pq)
	visited := make(map[string]bool)

	for pq.Len() > 0 {
		current := heap.Pop(pq).(AmphipodState)

		if current.isGoal() {
			return current.cost
		}

		stateKey := current.String()
		if visited[stateKey] {
			continue
		}
		visited[stateKey] = true

		for _, nextState := range getNextStates(current) {
			if !visited[nextState.String()] {
				heap.Push(pq, nextState)
			}
		}
	}

	return 0
}
