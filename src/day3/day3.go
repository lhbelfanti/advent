package day3

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Day3 struct{}

func (d Day3) Part1() {
	file, err := os.Open("src/day3/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	consumption := make(map[int][]int)

	for scanner.Scan() {
		n := scanner.Text()
		letters := strings.Split(n, "")
		for i, v := range letters {
			b, _ := strconv.Atoi(v)
			consumption[i] = append(consumption[i], b)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	var gamaRate, epsilonRate string

	length := len(consumption)
	for k := 0; k < length; k++ {
		v := consumption[k]
		gamaRate += countFreq(v, true)
		epsilonRate += countFreq(v, false)
	}

	gamaBinary, _ := strconv.ParseInt(gamaRate, 2, 64)
	epsilonBinary, _ := strconv.ParseInt(epsilonRate, 2, 64)

	fmt.Printf("The answer is: %d\n", gamaBinary*epsilonBinary)
}

func countFreq(arr []int, mostFreq bool) string {
	m := map[int]int{}
	var freq string
	for _, a := range arr {
		m[a]++
	} // 198

	if mostFreq {
		if m[1] >= m[0] {
			freq = "1"
		} else {
			freq = "0"
		}
	} else {
		if m[0] <= m[1] {
			freq = "0"
		} else {
			freq = "1"
		}
	}

	return freq
}

func (d Day3) Part2() {
	file, err := os.Open("src/day3/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	filteredOGR := make([]string, 0)
	filteredCS := make([]string, 0)

	for scanner.Scan() {
		n := scanner.Text()
		filteredOGR = append(filteredOGR, n)
		filteredCS = append(filteredCS, n)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	length := len(strings.Split(filteredOGR[0], ""))
	for k := 0; k < length; k++ {
		if len(filteredOGR) != 1 {
			filteredOGR = filterValues(filteredOGR, k, countFreqBySlice(filteredOGR, k, true))
		}
		if len(filteredCS) != 1 {
			filteredCS = filterValues(filteredCS, k, countFreqBySlice(filteredCS, k, false))
		}
	}

	oxygenGeneratorRating, _ := strconv.ParseInt(filteredOGR[0], 2, 64)
	co2Scrubber, _ := strconv.ParseInt(filteredCS[0], 2, 64)

	fmt.Printf("The answer is: %d\n", oxygenGeneratorRating*co2Scrubber)
}

func countFreqBySlice(values []string, index int, mostFreq bool) string {
	count := make(map[int][]int)
	for _, s := range values {
		letters := strings.Split(s, "")
		for i, v := range letters {
			b, _ := strconv.Atoi(v)
			count[i] = append(count[i], b)
		}
	}

	return countFreq(count[index], mostFreq)
}

func filterValues(values []string, index int, bit string) []string {
	var filteredValues []string

	for _, v := range values {
		letters := strings.Split(v, "")
		if letters[index] == bit {
			filteredValues = append(filteredValues, v)
		}
	}

	return filteredValues
}
