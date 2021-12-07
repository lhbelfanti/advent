package day6

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Day6 struct{}

const DefaultAge = 6
const Part1Days = 80
const Part2Days = 256

func (d Day6) Part1() {
	shoal := readFile()

	var days int
	for days < Part1Days {
		babies := make([]int, 0)
		for i, lf := range shoal {
			lf--
			if lf < 0 {
				babies = append(babies, DefaultAge+2)
				lf = DefaultAge
			}
			shoal[i] = lf
		}
		shoal = append(shoal, babies...)
		days += 1
	}

	fmt.Printf("The answer is: %d\n", len(shoal))
}

func (d Day6) Part2() {
	data := readFile()

	shoal := make(map[int]int, 8)
	for _, lf := range data {
		shoal[lf] += 1
	}

	var days int
	for days < Part2Days {
		first := shoal[0]
		for i := 0; i < 8; i++ {
			shoal[i] = shoal[i+1]
		}
		shoal[6] += first
		shoal[8] = first
		days += 1
	}

	var sum int
	for _, v := range shoal {
		sum += v
	}

	fmt.Printf("The answer is: %d\n", sum)
}

func readFile() []int {
	file, err := os.Open("src/day6/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	data := make([]int, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		l := scanner.Text()
		s := strings.Split(l, ",")
		for _, a := range s {
			lf, _ := strconv.Atoi(a)
			data = append(data, lf)
		}
	}

	return data
}
