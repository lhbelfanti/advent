package day10

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
)

type Day10 struct{}

func (d Day10) Part1() {
	data := readFile()
	var corruptedRunes []rune
	for _, line := range data {
		var stack []rune
		var corrupt bool
		for _, r := range line {
			if r == '(' || r == '[' || r == '{' || r == '<' {
				stack = append(stack, r)
			} else {
				if len(stack) == 0 {
					corrupt = true
					break
				}

				pop := stack[len(stack)-1:][0]
				stack = stack[:len(stack)-1]
				if pop == '(' {
					if r != ')' {
						corrupt = true
					}
				} else if pop == '[' {
					if r != ']' {
						corrupt = true
					}
				} else if pop == '{' {
					if r != '}' {
						corrupt = true
					}
				} else if pop == '<' {
					if r != '>' {
						corrupt = true
					}
				}
				if corrupt {
					corruptedRunes = append(corruptedRunes, r)
					break
				}
			}
		}
		if corrupt {
			continue
		}
	}

	var sum int
	for _, cr := range corruptedRunes {
		sum += getIllegalSyntaxPoints(cr)
	}

	fmt.Printf("The answer is: %d\n", sum)
}

func (d Day10) Part2() {
	data := readFile()
	var incompleteLines [][]rune
	for _, line := range data {
		var stack []rune
		var corrupt bool
		for _, r := range line {
			if r == '(' || r == '[' || r == '{' || r == '<' {
				stack = append(stack, r)
			} else {
				if len(stack) == 0 {
					corrupt = true
					break
				}

				pop := stack[len(stack)-1:][0]
				stack = stack[:len(stack)-1]
				if pop == '(' {
					if r != ')' {
						corrupt = true
					}
				} else if pop == '[' {
					if r != ']' {
						corrupt = true
					}
				} else if pop == '{' {
					if r != '}' {
						corrupt = true
					}
				} else if pop == '<' {
					if r != '>' {
						corrupt = true
					}
				}
				if corrupt {
					break
				}
			}
		}
		if corrupt {
			continue
		}

		if len(stack) > 0 {
			var missingPart []rune
			for len(stack) > 0 {
				pop := stack[len(stack)-1:][0]
				stack = stack[:len(stack)-1]
				if pop == '(' {
					missingPart = append(missingPart, ')')
				} else if pop == '[' {
					missingPart = append(missingPart, ']')
				} else if pop == '{' {
					missingPart = append(missingPart, '}')
				} else if pop == '<' {
					missingPart = append(missingPart, '>')
				}
			}
			incompleteLines = append(incompleteLines, missingPart)
		}
	}

	var total []int
	for _, line := range incompleteLines {
		var sum int
		for _, m := range line {
			sum *= 5
			sum += getAutocompleteToolsPoints(m)
		}
		total = append(total, sum)
	}

	sort.Ints(total)
	middle := len(total) / 2
	value := total[middle]

	fmt.Printf("The answer is: %d\n", value)
}

func readFile() [][]rune {
	file, err := os.Open("src/day10/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var data [][]rune
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		l := scanner.Text()
		line := []rune(l)
		data = append(data, line)
	}

	return data
}

func getIllegalSyntaxPoints(r rune) int {
	if r == ')' {
		return 3
	} else if r == ']' {
		return 57
	} else if r == '}' {
		return 1197
	} else if r == '>' {
		return 25137
	}

	return 0
}

func getAutocompleteToolsPoints(r rune) int {
	if r == ')' {
		return 1
	} else if r == ']' {
		return 2
	} else if r == '}' {
		return 3
	} else if r == '>' {
		return 4
	}

	return 0
}
