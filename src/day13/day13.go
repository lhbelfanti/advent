package day13

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type (
	Day13 struct{}

	Paper [][]int

	Point struct {
		x int
		y int
	}

	Fold struct {
		dir    string
		number int
	}
)

func (d Day13) Part1() {
	paper, folds := readFile()
	f := folds[0]
	if f.dir == "y" {
		paper = paper.foldUp(f.number)
	} else {
		paper = paper.foldLeft(f.number)
	}
	c := paper.countElements()

	fmt.Printf("The answer is: %d\n", c)
}

func (d Day13) Part2() {
	paper, folds := readFile()
	for _, f := range folds {
		if f.dir == "y" {
			paper = paper.foldUp(f.number)
		} else {
			paper = paper.foldLeft(f.number)
		}
	}
	c := paper.countElements()

	fmt.Printf("The answer is: %d\n", c)
	fmt.Println("But the real answer is:")
	paper.PrintPaper()
}

func readFile() (Paper, []Fold) {
	file, err := os.Open("src/day13/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var width, height int
	var points []Point
	var folds []Fold
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		l := scanner.Text()
		p := strings.Split(l, ",")

		if len(p) == 2 {
			x, _ := strconv.Atoi(p[0])
			y, _ := strconv.Atoi(p[1])
			points = append(points, Point{x: x, y: y})

			if x > width {
				width = x
			}
			if y > height {
				height = y
			}
		} else {
			fold := strings.Fields(l)
			if len(fold) == 3 {
				d := strings.Split(fold[2], "=")
				dir := d[0]
				number, _ := strconv.Atoi(d[1])
				folds = append(folds, Fold{dir: dir, number: number})
			}
		}
	}

	paper := make(Paper, height+1)
	for i := range paper {
		paper[i] = make([]int, width+1)
	}
	for _, p := range points {
		paper[p.y][p.x] = 1
	}

	return paper, folds
}

func (p *Paper) foldUp(x int) Paper {
	paper := *p
	y := len(paper)
	for j := range paper {
		if j > x {
			for i := range paper[j] {
				if paper[j][i] == 0 {
					continue
				}
				var diff int
				if j-y < 0 {
					diff = y - j - 1
				} else {
					diff = j - y - 1
				}

				paper[diff][i] = paper[j][i]
				paper[j][i] = 0
			}
		}
	}

	newPaper := make(Paper, x)
	width := len(paper[0])
	for j := range newPaper {
		if newPaper[j] == nil {
			newPaper[j] = make([]int, width)
		}
		for i := range newPaper[j] {
			newPaper[j][i] = paper[j][i]
		}
	}

	return newPaper
}

func (p *Paper) foldLeft(y int) Paper {
	paper := *p
	for j := range paper {
		for i := range paper[j] {
			x := len(paper[j])
			if i > y {
				if paper[j][i] == 0 {
					continue
				}
				var diff int
				if i-x < 0 {
					diff = x - i - 1
				} else {
					diff = i - x - 1
				}

				paper[j][diff] = paper[j][i]
				paper[j][i] = 0
			}
		}
	}

	newPaper := make(Paper, len(paper))
	width := y
	for j := range newPaper {
		if newPaper[j] == nil {
			newPaper[j] = make([]int, width)
		}
		for i := range newPaper[j] {
			newPaper[j][i] = paper[j][i]
		}
	}

	return newPaper
}

func (p *Paper) countElements() int {
	var count int
	paper := *p
	for i := range paper {
		for j := range paper[i] {
			if paper[i][j] == 1 {
				count += 1
			}
		}
	}

	return count
}

func (p *Paper) PrintPaper() {
	paper := *p
	for i := range paper {
		for j := range paper[i] {
			if paper[i][j] == 1 {
				fmt.Print("█")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}
