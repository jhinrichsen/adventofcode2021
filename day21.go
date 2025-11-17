package adventofcode2021

import (
	"strconv"
	"strings"
)

type DeterministicDie struct {
	next  int
	rolls uint
}

func (d *DeterministicDie) roll() int {
	val := d.next
	d.next++
	if d.next > 100 {
		d.next = 1
	}
	d.rolls++
	return val
}

func (d *DeterministicDie) roll3() int {
	return d.roll() + d.roll() + d.roll()
}

type Player struct {
	pos   int
	score uint
}

func (p *Player) move(spaces int) {
	p.pos += spaces
	// Circular board 1-10
	for p.pos > 10 {
		p.pos -= 10
	}
	p.score += uint(p.pos)
}

func parseDay21(lines []string) (int, int) {
	p1, p2 := 0, 0
	for _, line := range lines {
		if strings.Contains(line, "Player 1") {
			parts := strings.Split(line, ":")
			if len(parts) == 2 {
				val, err := strconv.Atoi(strings.TrimSpace(parts[1]))
				if err == nil {
					p1 = val
				}
			}
		} else if strings.Contains(line, "Player 2") {
			parts := strings.Split(line, ":")
			if len(parts) == 2 {
				val, err := strconv.Atoi(strings.TrimSpace(parts[1]))
				if err == nil {
					p2 = val
				}
			}
		}
	}
	return p1, p2
}

// Day21 solves day 21 puzzle
func Day21(lines []string, part1 bool) uint {
	p1Pos, p2Pos := parseDay21(lines)

	if part1 {
		die := DeterministicDie{next: 1}
		players := []Player{
			{pos: p1Pos, score: 0},
			{pos: p2Pos, score: 0},
		}

		currentPlayer := 0
		for {
			roll := die.roll3()
			players[currentPlayer].move(roll)

			if players[currentPlayer].score >= 1000 {
				// Game over
				loser := 1 - currentPlayer
				return players[loser].score * die.rolls
			}

			currentPlayer = 1 - currentPlayer
		}
	}

	// Part 2: Quantum dice
	// When rolling 3-sided die 3 times, count frequencies of each sum
	rollFreq := map[int]uint{
		3: 1, // (1,1,1)
		4: 3, // (1,1,2), (1,2,1), (2,1,1)
		5: 6, // (1,1,3), (1,3,1), (3,1,1), (1,2,2), (2,1,2), (2,2,1)
		6: 7, // (1,2,3), (1,3,2), (2,1,3), (2,3,1), (3,1,2), (3,2,1), (2,2,2)
		7: 6, // (1,3,3), (3,1,3), (3,3,1), (2,2,3), (2,3,2), (3,2,2)
		8: 3, // (2,3,3), (3,2,3), (3,3,2)
		9: 1, // (3,3,3)
	}

	type GameState struct {
		p1Pos, p1Score int
		p2Pos, p2Score int
		turn           int
	}

	memo := make(map[GameState][2]uint)

	var play func(p1Pos, p1Score, p2Pos, p2Score, turn int) [2]uint
	play = func(p1Pos, p1Score, p2Pos, p2Score, turn int) [2]uint {
		// Check if game is over
		if p1Score >= 21 {
			return [2]uint{1, 0}
		}
		if p2Score >= 21 {
			return [2]uint{0, 1}
		}

		state := GameState{p1Pos, p1Score, p2Pos, p2Score, turn}
		if result, ok := memo[state]; ok {
			return result
		}

		var totalWins [2]uint

		if turn == 0 {
			// Player 1's turn
			for roll, freq := range rollFreq {
				newPos := p1Pos + roll
				for newPos > 10 {
					newPos -= 10
				}
				newScore := p1Score + newPos
				wins := play(newPos, newScore, p2Pos, p2Score, 1)
				totalWins[0] += wins[0] * freq
				totalWins[1] += wins[1] * freq
			}
		} else {
			// Player 2's turn
			for roll, freq := range rollFreq {
				newPos := p2Pos + roll
				for newPos > 10 {
					newPos -= 10
				}
				newScore := p2Score + newPos
				wins := play(p1Pos, p1Score, newPos, newScore, 0)
				totalWins[0] += wins[0] * freq
				totalWins[1] += wins[1] * freq
			}
		}

		memo[state] = totalWins
		return totalWins
	}

	wins := play(p1Pos, 0, p2Pos, 0, 0)
	if wins[0] > wins[1] {
		return wins[0]
	}
	return wins[1]
}
