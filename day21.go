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

	// Part 2 would use quantum dice
	return 0
}
