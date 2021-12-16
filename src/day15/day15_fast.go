package day15

import (
	"bufio"
	"container/heap"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
)

type (
	Grid1 struct {
		Points map[Point]GridPoint
		X      int
		Y      int
	}

	Grid2 struct {
		Points map[Point]int
		X      int
		Y      int
	}

	GridPoint struct {
		Point Point
		Value int
	}

	State struct {
		Score int
		Pos   Point
	}

	QueueItem struct {
		Point Point
		Score int
		Index int
	}

	PriorityQueue []*QueueItem
)

var North = Point{0, 1}
var South = Point{0, -1}
var East = Point{1, 0}
var West = Point{-1, 0}
var Directions = []Point{North, South, East, West}

func (d Day15) Part1Fast() {
	lines := readLines()
	grid := parseGridPart1(lines)
	a := solvePart1(grid)

	fmt.Printf("The answer is: %d\n", a)
}

func (d Day15) Part2Fast() {
	lines := readLines()
	grid := parseGridPart2(lines)
	a := solvePart2(grid)

	fmt.Printf("The answer is: %d\n", a)
}

func readLines() []string {
	result := make([]string, 0)
	file, err := os.Open("src/day15/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		result = append(result, line)
	}

	return result
}

func parseGridPart1(lines []string) Grid1 {
	r := Grid1{make(map[Point]GridPoint), len([]rune(lines[0])), len(lines)}

	for j, ln := range lines {
		for i, c := range ln {
			r.Points[Point{x: i, y: j}] = GridPoint{Point{x: i, y: j}, int(c - '0')}
		}
	}

	return r
}

func parseGridPart2(lines []string) Grid2 {
	xsize := len([]rune(lines[0]))
	ysize := len(lines)
	m := make(map[Point]int)
	r := Grid2{m, xsize * 5, ysize * 5}

	for j := 0; j < r.Y; j++ {
		for i := 0; i < r.X; i++ {
			ln := lines[j%ysize]
			c := []rune(ln)[i%xsize]
			scoreIncrement := i/xsize + j/ysize
			r.Points[Point{x: i, y: j}] = (int(c-'0')-1+scoreIncrement)%9 + 1
		}
	}

	return r
}

func solvePart1(grid Grid1) int {
	best := make(map[Point]int)
	queue := make([]State, 0)

	start := Point{x: 0, y: 0}
	best[start] = 0
	queue = append(queue, State{0, start})

	for len(queue) > 0 {
		currentState := queue[0]
		queue = queue[1:]

		next := make([]GridPoint, 0)
		for _, d := range Directions {
			n, ok := grid.Points[currentState.Pos.add(d)]
			if ok {
				if b, ok := best[n.Point]; ok {
					vnext := currentState.Score + n.Value
					if b > vnext {
						best[n.Point] = vnext
						next = append(next, n)
					}
				} else {
					best[n.Point] = currentState.Score + n.Value
					next = append(next, n)
				}
			}
		}
		sort.Slice(next, func(i, j int) bool { return next[i].Value < next[j].Value })

		for _, gp := range next {
			queue = append(queue, State{currentState.Score + gp.Value, gp.Point})
		}
	}

	return best[Point{x: grid.X - 1, y: grid.Y - 1}]
}

func solvePart2(grid Grid2) int {
	queue := make(PriorityQueue, grid.X*grid.Y)
	allPoints := make(map[Point]*QueueItem)

	idx := 0
	for j := 0; j < grid.Y; j++ {
		for i := 0; i < grid.X; i++ {
			pt := Point{x: i, y: j}
			var qi = QueueItem{pt, math.MaxInt, idx}
			if i == 0 && j == 0 {
				qi.Score = 0
			}
			queue[idx] = &qi
			allPoints[pt] = &qi
			idx++
		}
	}

	heap.Init(&queue)

	for queue.Len() > 0 {
		item := heap.Pop(&queue).(*QueueItem)

		for _, d := range Directions {
			next := item.Point.add(d)
			value, ok := grid.Points[item.Point.add(d)]
			if ok {
				qi := allPoints[next]
				newScore := item.Score + value
				if newScore < qi.Score {
					queue.update(qi, newScore)
				}
			}
		}
	}

	return (*allPoints[Point{x: grid.X - 1, y: grid.Y - 1}]).Score
}

func (p Point) add(p2 Point) Point {
	return Point{p.x + p2.x, p.y + p2.y}
}

func (q PriorityQueue) Len() int { return len(q) }

func (q PriorityQueue) Less(i, j int) bool {
	return q[i].Score < q[j].Score
}

func (q PriorityQueue) Swap(i, j int) {
	q[i], q[j] = q[j], q[i]
	q[i].Index = i
	q[j].Index = j
}

func (q *PriorityQueue) Push(x interface{}) {
	n := len(*q)
	point := x.(*QueueItem)
	point.Index = n
	*q = append(*q, point)
}

func (q *PriorityQueue) Pop() interface{} {
	old := *q
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.Index = -1
	*q = old[0 : n-1]
	return item
}

func (q *PriorityQueue) update(item *QueueItem, score int) {
	item.Score = score
	heap.Fix(q, item.Index)
}
