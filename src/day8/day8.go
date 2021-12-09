package day8

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type (
	Day8 struct {
		Entries []Entry
	}

	Entry struct {
		input  map[int][]Segment
		output []string
	}

	Mapper map[string]string

	Segment string

	BySegmentLength []Segment

	RuneSlice []rune
)

func (a BySegmentLength) Len() int           { return len(a) }
func (a BySegmentLength) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a BySegmentLength) Less(i, j int) bool { return len(a[i]) < len(a[j]) }

func (p RuneSlice) Len() int           { return len(p) }
func (p RuneSlice) Less(i, j int) bool { return p[i] < p[j] }
func (p RuneSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func (d Day8) Part1() {
	d.Entries = readFile()

	var counter int
	for _, e := range d.Entries {
		for _, out := range e.output {
			n := len(out)
			if n == 2 || n == 3 || n == 4 || n == 7 { // 1, 7, 4 or 8
				counter++
			}
		}
	}

	fmt.Printf("The answer is: %d\n", counter)
}

func (d Day8) Part2() {
	d.Entries = readFile()

	digits := make(map[int]Segment, 10)

	var sum int
	for _, e := range d.Entries {
		var four, two string
		for i := 2; i < 8; i++ {
			values := e.input[i]
			switch i {
			case 2: // 1
				digits[1] = values[0]
			case 3: // 7
				digits[7] = values[0]
			case 4: // 4
				v := values[0]
				digits[4] = v
				s := []rune(digits[1])
				for _, l := range v {
					if !strings.Contains(string(s), string(l)) {
						four += string(l)
					}
				}
			case 5: // 2, 3 or 5
				for _, v := range values {
					d1 := []rune(digits[1])
					d4 := []rune(four)
					if strings.Contains(string(v), string(d1[0])) && strings.Contains(string(v), string(d1[1])) {
						digits[3] = v
					} else if strings.Contains(string(v), string(d4[0])) && strings.Contains(string(v), string(d4[1])) {
						digits[5] = v
					} else {
						digits[2] = v
						s := []rune(digits[1])
						for _, l := range v {
							if !strings.Contains(string(s), string(l)) {
								two += string(l)
							}
						}
					}
				}
			case 6: // 0, 6 or 9
				for _, v := range values {
					d4 := []rune(four)
					if strings.Contains(string(v), string(d4[0])) && strings.Contains(string(v), string(d4[1])) {
						d2 := []rune(two)
						var c int
						for _, l := range d2 {
							if strings.Contains(string(v), string(l)) {
								c++
							}
						}

						if c == len(values) {
							digits[9] = v
						} else {
							digits[6] = v
						}
					} else {
						digits[0] = v
					}
				}
			case 7: // 8
				digits[8] = values[0]
			}
		}

		var numberStr string
		for _, out := range e.output {
			numberStr += transformDigitToNumber(digits, out)
		}

		number, _ := strconv.Atoi(numberStr)
		sum += number
	}

	fmt.Printf("The answer is: %d\n", sum)
}

func readFile() []Entry {
	file, err := os.Open("src/day8/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	data := make([]Entry, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		l := scanner.Text()
		s := strings.Split(l, "|")
		inputValues := strings.Fields(s[0])
		input := make([]Segment, 10)
		for i, inp := range inputValues {
			ordered := []rune(inp)
			sort.Sort(RuneSlice(ordered))
			input[i] = Segment(ordered)
		}
		orderedByLength := make(map[int][]Segment, 10)
		sort.Sort(BySegmentLength(input))

		for _, inp := range input {
			orderedByLength[len(inp)] = append(orderedByLength[len(inp)], inp)
		}

		output := strings.Fields(s[1])
		data = append(data, Entry{
			input:  orderedByLength,
			output: output,
		})
	}

	return data
}

func transformDigitToNumber(digits map[int]Segment, outputValue string) string {
	for key, value := range digits {
		orderedValue := []rune(value)
		sort.Sort(RuneSlice(orderedValue))

		orderedOutput := []rune(outputValue)
		sort.Sort(RuneSlice(orderedOutput))

		if string(orderedValue) == string(orderedOutput) {
			return strconv.Itoa(key)
		}
	}

	return "error"
}
