package day9

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

type (
	Day9 struct{}

	HeightMap [][]int

	Size struct {
		width  int
		height int
	}

	Element struct {
		i      int
		j      int
		number int
	}
)

func (d Day9) Part1() {
	data, size := readFile()
	lowerPoints := getLowerPoints(data, size)
	var sum int
	for _, rl := range lowerPoints {
		sum += rl.number + 1
	}

	fmt.Printf("The answer is: %d\n", sum)
}

func (d Day9) Part2() {
	data, size := readFile()
	lowerPoints := getLowerPoints(data, size)
	basins := make(map[int][]Element, len(lowerPoints))
	for index, lp := range lowerPoints {
		basin := make([]Element, 0)
		neighbours := data.getNeighbours(lp.i, lp.j, size, true)
		for len(neighbours) > 0 {
			n := neighbours[0]
			basin = append(basin, n)
			neighbours = neighbours[1:]
			ne := data.getNeighbours(n.i, n.j, size, true)
			for _, a := range ne {
				check1 := isInSlice(neighbours, a)
				check2 := isInSlice(basin, a)

				if !check1 && !check2 {
					neighbours = append(neighbours, a)
				}
			}
		}

		basins[index] = basin
	}

	var lengths []int
	for _, v := range basins {
		lengths = append(lengths, len(v))
	}

	sort.Ints(lengths)
	last3 := lengths[len(lengths)-3:]
	multi := last3[0] * last3[1] * last3[2]

	fmt.Printf("The answer is: %d\n", multi)
}

func readFile() (HeightMap, Size) {
	file, err := os.Open("src/day9/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	data := make(HeightMap, 0)
	size := Size{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		l := scanner.Text()
		r := []rune(l)
		numbers := make([]int, 0)
		size.width = len(r)
		for _, v := range r {
			intValue, _ := strconv.Atoi(string(v))
			numbers = append(numbers, intValue)
		}
		data = append(data, numbers)
	}

	size.height = len(data)

	return data, size
}

func getLowerPoints(data HeightMap, size Size) []Element {
	lowerPoints := make([]Element, 0)
	for i := range data {
		for j := range data[i] {
			value := data[i][j]
			neighbours := data.getNeighbours(i, j, size, false)
			isLowPoint := true
			for _, n := range neighbours {
				if n.number <= value {
					isLowPoint = false
				}
			}

			if isLowPoint {
				lp := Element{i: i, j: j, number: value}
				lowerPoints = append(lowerPoints, lp)
			}
		}
	}

	return lowerPoints
}

func (hm HeightMap) getNeighbours(i, j int, size Size, filterNines bool) []Element {
	neighbours := make([]Element, 0)
	top := i - 1
	if top >= 0 {
		num := hm[top][j]
		if !(filterNines && num == 9) {
			neighbours = append(neighbours, Element{i: top, j: j, number: num})
		}
	}
	right := j + 1
	if right < size.width {
		num := hm[i][right]
		if !(filterNines && num == 9) {
			neighbours = append(neighbours, Element{i: i, j: right, number: num})
		}
	}
	bottom := i + 1
	if bottom < size.height {
		num := hm[bottom][j]
		if !(filterNines && num == 9) {
			neighbours = append(neighbours, Element{i: bottom, j: j, number: num})
		}
	}
	left := j - 1
	if left >= 0 {
		num := hm[i][left]
		if !(filterNines && num == 9) {
			neighbours = append(neighbours, Element{i: i, j: left, number: num})
		}
	}

	return neighbours
}

func (e Element) Equals(el Element) bool {
	return e.i == el.i && e.j == el.j
}

func isInSlice(slice []Element, e Element) bool {
	inSlice := false
	for _, b := range slice {
		if e.Equals(b) {
			inSlice = true
		}
	}

	return inSlice
}
