package day6

import (
	"bufio"
	"fmt"
	"log"
	"slices"
	"strings"

	"advent2024/src/reader"
)

type Day6 struct{}

func (d Day6) Part1() {
	file, scanner := reader.Read("src/day6/input.txt")
	defer file.Close()

	maze, guardY, guardX := scan(scanner)

	hasGuardLeftTheMaze := false
	for !hasGuardLeftTheMaze {
		//printMaze(maze)
		guardDirection := maze[guardY][guardX]
		switch guardDirection {
		case "^":
			if guardY == 0 {
				hasGuardLeftTheMaze = true
				maze[guardY][guardX] = "X"
				break
			}

			if maze[guardY-1][guardX] == "#" {
				maze[guardY][guardX] = ">"
			} else {
				maze[guardY-1][guardX], maze[guardY][guardX] = guardDirection, "X"
				guardY -= 1
			}
		case ">":
			if guardX == len(maze)-1 {
				hasGuardLeftTheMaze = true
				maze[guardY][guardX] = "X"
				break
			}

			if maze[guardY][guardX+1] == "#" {
				maze[guardY][guardX] = "v"
			} else {
				maze[guardY][guardX+1], maze[guardY][guardX] = guardDirection, "X"
				guardX += 1
			}
		case "v":
			if guardY == len(maze)-1 {
				hasGuardLeftTheMaze = true
				maze[guardY][guardX] = "X"
				break
			}

			if maze[guardY+1][guardX] == "#" {
				maze[guardY][guardX] = "<"
			} else {
				maze[guardY+1][guardX], maze[guardY][guardX] = guardDirection, "X"
				guardY += 1
			}
		case "<":
			if guardX == 0 {
				hasGuardLeftTheMaze = true
				maze[guardY][guardX] = "X"
				break
			}

			if maze[guardY][guardX-1] == "#" {
				maze[guardY][guardX] = "^"
			} else {
				maze[guardY][guardX-1], maze[guardY][guardX] = guardDirection, "X"
				guardX -= 1
			}
		}
	}

	var sum int
	for y := range maze {
		for x := range maze[y] {
			if maze[y][x] == "X" {
				sum++
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("The answer is: %d\n", sum)
}

func (d Day6) Part2() {
	file, scanner := reader.Read("src/day6/input.txt")
	defer file.Close()

	scan(scanner)

	var sum int

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("The answer is: %d\n", sum)
}

func scan(scanner *bufio.Scanner) ([][]string, int, int) {
	var i, guardY, guardX int
	maze := make([][]string, 0, 130)

	for scanner.Scan() {
		tiles := strings.Split(scanner.Text(), "")
		idx := slices.Index(tiles, "^")
		if idx != -1 {
			guardY = i
			guardX = idx
		}

		maze = append(maze, tiles)
		i++
	}

	return maze, guardY, guardX
}

func printMaze(maze [][]string) {
	fmt.Println()
	for y := range maze {
		for x := range maze[y] {
			fmt.Print(maze[y][x])
		}
		fmt.Println()
	}
	fmt.Println()
}

// --------------------------------------------------------------------------------------------------------------------

type (
	Point struct {
		X, Y int
	}

	Rectangle struct {
		TopLeft      *Point
		TopRight     *Point
		BottomLeft   *Point
		BottomRight  *Point
		MissingSide  *Point
		SideToRemove RectangleSide
		SidesCounter int
	}

	RectangleSide int
)

const (
	TopLeft RectangleSide = iota
	TopRight
	BottomLeft
	BottomRight
)

// Part2FirstAttempt Not working because not all the rectangles are calculated, only the ones where the guard walks
func (d Day6) Part2FirstAttempt() {
	file, scanner := reader.Read("src/day6/input.txt")
	defer file.Close()

	maze, guardY, guardX := scan(scanner)

	var rectangle Rectangle
	rectangles := make([]Rectangle, 0, 100)

	hasGuardLeftTheMaze := false
	for !hasGuardLeftTheMaze {
		guardDirection := maze[guardY][guardX]
		switch guardDirection {
		case "^":
			if guardY == 0 {
				hasGuardLeftTheMaze = true
				maze[guardY][guardX] = "."
				break
			}

			if maze[guardY-1][guardX] == "#" {
				maze[guardY][guardX] = ">"
				ProcessMovement(&rectangle, TopLeft, Point{X: guardX, Y: guardY - 1}, &rectangles, maze)
			} else {
				CalculateLoop(Point{X: guardX, Y: guardY - 1}, maze, &rectangles, TopLeft)
				maze[guardY-1][guardX], maze[guardY][guardX] = guardDirection, "."
				guardY -= 1
			}
		case ">":
			if guardX == len(maze)-1 {
				hasGuardLeftTheMaze = true
				maze[guardY][guardX] = "."
				break
			}

			if maze[guardY][guardX+1] == "#" {
				maze[guardY][guardX] = "v"
				ProcessMovement(&rectangle, TopRight, Point{X: guardX + 1, Y: guardY}, &rectangles, maze)
			} else {
				CalculateLoop(Point{X: guardX + 1, Y: guardY}, maze, &rectangles, TopRight)
				maze[guardY][guardX+1], maze[guardY][guardX] = guardDirection, "."
				guardX += 1
			}
		case "v":
			if guardY == len(maze)-1 {
				hasGuardLeftTheMaze = true
				maze[guardY][guardX] = "."
				break
			}

			if maze[guardY+1][guardX] == "#" {
				maze[guardY][guardX] = "<"
				ProcessMovement(&rectangle, BottomRight, Point{X: guardX, Y: guardY + 1}, &rectangles, maze)
			} else {
				CalculateLoop(Point{X: guardX, Y: guardY + 1}, maze, &rectangles, BottomRight)
				maze[guardY+1][guardX], maze[guardY][guardX] = guardDirection, "."
				guardY += 1
			}
		case "<":
			if guardX == 0 {
				hasGuardLeftTheMaze = true
				maze[guardY][guardX] = "."
				break
			}

			if maze[guardY][guardX-1] == "#" {
				maze[guardY][guardX] = "^"
				ProcessMovement(&rectangle, BottomLeft, Point{X: guardX - 1, Y: guardY}, &rectangles, maze)
			} else {
				CalculateLoop(Point{X: guardX - 1, Y: guardY}, maze, &rectangles, BottomLeft)
				maze[guardY][guardX-1], maze[guardY][guardX] = guardDirection, "."
				guardX -= 1
			}
		}
	}

	var sum int
	for _, rect := range rectangles {
		if rect.SidesCounter == 4 {
			sum++
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("The answer is: %d\n", sum)
}

func ProcessMovement(rectangle *Rectangle, sideToAdd RectangleSide, point Point, rectangles *[]Rectangle, maze [][]string) {
	rectangle.AddSide(sideToAdd, point)
	if rectangle.SidesCounter == 3 {
		*rectangles = append(*rectangles, *rectangle)
		CalculateLoop(point, maze, rectangles, sideToAdd)
		rectangle.RemoveSide()
	}
}

func (r *Rectangle) AddSide(side RectangleSide, point Point) {
	var sideToCalculate RectangleSide

	switch side {
	case TopLeft:
		r.TopLeft = &point
		r.SidesCounter++
		sideToCalculate = TopRight
	case TopRight:
		r.TopRight = &point
		r.SidesCounter++
		sideToCalculate = BottomRight
	case BottomRight:
		r.BottomRight = &point
		r.SidesCounter++
		sideToCalculate = BottomLeft
	case BottomLeft:
		r.BottomLeft = &point
		r.SidesCounter++
		sideToCalculate = TopLeft
	}

	if r.SidesCounter == 3 {
		switch sideToCalculate {
		case TopLeft:
			r.MissingSide = &Point{X: r.BottomLeft.X + 1, Y: r.TopRight.Y - 1}
			r.SideToRemove = TopRight
		case TopRight:
			r.MissingSide = &Point{X: r.BottomRight.X + 1, Y: r.TopLeft.Y + 1}
			r.SideToRemove = BottomRight
		case BottomRight:
			r.MissingSide = &Point{X: r.TopRight.X - 1, Y: r.BottomLeft.Y + 1}
			r.SideToRemove = BottomLeft
		case BottomLeft:
			r.MissingSide = &Point{X: r.TopLeft.X - 1, Y: r.BottomRight.Y - 1}
			r.SideToRemove = TopLeft
		}
	}
}

func (r *Rectangle) RemoveSide() {
	switch r.SideToRemove {
	case TopLeft:
		r.TopLeft = nil
	case TopRight:
		r.TopRight = nil
	case BottomRight:
		r.BottomRight = nil
	case BottomLeft:
		r.BottomLeft = nil
	}

	r.SidesCounter--
}

func CalculateLoop(point Point, maze [][]string, rectangles *[]Rectangle, direction RectangleSide) {
	for i, rectangle := range *rectangles {
		if rectangle.SidesCounter == 3 {
			if rectangle.MissingSide.X == point.X && rectangle.MissingSide.Y == point.Y {
				rectangle.SidesCounter++
				(*rectangles)[i] = rectangle

				maze[point.Y][point.X] = "O"
				printMaze(maze)
				/*anyObstruction := isThereAnObstructionInTheMiddle(point, calculatedPosition, direction, maze)

				if !anyObstruction {
					rectangle.SidesCounter++
				}*/
			}
		}
	}
}

func isThereAnObstructionInTheMiddle(from, to Point, direction RectangleSide, maze [][]string) bool {
	var fromY, toY, fromX, toX int
	switch direction {
	case TopLeft: // ^
		fromY = to.Y
		toY = from.Y
		fromX = from.X + 1
		toX = to.X + 1
	case TopRight: // >
		fromX = from.X
		toX = to.X
		fromY = from.Y + 1
		toY = to.Y + 1
	case BottomRight: // v
		fromY = from.Y
		toY = to.Y
		fromX = from.X - 1
		toX = to.X + 1
	case BottomLeft: // <
		fromX = to.X
		toX = from.X
		fromY = to.Y
		toY = from.Y
	}

	for i := fromY; i < toY; i++ {
		for j := fromX; j < toX; j++ {
			if maze[i][j] == "#" {
				return true
			}
		}
	}

	maze[to.Y][to.X] = "O"
	printMaze(maze)

	return false
}
