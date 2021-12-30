package day24

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type (
	Day24 struct{}
)

// The day24.xlsx file contains all the steps to get the answers.

func (d Day24) Part1() {
	data := readFile()

	// To validate theories from the spreadsheet
	/*
		instructions := parseInstructions(data[18*3 : 18*11])
		inputs := genInputs(nil, 8)

		for {
			input, ok := <-inputs
			if !ok {
				break
			}

			st := NewState(input)
			for _, i := range instructions {
				i.Exec(st)
			}

			if st.Regs[RegZ] == 0 {
				fmt.Printf("%v %+v\n", input, st)
			}
		}
	*/

	// Checking the answer obtaining the input from the spreadsheet
	instructions := parseInstructions(data)
	in := []int{9, 9, 5, 9, 8, 9, 6, 3, 9, 9, 9, 9, 7, 1}
	st := NewState(in)
	for _, i := range instructions {
		i.Exec(st)
	}

	// To test z == 0
	//fmt.Printf("%v %+v\n", in, st)

	answer := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(in)), ""), "[]")
	fmt.Printf("The answer is: %s\n", answer)
}

func (d Day24) Part2() {
	data := readFile()

	// Checking the answer obtaining the input from the spreadsheet
	instructions := parseInstructions(data)
	in := []int{9, 3, 1, 5, 1, 4, 1, 1, 7, 1, 1, 2, 1, 1}
	st := NewState(in)
	for _, i := range instructions {
		i.Exec(st)
	}

	// To test z == 0
	//fmt.Printf("%v %+v\n", in, st)

	answer := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(in)), ""), "[]")
	fmt.Printf("The answer is: %s\n", answer)
}

func readFile() []string {
	file, err := os.Open("src/day24/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	data := make([]string, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		data = append(data, scanner.Text())
	}

	return data
}

func NewState(input []int) *State {
	i := make([]int, len(input))
	copy(i, input)
	return &State{
		Regs:  []int{0, 0, 0, 0},
		input: i,
	}
}

func genInputs(prefix []int, length int) chan []int {
	c := make(chan []int)

	var gen func(in []int)
	gen = func(in []int) {
		for i := 9; i >= 1; i-- {
			//duplicate input slice
			next := make([]int, len(in)+1)
			copy(next, in)
			next[len(in)] = i

			// if long enough
			if len(next) == length {
				c <- next
			} else {
				// otherwise, recursively generate the next turn
				gen(next)
			}
		}
		// close channel if this is the top level to indicate completion
		if len(in) == len(prefix) {
			close(c)
		}
	}

	go gen(prefix)
	return c
}
