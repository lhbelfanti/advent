package day5

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type (
	Day5 struct {
		lines   []Line
		diagram Diagram
	}
	Line struct {
		p1 Point
		p2 Point
	}

	Point struct {
		x int
		y int
	}

	Diagram [][]int
)

func (d Day5) Part1() {
	data := d.readFile()

	for _, l := range data.lines {
		if l.isHorizontalOrVerticalLine() {
			points := l.getHorizontalAndVerticalCoverPoints()
			for _, p := range points {
				data.diagram[p.x][p.y] += 1
			}
		}
	}

	overlaps := data.diagram.countOverlap()

	fmt.Printf("The answer is: %d\n", overlaps)
}

func (d Day5) Part2() {
	data := d.readFile()

	for _, l := range data.lines {
		if l.isHorizontalOrVerticalLine() {
			points := l.getHorizontalAndVerticalCoverPoints()
			for _, p := range points {
				data.diagram[p.x][p.y] += 1
			}
		} else if l.isDiagonalLine() {
			points := l.getDiagonalCoverPoints()
			for _, p := range points {
				data.diagram[p.x][p.y] += 1
			}
		}
	}

	overlaps := data.diagram.countOverlap()

	fmt.Printf("The answer is: %d\n", overlaps)
}

func (d Day5) readFile() Day5 {
	file, err := os.Open("src/day5/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	d.lines = make([]Line, 0)

	var lengthX, lengthY int

	for scanner.Scan() {
		l := scanner.Text()
		s := strings.Split(l, " -> ")
		p1 := getPoint(s[0])
		p2 := getPoint(s[1])
		lengthX = getXLength(p1, p2, lengthX)
		lengthY = getYLength(p1, p2, lengthY)
		line := Line{p1: p1, p2: p2}
		d.lines = append(d.lines, line)
	}

	d.diagram = make([][]int, lengthX+1)
	for i := range d.diagram {
		d.diagram[i] = make([]int, lengthY+1)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return d
}

func getPoint(s string) Point {
	ps := strings.Split(s, ",")
	x, _ := strconv.Atoi(ps[0])
	y, _ := strconv.Atoi(ps[1])
	return Point{x: x, y: y}
}

func getXLength(p1, p2 Point, xLength int) int {
	if p1.x > xLength {
		xLength = p1.x
	}
	if p2.x > xLength {
		xLength = p2.x
	}

	return xLength
}

func getYLength(p1, p2 Point, yLength int) int {
	if p1.y > yLength {
		yLength = p1.y
	}
	if p2.y > yLength {
		yLength = p2.y
	}

	return yLength
}

func (l Line) getOrderedPoints() []Point {
	if l.p1.x == l.p2.x { // Horizontal
		if l.p1.y < l.p2.y {
			return []Point{l.p1, l.p2}
		} else {
			return []Point{l.p2, l.p1}
		}
	} else if l.p1.y == l.p2.y { // Vertical
		if l.p1.x < l.p2.x {
			return []Point{l.p1, l.p2}
		} else {
			return []Point{l.p2, l.p1}
		}
	} else { // Diagonal
		if l.p1.x < l.p2.x {
			return []Point{l.p1, l.p2}
		} else {
			return []Point{l.p2, l.p1}
		}
	}
}

func (l Line) isHorizontalOrVerticalLine() bool {
	return l.p1.x == l.p2.x || l.p1.y == l.p2.y
}

func (l Line) isDiagonalLine() bool {
	var a, b int
	if l.p1.x > l.p2.x {
		a = l.p1.x - l.p2.x
	} else {
		a = l.p2.x - l.p1.x
	}

	if l.p1.y > l.p2.y {
		b = l.p1.y - l.p2.y
	} else {
		b = l.p2.y - l.p1.y
	}
	return a == b
}

func (l Line) getHorizontalAndVerticalCoverPoints() []Point {
	points := make([]Point, 0)
	ordered := l.getOrderedPoints()
	p1 := ordered[0]
	p2 := ordered[1]
	diffX := p2.x - p1.x
	diffY := p2.y - p1.y

	if diffX > 1 {
		for i := p1.x; i < p1.x+diffX+1; i++ {
			points = append(points, Point{x: i, y: p1.y})
		}
	} else if diffY > 1 {
		for i := p1.y; i < p1.y+diffY+1; i++ {
			points = append(points, Point{x: p1.x, y: i})
		}
	} else {
		points = append(points, p1, p2)
	}

	return points
}

func (l Line) getDiagonalCoverPoints() []Point {
	points := make([]Point, 0)
	ordered := l.getOrderedPoints()
	p1 := ordered[0]
	p2 := ordered[1]
	diffX := p2.x - p1.x
	diffY := p2.y - p1.y

	if diffX > 1 || diffY > 1 {
		for i := 0; i < diffX+1; i++ {
			if diffX != diffY {
				points = append(points, Point{x: p1.x + i, y: p1.y - i})
			} else {
				points = append(points, Point{x: p1.x + i, y: p1.y + i})
			}
		}
	} else {
		points = append(points, p1, p2)
	}

	return points
}

func (d Diagram) countOverlap() int {
	var sum int
	for i := range d {
		for j := range d[i] {
			if d[i][j] >= 2 {
				sum += 1
			}
		}
	}

	return sum
}
