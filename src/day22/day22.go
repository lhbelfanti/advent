package day22

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type (
	Day22 struct{}
)

func (d Day22) Part1() {
	data := readFile()
	answer := solve(1, data)

	fmt.Printf("The answer is: %d\n", answer)
}

func (d Day22) Part2() {
	data := readFile()
	answer := solve(2, data)

	fmt.Printf("The answer is: %d\n", answer)
}

func readFile() [][]int {
	file, err := os.Open("src/day22/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	data := make([][]int, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		row := make([]int, 0)
		t := scanner.Text() // on x=-3..43,y=-28..22,z=-6..38
		s1 := strings.Split(t, " ")
		s := strings.Replace(strings.Replace(strings.Replace(s1[1], "x=", " ", -1), ",y=", " ", -1), ",z=", " ", -1)
		s = strings.Replace(s, "..", " ", -1)
		s2 := strings.Fields(s)
		var a, b, c, d, e, f, g int
		if s1[0] == "on" {
			a = 1
		} else {
			a = 0
		}
		b, _ = strconv.Atoi(s2[0])
		c, _ = strconv.Atoi(s2[1])
		d, _ = strconv.Atoi(s2[2])
		e, _ = strconv.Atoi(s2[3])
		f, _ = strconv.Atoi(s2[4])
		g, _ = strconv.Atoi(s2[5])

		row = append(row, a, b, c, d, e, f, g)
		data = append(data, row)
	}

	return data
}

func solve(part int, data [][]int) int {
	regions := make([][]int, 0)
	nonOverlappingRegions := []int{0, -100000, 100000, -100000, 100000, -100000, 100000}
	regions = append(regions, nonOverlappingRegions)

	var counter int
	for _, r := range data {
		state, x1, x2, y1, y2, z1, z2 := r[0], r[1], r[2], r[3], r[4], r[5], r[6]

		if part == 1 {
			if x1 < -50 {
				x1 = -50
			}
			if y1 < -50 {
				y1 = -50
			}
			if z1 < -50 {
				z1 = -50
			}
			if x2 > 50 {
				x2 = 50
			}
			if y2 > 50 {
				y2 = 50
			}
			if z2 > 50 {
				z2 = 50
			}
			if x1 > x2 || y1 > y2 || z1 > z2 {
				continue
			}
		}

		newRegions := make([][]int, 0)
		for _, region := range regions {
			if !overlap3D(region, r) {
				newRegions = append(newRegions, region)
				continue
			}
			if region[1] < x1 {
				nr := []int{region[0], region[1], x1 - 1, region[3], region[4], region[5], region[6]}
				newRegions = append(newRegions, nr)
				region[1] = x1
			}
			if region[2] > x2 {
				nr := []int{region[0], x2 + 1, region[2], region[3], region[4], region[5], region[6]}
				newRegions = append(newRegions, nr)
				region[2] = x2
			}
			if region[3] < y1 {
				nr := []int{region[0], region[1], region[2], region[3], y1 - 1, region[5], region[6]}
				newRegions = append(newRegions, nr)
				region[3] = y1
			}
			if region[4] > y2 {
				nr := []int{region[0], region[1], region[2], y2 + 1, region[4], region[5], region[6]}
				newRegions = append(newRegions, nr)
				region[4] = y2
			}
			if region[5] < z1 {
				nr := []int{region[0], region[1], region[2], region[3], region[4], region[5], z1 - 1}
				newRegions = append(newRegions, nr)
				region[5] = z1
			}
			if region[6] > z2 {
				nr := []int{region[0], region[1], region[2], region[3], region[4], z2 + 1, region[6]}
				newRegions = append(newRegions, nr)
				region[6] = z2
			}
			region[0] = state
			newRegions = append(newRegions, region)
		}
		regions = newRegions
		checkRegions(regions)
		counter++
	}

	var answer int
	for _, r := range regions {
		state, x1, x2, y1, y2, z1, z2 := r[0], r[1], r[2], r[3], r[4], r[5], r[6]
		if state == 1 {
			answer += (x2 - x1 + 1) * (y2 - y1 + 1) * (z2 - z1 + 1)
		}
	}

	return answer
}

func overlap1D(box1Min, box1Max, box2Min, box2Max int) bool {
	return box1Max >= box2Min && box2Max >= box1Min
}

func overlap3D(coords1, coords2 []int) bool {
	return overlap1D(coords1[1], coords1[2], coords2[1], coords2[2]) &&
		overlap1D(coords1[3], coords1[4], coords2[3], coords2[4]) &&
		overlap1D(coords1[5], coords1[6], coords2[5], coords2[6])
}

func checkRegions(regions [][]int) {
	for _, r := range regions {
		x1, x2, y1, y2, z1, z2 := r[1], r[2], r[3], r[4], r[5], r[6]
		if x2 < x1 || y2 < y1 || z2 < z1 {
			fmt.Printf("Invalid region: %d, %d, %d, %d, %d, %d", x1, x2, y1, y2, z1, z2)
		}
	}
}
