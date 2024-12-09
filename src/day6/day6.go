package day6

import (
	"bufio"
	"fmt"
	"log"
	"slices"
	"strings"

	"advent2024/src/reader"
)

type (
	Day6 struct{}

	Point struct {
		X, Y int
	}

	VisitedPoint map[string]int

	Move struct {
		dx, dy          int
		changeDirection map[bool]string // true for encountering `#`
	}
)

const (
	Top   string = "^"
	Right string = ">"
	Down  string = "v"
	Left  string = "<"
)

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

	maze, originalGuardY, originalGuardX := scan(scanner)

	visitedPointsMaze := make([][]VisitedPoint, len(maze))
	for a := range maze {
		visitedPointsMaze[a] = make([]VisitedPoint, len(maze[a]))
		for b := range maze[a] {
			visitedPointsMaze[a][b] = map[string]int{Top: 0, Right: 0, Down: 0, Left: 0}
		}
	}

	moves := map[string]Move{
		Top:   {dx: 0, dy: -1, changeDirection: map[bool]string{true: Right, false: Top}},
		Right: {dx: 1, dy: 0, changeDirection: map[bool]string{true: Down, false: Right}},
		Down:  {dx: 0, dy: 1, changeDirection: map[bool]string{true: Left, false: Down}},
		Left:  {dx: -1, dy: 0, changeDirection: map[bool]string{true: Top, false: Left}},
	}

	walkablePoints := getWalkablePoints(maze, originalGuardX, originalGuardY, moves)

	var loopCount int
	for _, walkablePoint := range walkablePoints {
		y, x := walkablePoint.Y, walkablePoint.X
		if y == originalGuardY && x == originalGuardX {
			continue
		}

		maze[y][x] = "#"

		resetVisitedPoints(visitedPointsMaze)

		guardX, guardY := originalGuardX, originalGuardY

		for {
			move := moves[maze[guardY][guardX]]

			guardNextY, guardNextX := guardY+move.dy, guardX+move.dx

			if guardNextY < 0 || guardNextY >= len(maze) || guardNextX < 0 || guardNextX >= len(maze[0]) {
				maze[guardY][guardX] = "."
				break // Guard left the maze
			}

			if maze[guardNextY][guardNextX] == "#" {
				maze[guardY][guardX] = move.changeDirection[true]
			} else {
				direction := move.changeDirection[false]
				maze[guardNextY][guardNextX], maze[guardY][guardX] = direction, "."

				visitedPointsMaze[guardY][guardX][direction]++
				if visitedPointsMaze[guardY][guardX][direction] >= 2 {
					//maze[y][x] = "O"
					//printMaze(maze)
					loopCount++
					break
				}

				guardY, guardX = guardNextY, guardNextX
			}
		}

		maze[y][x] = "."
		maze[originalGuardY][originalGuardX] = "^"
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("The answer is: %d\n", loopCount)
}

func getWalkablePoints(maze [][]string, guardX, guardY int, moves map[string]Move) []Point {
	originalGuardY, originalGuardX := guardY, guardX
	walkablePoints := make([]Point, 0)

	for {
		move := moves[maze[guardY][guardX]]

		point := Point{X: guardX, Y: guardY}
		walkablePoints = append(walkablePoints, point)

		nextY, nextX := guardY+move.dy, guardX+move.dx

		if nextY < 0 || nextY >= len(maze) || nextX < 0 || nextX >= len(maze[0]) {
			maze[guardY][guardX] = "."
			break
		}

		if maze[nextY][nextX] == "#" {
			maze[guardY][guardX] = move.changeDirection[true]
		} else {
			maze[nextY][nextX], maze[guardY][guardX] = move.changeDirection[false], "."
			guardY, guardX = nextY, nextX
		}
	}

	maze[originalGuardY][originalGuardX] = "^"

	return unique(walkablePoints)
}

func unique[T comparable](elements []T) []T {
	alreadyAdded := make(map[T]bool)
	var result []T
	for _, element := range elements {
		if _, ok := alreadyAdded[element]; !ok {
			alreadyAdded[element] = true
			result = append(result, element)
		}
	}

	return result
}

func resetVisitedPoints(visitedPoints [][]VisitedPoint) {
	for i := range visitedPoints {
		for j := range visitedPoints[i] {
			visitedPoints[i][j] = map[string]int{Top: 0, Right: 0, Down: 0, Left: 0}
		}
	}
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
