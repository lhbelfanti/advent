package day1

import (
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"

	"advent2024/src/reader"
)

type Day1 struct{}

func (d Day1) Part1() {
	file, scanner := reader.Read("src/day1/input.txt")
	defer file.Close()

	list1 := make([]int, 0, 1000)
	list2 := make([]int, 0, 1000)

	for scanner.Scan() {
		elements := strings.Split(scanner.Text(), "   ")
		elementList1, _ := strconv.Atoi(elements[0])
		elementList2, _ := strconv.Atoi(elements[1])
		list1 = append(list1, elementList1)
		list2 = append(list2, elementList2)
	}

	sort.Ints(list1)
	sort.Ints(list2)

	var sum int
	for i := range list1 {
		sum += abs(list1[i] - list2[i])
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("The answer is: %d\n", sum)
}

func (d Day1) Part2() {
	file, scanner := reader.Read("src/day1/input.txt")
	defer file.Close()

	list1 := make([]int, 0, 1000)
	list2 := make([]int, 0, 1000)

	for scanner.Scan() {
		elements := strings.Split(scanner.Text(), "   ")
		elementList1, _ := strconv.Atoi(elements[0])
		elementList2, _ := strconv.Atoi(elements[1])
		list1 = append(list1, elementList1)
		list2 = append(list2, elementList2)
	}

	sort.Ints(list1)
	sort.Ints(list2)

	var sum int
	for _, n1 := range list1 {
		multiplier := 0
		for _, n2 := range list2 {
			if n2 == n1 {
				multiplier += 1
			}

		}
		sum += n1 * multiplier
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("The answer is: %d\n", sum)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
