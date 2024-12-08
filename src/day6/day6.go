package day6

import (
	"fmt"
	"log"

	"advent2024/src/reader"
)

type Day6 struct{}

func (d Day6) Part1() {
	file, scanner := reader.Read("src/day6/input.txt")
	defer file.Close()

	for scanner.Scan() {

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("The answer is: %d\n", 0)
}

func (d Day6) Part2() {
	file, scanner := reader.Read("src/day6/input.txt")
	defer file.Close()

	for scanner.Scan() {

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("The answer is: %d\n", 0)
}
