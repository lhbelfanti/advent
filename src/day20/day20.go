package day20

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type (
	Day20 struct{}

	xy struct {
		x, y int
	}
)

func (d Day20) Part1() {
	data := readFile()

	grid := NewGrid()
	for y, row := range data[2:] {
		for x, v := range row {
			if v == '#' {
				grid.Flag(xy{x, y})
			}
		}
	}

	fmt.Printf("The answer is: %d\n", len(grid.Step(data[0]).Step(data[0]).grid))
}

func (d Day20) Part2() {
	data := readFile()

	grid := NewGrid()
	for y, row := range data[2:] {
		for x, v := range row {
			if v == '#' {
				grid.Flag(xy{x, y})
			}
		}
	}

	for i := 0; i < 50; i++ {
		grid = grid.Step(data[0])
	}

	fmt.Printf("The answer is: %d\n", len(grid.grid))
}

func readFile() []string {
	file, err := os.Open("src/day20/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	data := make([]string, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		data = append(data, scanner.Text())
	}

	return data
}
