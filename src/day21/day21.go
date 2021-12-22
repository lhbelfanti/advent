package day21

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type (
	Day21 struct{}

	Dice struct {
		role int
	}

	Win struct {
		w1, w2 int
	}

	State struct {
		pos1, pos2, score1, score2 int
	}
)

func (d Day21) Part1() {
	p1Pos, p2Pos := readFile()
	var p1Score, p2Score int
	var dice Dice
	var losingPlayerPoints int
	for {
		p1Pos = dice.roleDice(p1Pos)
		p1Score += p1Pos
		if p1Score >= 1000 {
			losingPlayerPoints = p2Score
			break
		}
		p2Pos = dice.roleDice(p2Pos)
		p2Score += p2Pos
		if p2Score >= 1000 {
			losingPlayerPoints = p1Score
			break
		}
	}

	fmt.Printf("The answer is: %d\n", losingPlayerPoints*dice.role)
}

func (d Day21) Part2() {
	pos1, pos2 := readFile()
	DP := make(map[State]Win)
	w := countWin(pos1-1, pos2-1, 0, 0, DP)
	var winner int
	if w.w1 > w.w2 {
		winner = w.w1
	} else {
		winner = w.w2
	}

	fmt.Printf("The answer is: %d\n", winner)
}

func readFile() (int, int) {
	file, err := os.Open("src/day21/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	data := make([]string, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		data = append(data, scanner.Text())
	}

	player1, _ := strconv.Atoi(strings.Split(data[0], " starting position: ")[1])
	player2, _ := strconv.Atoi(strings.Split(data[1], " starting position: ")[1])

	return player1, player2
}

func (d *Dice) roleDice(position int) int {
	move := d.role + 1 + d.role + 2 + d.role + 3
	d.role += 3
	for i := 0; i < move; i++ {
		if position == 10 {
			position = 1
		} else {
			position++
		}
	}

	return position
}

func countWin(p1, p2, s1, s2 int, DP map[State]Win) Win {
	s := State{p1, p2, s1, s2}
	if s1 >= 21 {
		return Win{1, 0}
	}
	if s2 >= 21 {
		return Win{0, 1}
	}
	if _, ok := DP[s]; ok {
		return DP[s]
	}
	answer := Win{0, 0}
	quantumDice := []int{1, 2, 3}

	for _, d1 := range quantumDice {
		for _, d2 := range quantumDice {
			for _, d3 := range quantumDice {
				newP1 := (p1 + d1 + d2 + d3) % 10
				newS1 := s1 + newP1 + 1

				w := countWin(p2, newP1, s2, newS1, DP)
				answer.w1 += w.w2
				answer.w2 += w.w1
			}
		}
	}
	DP[s] = answer
	return answer
}
