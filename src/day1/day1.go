package day1

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type Day1 struct{}

func (d Day1) Part1() {
	file, err := os.Open("src/day1/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	previousMeasure := -999999
	counter := 0

	for scanner.Scan() {
		currentMeasure, _ := strconv.Atoi(scanner.Text())
		if previousMeasure < currentMeasure && previousMeasure != -999999 {
			counter++
		}
		previousMeasure = currentMeasure
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("The answer is: %d\n", counter)
}

func (d Day1) Part2() {
	file, err := os.Open("src/day1/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var measurements []int
	for scanner.Scan() {
		measure, _ := strconv.Atoi(scanner.Text())
		measurements = append(measurements, measure)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	sliding := 3
	length := len(measurements)
	counter := 0
	for i := range measurements {
		if i+sliding < length {
			prev := sum(measurements[i : i+sliding])
			curr := sum(measurements[i+1 : i+1+sliding])

			if prev < curr {
				counter++
			}
		}
	}

	fmt.Printf("The answer is: %d\n", counter)
}

func sum(m []int) int {
	result := 0
	for _, numb := range m {
		result += numb
	}
	return result
}
