package day7

import (
	"advent2024/src/reader"
	"bufio"
	"fmt"
	"log"
)

type Day7 struct{}

func (d Day7) Part1() {
	file, scanner := reader.Read("src/day7/input.txt")
	defer file.Close()

	scan(scanner)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("The answer is: %d\n", 0)
}

func (d Day7) Part2() {
	file, scanner := reader.Read("src/day7/input.txt")
	defer file.Close()

	scan(scanner)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("The answer is: %d\n", 0)
}

func scan(scanner *bufio.Scanner) {

}
