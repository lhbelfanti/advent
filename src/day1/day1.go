package day1

import (
	"bufio"
	"log"
	"os"
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
