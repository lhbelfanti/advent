package day19

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type (
	Day19 struct{}

	ScanResult map[Vector3]bool
)

func (d Day19) Part1() {
	scans := processData(readFile())
	_, ocean := solve(scans)

	fmt.Printf("The answer is: %d\n", len(ocean))
}

func (d Day19) Part2() {
	scans := processData(readFile())
	scannerLocs, _ := solve(scans)

	maxDist := 0
	for i := 0; i < len(scannerLocs); i++ {
		for j := 1; j < len(scannerLocs); j++ {
			dist := scannerLocs[i].Dist(scannerLocs[j])
			if dist > maxDist {
				maxDist = dist
			}
		}
	}

	fmt.Printf("The answer is: %d\n", maxDist)
}

func readFile() []string {
	file, err := os.Open("src/day19/input.txt")
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

func processData(data []string) []ScanResult {
	scans := []ScanResult{}
	var scan ScanResult
	for i := 0; i < len(data); i++ {
		if data[i] == "" {
			continue
		}

		if strings.Index(data[i], "scanner") >= 0 {
			scan = ScanResult{}
			scans = append(scans, scan)
			continue
		}

		v := Vector3{}
		fmt.Sscanf(data[i], "%d,%d,%d", &v.x, &v.y, &v.z)
		scan[v] = true
	}

	return scans
}

func (s ScanResult) String() string {
	var sb strings.Builder
	sb.WriteByte('[')
	sep := ""
	for p := range s {
		sb.WriteString(sep)
		sb.WriteString(fmt.Sprintf("%v", p))
		sep = " "
	}
	sb.WriteByte(']')
	return sb.String()
}

func solve(scans []ScanResult) ([]Vector3, ScanResult) {
	ocean := ScanResult{}
	for p := range scans[0] {
		ocean[p] = true
	}

	scannerLocs := []Vector3{{0, 0, 0}}

	scans = scans[1:]

	for len(scans) > 0 {
	outer:
		for i := len(scans) - 1; i >= 0; i-- {
			for rotID := 0; rotID < 24; rotID++ {
				offsets := map[Vector3]int{}
				for knownPoint := range ocean {
					for p := range scans[i] {
						offset := p.Rotate(rotID).Sub(knownPoint)
						offsets[offset]++
					}
				}

				for offset, count := range offsets {
					if count >= 12 {
						scanner := offset.Inverse()
						scannerLocs = append(scannerLocs, scanner)

						for p := range scans[i] {
							mappedPoint := p.Rotate(rotID).Add(scanner)
							ocean[mappedPoint] = true
						}

						// fmt.Println(rotID, scanner, count)

						scans = append(scans[:i], scans[i+1:]...)

						continue outer
					}
				}
			}
		}
	}

	return scannerLocs, ocean
}
