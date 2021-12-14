package day14

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strings"
)

type (
	Day14 struct{}

	Pairs map[string]string

	Polymer string

	Polymers []string

	Step struct {
		Pairs map[string]int
		Count map[string]int
	}
)

func (d Day14) Part1() {
	polymer, pairs := readFile()

	steps := 0
	for steps < 10 {
		pr := []rune(polymer)
		l := len(pr)
		var i int
		var newPolymer string
		for i < l {
			if i+1 < l {
				pair := string(pr[i : i+2])
				newPolymer += string(pr[i])
				value, ok := pairs[pair]
				if ok {
					newPolymer += value
				}
			} else {
				newPolymer += string(pr[i])
			}
			i++
		}

		polymer = Polymer(newPolymer)
		steps++
	}

	a := polymer.getAnswer()

	fmt.Printf("The answer is: %d\n", a)
}

func (d Day14) Part2() {
	polymer, pairs := readFile()

	s := createStep(string(polymer))

	steps := 0
	for steps < 40 {
		s = nextStep(s, pairs)
		steps++
	}

	min := math.MaxInt
	max := 0
	for _, v := range s.Count {
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}

	a := max - min + 1
	fmt.Printf("The answer is: %d\n", a)
}

func readFile() (Polymer, Pairs) {
	file, err := os.Open("src/day14/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var polymer Polymer
	pairs := make(Pairs)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		l := scanner.Text()
		if len(polymer) == 0 {
			polymer = Polymer(l)
		} else if l != "" {
			s := strings.Split(l, " -> ")
			pair := s[0]
			element := s[1]
			pairs[pair] = string(element[0])
		}
	}

	return polymer, pairs
}

func (p Polymer) getAnswer() int {
	data := make(map[string]int)
	runes := []rune(p)
	for _, r := range runes {
		data[string(r)] += 1
	}
	lessCommon := math.MaxInt
	var mostCommon int
	for _, v := range data {
		if v < lessCommon {
			lessCommon = v
		}

		if v > mostCommon {
			mostCommon = v
		}
	}

	return mostCommon - lessCommon
}

func nextStep(step *Step, pairs Pairs) *Step {
	newStep := &Step{
		Pairs: map[string]int{},
		Count: map[string]int{},
	}

	for k, v := range step.Count {
		newStep.Count[k] = v
	}

	for p, count := range step.Pairs {
		t := pairs[p]
		r := []rune(p)
		pair1 := string(r[0]) + t
		pair2 := t + string(r[1])
		newStep.Pairs[pair1] += count
		newStep.Pairs[pair2] += count
		newStep.Count[t] += count
	}

	return newStep
}

func createStep(polymer string) *Step {
	s := &Step{
		Pairs: map[string]int{},
		Count: map[string]int{},
	}

	for i := 0; i < len(polymer)-1; i++ {
		pair := polymer[i : i+2]
		s.Pairs[pair]++
		s.Count[string(polymer[i])]++
	}

	return s
}
