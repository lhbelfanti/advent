package day23

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type (
	Day23 struct{}
)

func (d Day23) Part1() {
	lines := readFile()
	fmt.Printf("The answer is: %d\n", len(lines))
}

func (d Day23) Part2() {
	lines := readFile()
	fmt.Printf("The answer is: %d\n", len(lines))
}

func readFile() []string {
	file, err := os.Open("src/day23/input.txt")
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
