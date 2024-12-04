package day3

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"

	"advent2024/src/reader"
)

type Day3 struct{}

func (d Day3) Part1() {
	file, scanner := reader.Read("src/day3/input.txt")
	defer file.Close()

	var corruptedMemory string

	for scanner.Scan() {
		corruptedMemory += scanner.Text()
	}

	r, _ := regexp.Compile("mul\\([0-9]{1,3},[0-9]{1,3}\\)")

	var sum int
	mulOps := r.FindAllString(corruptedMemory, -1)
	for _, mul := range mulOps {
		numbers := strings.Split(strings.Replace(strings.Replace(mul, "mul(", "", -1), ")", "", -1), ",")
		n1, _ := strconv.Atoi(numbers[0])
		n2, _ := strconv.Atoi(numbers[1])
		sum += n1 * n2
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("The answer is: %d\n", sum)
}

func (d Day3) Part2() {
	file, scanner := reader.Read("src/day3/input.txt")
	defer file.Close()

	var corruptedMemory string

	for scanner.Scan() {
		corruptedMemory += scanner.Text()
	}

	r, _ := regexp.Compile("(mul\\([0-9]{1,3},[0-9]{1,3}\\))|do\\(\\)|don't\\(\\)")

	var sum int
	operations := r.FindAllString(corruptedMemory, -1)
	mulEnabled := true
	for _, op := range operations {
		if strings.Contains(op, "mul") && mulEnabled {
			numbers := strings.Split(strings.Replace(strings.Replace(op, "mul(", "", -1), ")", "", -1), ",")
			n1, _ := strconv.Atoi(numbers[0])
			n2, _ := strconv.Atoi(numbers[1])
			sum += n1 * n2
		} else if op == "don't()" {
			mulEnabled = false
		} else if op == "do()" {
			mulEnabled = true
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("The answer is: %d\n", sum)
}
