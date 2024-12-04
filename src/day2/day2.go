package day2

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	
	"advent2024/src/reader"
)

type Day2 struct{}

func (d Day2) Part1() {
	file, scanner := reader.Read("src/day2/input.txt")
	defer file.Close()

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
		if isLevelSafe(level) {
			count++
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
		if isLevelSafe(level) {
			count++
		} else {
			for i := 0; i < len(level); i++ {
				newLevel := removeIndex(level, i)
				if isLevelSafe(newLevel) {
					count++
					break
				}
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("The answer is: %d\n", count)
}

func isLevelSafe(level []int) bool {
	prev := 0
	for i := 0; i < len(level)-1; i++ {
		value := level[i] - level[i+1]
		if abs(value) < 1 || abs(value) > 3 {
			return false
		}

		if i > 0 {
			if prev > 0 && value < 0 {
				return false
			}

			if prev < 0 && value > 0 {
				return false
			}
		}

		prev = value
	}

	return true
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func removeIndex(s []int, index int) []int {
	ret := make([]int, 0, len(s)-1)
	ret = append(ret, s[:index]...)
	return append(ret, s[index+1:]...)
}
