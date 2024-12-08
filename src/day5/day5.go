package day5

import (
	"bufio"
	"fmt"
	"log"
	"slices"
	"strconv"
	"strings"

	"advent2024/src/reader"
)

type Day5 struct{}

func (d Day5) Part1() {
	file, scanner := reader.Read("src/day5/input.txt")
	defer file.Close()

	rules, updates := scan(scanner)

	correctUpdates, _ := processUpdates(rules, updates)

	sum := sumCorrectUpdates(correctUpdates)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("The answer is: %d\n", sum)
}

func (d Day5) Part2() {
	file, scanner := reader.Read("src/day5/input.txt")
	defer file.Close()

	rules, updates := scan(scanner)

	_, incorrectUpdates := processUpdates(rules, updates)

	ordered := false
	for !ordered {
		ordered = orderIncorrectUpdate(rules, &incorrectUpdates)
	}

	sum := sumCorrectUpdates(incorrectUpdates)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("The answer is: %d\n", sum)
}

func scan(scanner *bufio.Scanner) (map[int][]int, [][]int) {
	rules := make(map[int][]int)
	updates := make([][]int, 0)

	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "|") {
			pages := strings.Split(line, "|")
			pageX, _ := strconv.Atoi(pages[0])
			pageY, _ := strconv.Atoi(pages[1])

			if _, ok := rules[pageX]; ok {
				rules[pageX] = append(rules[pageX], pageY)
			} else {
				rules[pageX] = []int{pageY}
			}
		} else if strings.Contains(line, ",") {
			pagesPerUpdateStr := strings.Split(line, ",")
			pagesPerUpdate := make([]int, 0, len(pagesPerUpdateStr))

			for _, pageStr := range pagesPerUpdateStr {
				page, _ := strconv.Atoi(pageStr)
				pagesPerUpdate = append(pagesPerUpdate, page)
			}

			updates = append(updates, pagesPerUpdate)
		}
	}

	return rules, updates
}

func processUpdates(rules map[int][]int, updates [][]int) ([][]int, [][]int) {
	correctUpdates := make([][]int, 0)
	incorrectUpdates := make([][]int, 0)
	for _, update := range updates {
		isCorrect := true
		for i, page := range update {
			rule, ok := rules[page]
			if !ok {
				continue
			}

			beforeNumbers := update[i+1:]
			afterNumbers := update[0:i]

			// page must be before numbers in the slice 'beforeNumbers'
			for _, b := range beforeNumbers {
				if !slices.Contains(rule, b) {
					isCorrect = false
				}
			}

			// search by the negative page must not contain any numbers of the slice 'afterNumbers'
			for _, a := range afterNumbers {
				if slices.Contains(rule, a) {
					isCorrect = false
				}
			}
		}

		if isCorrect {
			correctUpdates = append(correctUpdates, update)
		} else {
			incorrectUpdates = append(incorrectUpdates, update)
		}
	}

	return correctUpdates, incorrectUpdates
}

func orderIncorrectUpdate(rules map[int][]int, incorrectUpdates *[][]int) bool {
	for idx, update := range *incorrectUpdates {
		for i, page := range update { // i.e. update = [61, 13, 29], i = 2, page = 29
			rule, ok := rules[page] // i.e. [13]
			if !ok {
				continue
			}

			afterNumbers := update[0:i] // i.e. [61, 13]

			for k, a := range afterNumbers { // i.e. k = 1
				if slices.Contains(rule, a) { // i.e. [61, 13] contains 13 --> swap 13 and 29
					// 13 = update[k]
					// 29 = update[i]
					update[i], update[k] = update[k], update[i]
					(*incorrectUpdates)[idx] = update
					return false
				}
			}

		}
	}

	return true
}

func sumCorrectUpdates(correctUpdates [][]int) int {
	var sum int
	for _, correctUpdate := range correctUpdates {
		sum += correctUpdate[len(correctUpdate)/2]
	}

	return sum
}
