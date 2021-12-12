package day11

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type (
	Day11 struct{}

	Octopus struct {
		energyLevel int
		flashed     bool
		i           int
		j           int
	}

	Size struct {
		width  int
		height int
	}

	OctopusesMap [][]Octopus
)

func (d Day11) Part1() {
	octopusesMap, size := readFile()
	var flashes, steps int
	for steps < 100 {
		octopusesMap.increaseEnergy()
		willFlash := octopusesMap.getOctopusesThatWillFlash()
		for len(willFlash) > 0 {
			for _, o := range willFlash {
				if !o.flashed {
					o.flashed = true
					flashes++
					neighbours := o.getNeighbours(size, octopusesMap)
					for _, n := range neighbours {
						n.energyLevel += 1
					}
				}
			}
			willFlash = octopusesMap.getOctopusesThatWillFlash()
		}

		octopusesMap.resetFlashed()
		steps++
	}

	fmt.Printf("The answer is: %d\n", flashes)
}

func (d Day11) Part2() {
	octopusesMap, size := readFile()
	var flashes, steps int
	for !octopusesMap.willFlashTogether() {
		octopusesMap.increaseEnergy()
		willFlash := octopusesMap.getOctopusesThatWillFlash()
		for len(willFlash) > 0 {
			for _, o := range willFlash {
				if !o.flashed {
					o.flashed = true
					flashes++
					neighbours := o.getNeighbours(size, octopusesMap)
					for _, n := range neighbours {
						n.energyLevel += 1
					}
				}
			}
			willFlash = octopusesMap.getOctopusesThatWillFlash()
		}

		octopusesMap.resetFlashed()
		steps++
	}

	fmt.Printf("The answer is: %d\n", steps)
}

func readFile() (*OctopusesMap, Size) {
	file, err := os.Open("src/day11/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var octopusesMap OctopusesMap
	var size Size
	var i int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		l := scanner.Text()
		line := []rune(l)
		var octopusesLine []Octopus
		for j, s := range line {
			value, _ := strconv.Atoi(string(s))
			octopusesLine = append(octopusesLine, Octopus{energyLevel: value, i: i, j: j})
		}
		size.width = len(line)
		i++
		octopusesMap = append(octopusesMap, octopusesLine)
	}
	size.height = len(octopusesMap)

	return &octopusesMap, size
}

func (om *OctopusesMap) increaseEnergy() {
	omPtr := *om
	for i := range omPtr {
		for j := range omPtr[i] {
			omPtr[i][j].energyLevel += 1
		}
	}
}

func (om *OctopusesMap) getOctopusesThatWillFlash() []*Octopus {
	var willFlash []*Octopus
	omPtr := *om
	for i := range omPtr {
		for j := range omPtr[i] {
			if omPtr[i][j].energyLevel > 9 && !omPtr[i][j].flashed {
				willFlash = append(willFlash, &omPtr[i][j])
			}
		}
	}

	return willFlash
}

func (om *OctopusesMap) resetFlashed() {
	omPtr := *om
	for i := range omPtr {
		for j := range omPtr[i] {
			if omPtr[i][j].flashed {
				omPtr[i][j].energyLevel = 0
			}
			omPtr[i][j].flashed = false
		}
	}
}

func (om *OctopusesMap) willFlashTogether() bool {
	omPtr := *om
	for i := range omPtr {
		for j := range omPtr[i] {
			if omPtr[i][j].energyLevel != 0 {
				return false
			}
		}
	}

	return true
}

func (om *OctopusesMap) Print() {
	for _, l := range *om {
		for _, o := range l {
			if o.energyLevel > 9 {
				fmt.Printf("X")
			} else {
				fmt.Printf("%d", o.energyLevel)
			}
		}
		fmt.Println()
	}
}

func (o *Octopus) getNeighbours(size Size, octopusesMap *OctopusesMap) []*Octopus {
	var neighbours []*Octopus
	i := o.i
	j := o.j
	om := *octopusesMap
	top := i - 1
	if top >= 0 {
		neighbours = append(neighbours, &om[top][j])
	}
	right := j + 1
	if right < size.width {
		neighbours = append(neighbours, &om[i][right])
	}
	bottom := i + 1
	if bottom < size.height {
		neighbours = append(neighbours, &om[bottom][j])
	}
	left := j - 1
	if left >= 0 {
		neighbours = append(neighbours, &om[i][left])
	}
	if top >= 0 && left >= 0 {
		neighbours = append(neighbours, &om[top][left])
	}
	if top >= 0 && right < size.width {
		neighbours = append(neighbours, &om[top][right])
	}
	if bottom < size.height && right < size.width {
		neighbours = append(neighbours, &om[bottom][right])
	}
	if bottom < size.height && left >= 0 {
		neighbours = append(neighbours, &om[bottom][left])
	}

	return neighbours
}
