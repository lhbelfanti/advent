package day4

import (
	"fmt"
	"log"

	"advent2024/src/reader"
)

type Day4 struct{}

func (d Day4) Part1() {
	file, scanner := reader.Read("src/day4/input.txt")
	defer file.Close()

	for scanner.Scan() {

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("The answer is: %d\n", 0)
}

func (d Day4) Part2() {
	file, scanner := reader.Read("src/day4/input.txt")
	defer file.Close()

	for scanner.Scan() {

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("The answer is: %d\n", 0)
}
