package day2

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Day2 struct{}

func (d Day2) Part1() {
	file, err := os.Open("src/day2/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	levels := make([][]int, 0)

	n := 0
	for scanner.Scan() {
		elements := strings.Split(scanner.Text(), " ")
		levels = append(levels, make([]int, 0, len(elements)))
		for _, element := range elements {
			intElement, _ := strconv.Atoi(element)
			levels[n] = append(levels[n], intElement)
		}
		n++
	}

	count := 0
	for _, level := range levels {
		count++
		prev := 0
		for i := 0; i < len(level)-1; i++ {
			value := level[i] - level[i+1]
			if abs(value) < 1 || abs(value) > 3 {
				count--
				break
			}

			if i > 0 {
				if prev > 0 && value < 0 {
					count--
					break
				}

				if prev < 0 && value > 0 {
					count--
					break
				}
			}

			prev = value
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("The answer is: %d\n", count)
}

func (d Day2) Part2() {
	file, err := os.Open("src/day2/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	levels := make([][]int, 0)

	i := 0
	for scanner.Scan() {
		elements := strings.Split(scanner.Text(), " ")
		levels = append(levels, make([]int, 0))
		for _, element := range elements {
			intElement, _ := strconv.Atoi(element)
			levels[i] = append(levels[i], intElement)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("The answer is: %d\n", 0)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
