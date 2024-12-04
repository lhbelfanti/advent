package reader

import (
	"bufio"
	"log"
	"os"
)

func Read(inputFile string) (*os.File, *bufio.Scanner) {
	file, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	return file, scanner
}
