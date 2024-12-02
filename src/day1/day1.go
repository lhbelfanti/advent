package day1

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Day1 struct{}

func (d Day1) Part1() {
	file, err := os.Open("src/day1/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

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
	file, err := os.Open("src/day1/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	list1 := make([]int, 0, 1000)
	list2 := make([]int, 0, 1000)

	for scanner.Scan() {
		elements := strings.Split(scanner.Text(), "   ")
		elementList1, _ := strconv.Atoi(elements[0])
		elementList2, _ := strconv.Atoi(elements[1])
		list1 = append(list1, elementList1)
		list2 = append(list2, elementList2)
	}

}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
