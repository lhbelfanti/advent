package day6

import (
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

	for scanner.Scan() {

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("The answer is: %d\n", 0)
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
