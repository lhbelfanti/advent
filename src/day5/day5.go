package day5

import (
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

	correctUpdates := make([][]int, 0)
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
		}

	}

	var sum int
	for _, correctUpdate := range correctUpdates {
		sum += correctUpdate[len(correctUpdate)/2]
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("The answer is: %d\n", sum)
}

func (d Day5) Part2() {
	file, scanner := reader.Read("src/day5/input.txt")
	defer file.Close()

	for scanner.Scan() {

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("The answer is: %d\n", 0)
}

func isContained(rule []int, numbers []int) bool {
	for _, n := range numbers {
		if !slices.Contains(rule, n) {
			return false
		}
	}

	return true
}
