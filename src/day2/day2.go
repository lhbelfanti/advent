package day2

import (
"bufio"
"fmt"
"log"
"os"
"strconv"
	"strings"
)

type Day2 struct {}

func (d Day2) Part1() {
	file, err := os.Open("src/day2/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var horizontalPosition, depth int

	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		value, _ := strconv.Atoi(fields[1])
		switch fields[0] {
		case "forward":
			horizontalPosition += value
		case "up":
			depth -= value
		case "down":
			depth += value
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("The answer is: %d\n", horizontalPosition * depth)
}

func (d Day2) Part2() {
	file, err := os.Open("src/day2/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var horizontalPosition, depth, aim int

	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		value, _ := strconv.Atoi(fields[1])
		switch fields[0] {
		case "forward":
			horizontalPosition += value
			depth += aim * value
		case "up":
			aim -= value
		case "down":
			aim += value
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("The answer is: %d\n", horizontalPosition * depth)
}
