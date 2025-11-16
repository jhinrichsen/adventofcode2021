package adventofcode2021

import (
	"container/heap"
)

type AmphipodState struct {
	hallway [11]byte        // hallway positions (0=empty, 'A'/'B'/'C'/'D')
	rooms   [4][2]byte      // 4 rooms, each with depth 2
	cost    uint            // total cost to reach this state
}

func (s AmphipodState) String() string {
	return string(s.hallway[:]) + string(s.rooms[0][:]) + string(s.rooms[1][:]) + string(s.rooms[2][:]) + string(s.rooms[3][:])
}

func (s AmphipodState) isGoal() bool {
	return s.rooms[0][0] == 'A' && s.rooms[0][1] == 'A' &&
		s.rooms[1][0] == 'B' && s.rooms[1][1] == 'B' &&
		s.rooms[2][0] == 'C' && s.rooms[2][1] == 'C' &&
		s.rooms[3][0] == 'D' && s.rooms[3][1] == 'D'
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

func parseDay23(lines []string) AmphipodState {
	var state AmphipodState

	// Initialize empty
	for i := range state.hallway {
		state.hallway[i] = '.'
	}
	for i := range state.rooms {
		state.rooms[i][0] = '.'
		state.rooms[i][1] = '.'
	}

	// Parse the diagram
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
		if state.rooms[roomIdx][0] != '.' {
			amphipod = state.rooms[roomIdx][0]
			depth = 0
		} else if state.rooms[roomIdx][1] != '.' {
			amphipod = state.rooms[roomIdx][1]
			depth = 1
		} else {
			continue // room is empty
		}

		// Check if this amphipod should stay (it's in the right room and all below are correct)
		targetRoom := roomIndex(amphipod)
		if targetRoom == roomIdx {
			if depth == 1 {
				continue // bottom position, stays
			}
			if state.rooms[roomIdx][1] == amphipod || state.rooms[roomIdx][1] == '.' {
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
		for _, occupant := range state.rooms[targetRoom] {
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
		var depth int
		if state.rooms[targetRoom][1] == '.' {
			depth = 1
		} else if state.rooms[targetRoom][0] == '.' {
			depth = 0
		} else {
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
	initialState := parseDay23(lines)

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
