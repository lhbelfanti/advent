package day4

import (
	"fmt"
	"log"
	"strings"

	"advent2024/src/reader"
)

type Day4 struct{}

func (d Day4) Part1() {
	file, scanner := reader.Read("src/day4/input.txt")
	defer file.Close()

	wordSearchMatrix := make([][]string, 0, 140)

	for scanner.Scan() {
		wordSearchMatrix = append(wordSearchMatrix, strings.Split(scanner.Text(), ""))
	}

	const wordToSearch string = "XMAS"
	wordToSearchLen := len(wordToSearch)

	counter := 0

	for n := 0; n < 4; n++ {
		wordSearchMatrixLen := len(wordSearchMatrix)

		for i := 0; i < wordSearchMatrixLen; i++ {
			row := wordSearchMatrix[i]

			shouldCheckDiagonal := i < (wordSearchMatrixLen - (wordToSearchLen - 1))
			rowLen := len(row) - (wordToSearchLen - 1)
			for j := 0; j < rowLen; j++ {
				if row[j] != "X" {
					continue
				}

				horizontalWord := strings.Join(row[j:j+wordToSearchLen], "")
				if horizontalWord == wordToSearch {
					counter++
				}

				if shouldCheckDiagonal {
					diagonalWord := row[j]
					for k := 1; k < wordToSearchLen; k++ {
						diagonalWord += wordSearchMatrix[i+k][j+k]
					}

					if diagonalWord == wordToSearch {
						counter++
					}
				}
			}
		}

		wordSearchMatrix = rotateWordSearchMatrix90Deg(wordSearchMatrix)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("The answer is: %d\n", counter)
}

func (d Day4) Part2() {
	file, scanner := reader.Read("src/day4/input.txt")
	defer file.Close()

	for scanner.Scan() {

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("The answer is: %d\n", 0)
}

func rotateWordSearchMatrix90Deg(wordSearchMatrix [][]string) [][]string {
	n := len(wordSearchMatrix)
	rotatedWordSearchMatrix := make([][]string, len(wordSearchMatrix))

	for i := range rotatedWordSearchMatrix {
		rotatedWordSearchMatrix[i] = make([]string, len(rotatedWordSearchMatrix))
	}

	for i := 0; i < n; i++ {
		for j := 0; j < len(wordSearchMatrix[0]); j++ {
			rotatedWordSearchMatrix[j][n-1-i] = wordSearchMatrix[i][j]
		}
	}

	return rotatedWordSearchMatrix
}
