package day17

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type (
	Day17 struct{}

	Range struct {
		min int
		max int
	}

	TargetArea struct {
		x Range
		y Range
	}

	Vector2 struct {
		x, y int
	}
)

func (d Day17) Part1() {
	data := readFile()

	absMin := int(math.Abs(float64(data.y.min)))
	a := (absMin - 1) * absMin / 2
	fmt.Printf("The answer is: %d\n", a)
}

func (d Day17) Part2() {
	data := readFile()
	tx1, tx2 := data.x.min, data.x.max
	ty1, ty2 := data.y.min, data.y.max

	absYMin := int(math.Abs(float64(ty1)))
	yRange := Range{min: ty1, max: absYMin - 1}
	sqrt := int(math.Sqrt(float64(tx1) * 2))
	xRange := Range{min: sqrt, max: tx2}

	var onTarget []Vector2
	for iy := yRange.min; iy <= yRange.max; iy++ {
		for ix := xRange.min; ix <= xRange.max; ix++ {
			x, y := 0, 0
			dx, dy := ix, iy

			// While not beyond target area
			for x <= tx2 && y >= ty1 {
				x += dx
				y += dy
				if dx > 0 {
					dx--
				}
				dy--

				// If in target area
				if x >= tx1 && x <= tx2 && y >= ty1 && y <= ty2 {
					onTarget = append(onTarget, Vector2{x, y})
					break
				}
			}
		}
	}

	fmt.Printf("The answer is: %d\n", len(onTarget))
}

func readFile() TargetArea {
	file, err := os.Open("src/day17/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var s string
	var x []string
	var y []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s = scanner.Text()
		s1 := strings.Split(s, "target area: x=")
		s2 := strings.Split(s1[1], ", y=")
		x = strings.Split(s2[0], "..")
		y = strings.Split(s2[1], "..")
	}

	xMin, _ := strconv.Atoi(x[0])
	xMax, _ := strconv.Atoi(x[1])

	yMin, _ := strconv.Atoi(y[0])
	yMax, _ := strconv.Atoi(y[1])

	return TargetArea{
		x: Range{
			min: xMin,
			max: xMax,
		},
		y: Range{
			min: yMin,
			max: yMax,
		},
	}
}
