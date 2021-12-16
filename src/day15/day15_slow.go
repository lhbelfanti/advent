package day15

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
)

type (
	Day15 struct{}

	Data struct {
		size Size
		grid [][]int
	}

	Size struct {
		width  int
		height int
	}

	Point struct {
		x int
		y int
	}
)

// This is the first implementation of the algorithms. It works, and I used them to pass the challenge, but they are really slow.
// Took me 30 min until I got the second answer.

func (d Day15) Part1Slow() {
	data := readFile()
	a := dijkstra(data)

	fmt.Printf("The answer is: %d\n", a)
}

func (d Day15) Part2Slow() {
	data := readFile()
	a := dijkstraXL(data)

	fmt.Printf("The answer is: %d\n", a)
}

func readFile() Data {
	file, err := os.Open("src/day15/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var size Size
	data := Data{
		size: Size{},
		grid: make([][]int, 0),
	}
	fileData := make([]string, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		l := scanner.Text()
		fileData = append(fileData, l)
	}

	grid := make([][]int, len(fileData))
	for i := range grid {
		grid[i] = make([]int, len(fileData[i]))
		for j, v := range fileData[i] {
			grid[i][j] = int(v - '0')
		}
	}

	size.width = len(grid[0])
	size.height = len(grid)
	data.grid = grid
	data.size = size

	return data
}

func dijkstra(data Data) int {
	grid := data.grid
	size := data.size
	maxX := data.size.width
	maxY := data.size.height
	distance := make(map[Point]int)
	visited := make(map[Point]bool)
	for y := 0; y < maxY; y++ {
		for x := 0; x < maxX; x++ {
			if x == 0 && y == 0 {
				distance[Point{x, y}] = 0
			} else {
				distance[Point{x, y}] = math.MaxInt
			}
		}
	}

	var a int
	current := Point{0, 0}

	for {
		for _, n := range neighbours(current, size) {
			if visited[n] {
				continue
			}
			newDistance := distance[current] + grid[n.y][n.x]
			if newDistance < distance[n] {
				distance[n] = newDistance
			}
		}
		visited[current] = true

		if visited[Point{maxX - 1, maxY - 1}] {
			a = distance[Point{maxX - 1, maxY - 1}]
			break
		}

		minDist := math.MaxInt
		current = Point{maxX, maxY}
		for p, v := range distance {
			if !visited[p] && v < minDist {
				minDist = v
				current = p
			}
		}
	}

	return a
}

func dijkstraXL(data Data) int {
	grid := data.grid
	maxX := data.size.width * 5
	maxY := data.size.height * 5
	sizeXL := Size{maxX, maxY}
	distance := make(map[Point]int)
	visited := make(map[Point]bool)
	distance[Point{0, 0}] = 0

	var a int
	current := Point{0, 0}

	for {
		for _, n := range neighbours(current, sizeXL) {
			if visited[n] {
				continue
			}

			y := n.y % len(grid)
			x := n.x % len(grid[0])
			val := grid[y][x]
			val += n.y/len(grid) + n.x/len(grid[0])
			if val > 9 {
				val -= 9
			}

			newDistance := distance[current] + val
			if _, ok := distance[n]; !ok {
				distance[n] = newDistance
			} else if newDistance < distance[n] {
				distance[n] = newDistance
			}
		}
		visited[current] = true

		if visited[Point{maxX - 1, maxY - 1}] {
			a = distance[Point{maxX - 1, maxY - 1}]
			break
		}

		minDist := math.MaxInt
		current = Point{maxX, maxY}
		for p, v := range distance {
			if !visited[p] && v < minDist {
				minDist = v
				current = p
			}
		}
	}

	return a
}

func neighbours(p Point, size Size) []Point {
	points := make([]Point, 0)
	maxX := size.width
	maxY := size.height

	if p.x > 0 {
		points = append(points, Point{p.x - 1, p.y})
	}

	if p.x < maxX-1 {
		points = append(points, Point{p.x + 1, p.y})
	}

	if p.y > 0 {
		points = append(points, Point{p.x, p.y - 1})
	}

	if p.y < maxY-1 {
		points = append(points, Point{p.x, p.y + 1})
	}

	return points
}
