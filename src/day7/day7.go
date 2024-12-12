package day7

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"

	"advent2024/src/reader"
)

type (
	Day7 struct{}

	Calculation struct {
		Result int
		Values []int
	}
)

func (d Day7) Part1() {
	file, scanner := reader.Read("src/day7/input.txt")
	defer file.Close()

	calculations := scan(scanner)

	ops := []string{"*", "+"}

	possibilities := make(map[int][]string)

	for _, c := range calculations {
		numbers := len(c.Values) - 1
		if _, ok := possibilities[numbers]; !ok {
			possibilities[numbers] = make([]string, 0)
			format := "%0" + strconv.Itoa(numbers) + "b"
			p := pow(len(ops), numbers)
			for i := range p {
				binaryStr := fmt.Sprintf(fmt.Sprintf(format, i))
				operations := strings.Replace(binaryStr, "1", ops[0], -1)
				operations = strings.Replace(operations, "0", ops[1], -1)
				possibilities[numbers] = append(possibilities[numbers], operations)
			}
		}

	}

	var sum int
	for _, calculation := range calculations {
		combinations := possibilities[len(calculation.Values)-1]
		for _, combination := range combinations {
			operators := strings.Split(combination, "")
			maths := doMaths(calculation.Values, operators)
			if maths == calculation.Result {
				sum += calculation.Result
				break
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("The answer is: %d\n", sum)
}

func (d Day7) Part2() {
	file, scanner := reader.Read("src/day7/input.txt")
	defer file.Close()

	scan(scanner)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("The answer is: %d\n", 0)
}

func scan(scanner *bufio.Scanner) []Calculation {
	calculations := make([]Calculation, 0, 850)

	for scanner.Scan() {
		calculationString := strings.Split(scanner.Text(), ": ")
		result, _ := strconv.Atoi(calculationString[0])
		values := strings.Split(calculationString[1], " ")
		calculation := Calculation{Result: result, Values: make([]int, 0, len(values))}
		for _, value := range values {
			intValue, _ := strconv.Atoi(value)
			calculation.Values = append(calculation.Values, intValue)
		}

		calculations = append(calculations, calculation)
	}

	return calculations
}

func pow(x, y int) int {
	return int(math.Pow(float64(x), float64(y)))
}

func doMaths(values []int, operators []string) int {
	maths := values[0]
	for i := 1; i < len(values); i++ {
		switch operators[i-1] {
		case "*":
			maths *= values[i]
		case "+":
			maths += values[i]
		}
	}

	return maths
}
