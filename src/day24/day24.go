package day24

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type (
	Day24 struct{}
)

func (d Day24) Part1() {
	lines := readFile()
	fmt.Printf("The answer is: %d\n", len(lines))
}

func (d Day24) Part2() {
	lines := readFile()
	fmt.Printf("The answer is: %d\n", len(lines))
}

func readFile() []string {
	file, err := os.Open("src/day24/input.txt")
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
