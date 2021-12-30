package day25

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type (
	Day25 struct{}

	Grid [][]string
)

func (d Day25) Part1() {
	grid := readFile()

	var count int
	for ok := true; ok; grid, ok = step(grid, 1) {
		count++
	}

	fmt.Printf("The answer is: %d\n", count)
}

func (d Day25) Part2() {
	fmt.Printf("The answer is: Click on the link")
}

func readFile() Grid {
	file, err := os.Open("src/day25/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	data := make([]string, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		data = append(data, scanner.Text())
	}

	var grid = make(Grid, len(data))
	for y, row := range data {
		grid[y] = make([]string, len(row))
		for x, ch := range row {
			if ch != '.' {
				grid[y][x] = string(ch)
			}
		}
	}

	return grid
}

func step(g Grid, n int) (next Grid, moved bool) {
	var m bool
	next = make(Grid, len(g))
	for y := range g {
		next[y] = make([]string, len(g[y]))
	}

	for y, row := range g {
		for x, ch := range row {
			switch ch {
			case "":
				continue
			case "v":
				if n == 1 {
					next[y][x] = "v"
				} else {
					idx := (y + 1) % len(g)
					if g[idx][x] == "" {
						next[idx][x] = "v"
						moved = true
					} else {
						next[y][x] = "v"
					}
				}
			case ">":
				if n == 1 {
					idx := (x + 1) % len(row)
					if g[y][idx] == "" {
						next[y][idx] = ">"
						moved = true
					} else {
						next[y][x] = ">"
					}
				} else {
					next[y][x] = ">"
				}
			}
		}
	}

	if n == 1 {
		next, m = step(next, 2)
		if !moved {
			moved = m
		}
	}

	return next, moved
}

func (g Grid) String() string {
	var sb strings.Builder
	for _, row := range g {
		for _, ch := range row {
			if ch == "" {
				sb.WriteByte('.')
			} else {
				sb.WriteString(ch)
			}
		}
		sb.WriteString("\n")
	}

	return sb.String()
}
