package day7

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type Day7 struct{}

func (d Day7) Part1() {
	crabs, _ := readFile()
	fuel := make([]int, len(crabs))

	for i, c := range crabs {
		for j, o := range crabs {
			if i != j {
				f := int(math.Abs(float64(c - o)))
				fuel[i] += f
			}
		}
	}

	min := -1
	for _, f := range fuel {
		if f < min || min == -1 {
			min = f
		}
	}

	fmt.Printf("The answer is: %d\n", min)
}

func (d Day7) Part2() {
	crabs, max := readFile()
	fuel := make([]int, len(crabs))
	cache := make([]int, max+1)

	for i, c := range crabs {
		for j, o := range crabs {
			if i != j {
				f := int(math.Abs(float64(c - o)))
				if cache[f] == 0 {
					var sum int
					for k := 0; k < f; k++ {
						sum += k + 1
					}
					cache[f] = sum
				}

				fuel[i] += cache[f]
			}
		}
	}

	min := -1
	for _, f := range fuel {
		if f < min || min == -1 {
			min = f
		}
	}

	fmt.Printf("The answer is: %d\n", min)
}

func readFile() ([]int, int) {
	file, err := os.Open("src/day7/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	data := make([]int, 0)
	scanner := bufio.NewScanner(file)
	var max int
	for scanner.Scan() {
		l := scanner.Text()
		s := strings.Split(l, ",")
		for _, a := range s {
			crabPosition, _ := strconv.Atoi(a)
			data = append(data, crabPosition)
			if crabPosition > max {
				max = crabPosition
			}
		}
	}

	return data, max
}
