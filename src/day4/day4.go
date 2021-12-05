package day4

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type (
	Day4 struct {
		numbers []string
		boards  []Board
	}

	Board struct {
		data     [][]string
		complete bool
	}
)

func (d Day4) Part1() {
	data := d.readFile()
	var score int

	for _, number := range data.numbers {
		for _, board := range data.boards {
			board.markNumber(number)
			if board.hasBingo() {
				num, _ := strconv.Atoi(number)
				score = board.getScore(num)
				break
			}
		}
		if score != 0 {
			break
		}
	}

	fmt.Printf("The answer is: %d\n", score)
}

func (d Day4) Part2() {
	data := d.readFile()
	var score int

	for _, number := range data.numbers {
		for i, board := range data.boards {
			if !board.complete {
				board.markNumber(number)

				if board.hasBingo() {
					board.complete = true
					data.boards[i] = board
					num, _ := strconv.Atoi(number)
					score = board.getScore(num)
				}
			}
		}
	}

	fmt.Printf("The answer is: %d\n", score)
}

func (d Day4) readFile() Day4 {
	file, err := os.Open("src/day4/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	d.boards = make([]Board, 0)

	var board Board
	var index int

	for scanner.Scan() {
		n := scanner.Text()
		if len(d.numbers) == 0 {
			d.numbers = strings.Split(n, ",")
			continue
		}
		if n == "" {
			if board.data != nil {
				board.complete = false
				d.boards = append(d.boards, board)
			}
			board.data = make([][]string, 0)
			index = 0
			continue
		}

		nums := strings.Fields(n)
		if len(board.data) == 0 {
			board.data = make([][]string, len(nums))
		}
		board.data[index] = append(board.data[index], nums...)
		index++
	}

	// Append the last one
	if board.data != nil {
		board.complete = false
		d.boards = append(d.boards, board)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return d
}

func (b Board) markNumber(num string) {
	for i := range b.data {
		for j := range b.data[i] {
			if b.data[i][j] == num {
				b.data[i][j] = "-1"
			}
		}
	}
}

func (b Board) hasBingo() bool {
	var bingoInRow int
	var bingoInColumn int

	for i := range b.data {
		bingoInRow = 0
		bingoInColumn = 0

		for j := range b.data[i] {
			if b.data[i][j] == "-1" {
				bingoInRow++
			}

			if b.data[j][i] == "-1" {
				bingoInColumn++
			}

			if bingoInRow == len(b.data[i]) || bingoInColumn == len(b.data[i]) {
				return true
			}
		}
	}

	return false
}

func (b Board) getScore(num int) int {
	var sum int
	for i := range b.data {
		for j := range b.data[i] {
			n, _ := strconv.Atoi(b.data[i][j])
			if n != -1 {
				sum += n
			}
		}
	}

	return sum * num
}
