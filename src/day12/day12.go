package day12

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"unicode"
)

type (
	Day12 struct{}

	Graph map[string][]string
)

func (d Day12) Part1() {
	graph := readFile()
	visitedNodes := make(map[string]int, 0)
	paths := getPaths(graph, "start", visitedNodes)
	fmt.Printf("The answer is: %d\n", paths)
}

func (d Day12) Part2() {
	graph := readFile()
	visitedNodes := make(map[string]int, 0)
	paths := getPaths2(graph, "start", visitedNodes, true)
	fmt.Printf("The answer is: %d\n", paths)
}

func getPaths(graph Graph, currentNode string, visitedNodes map[string]int) int {
	if currentNode == "end" {
		return 1
	}

	visitedNodes[currentNode] += 1
	total := 0
	for _, node := range graph[currentNode] {
		if visitedNodes[node] == 0 || toUppercase(node) == node {
			total += getPaths(graph, node, visitedNodes)
		}
	}
	visitedNodes[currentNode] -= 1
	return total
}

func getPaths2(graph Graph, currentNode string, visitedNodes map[string]int, allowsTwice bool) int {
	if currentNode == "end" {
		return 1
	}

	visitedNodes[currentNode] += 1
	total := 0
	for _, node := range graph[currentNode] {
		if visitedNodes[node] == 0 || toUppercase(node) == node {
			total += getPaths2(graph, node, visitedNodes, allowsTwice)
		} else if allowsTwice && node != "start" {
			total += getPaths2(graph, node, visitedNodes, false)
		}
	}
	visitedNodes[currentNode] -= 1
	return total
}

func readFile() Graph {
	file, err := os.Open("src/day12/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	graph := make(Graph, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		l := scanner.Text()
		d := strings.Split(l, "-")
		graph.addEdge(d[0], d[1])
		graph.addEdge(d[1], d[0])
	}

	return graph
}

func (g *Graph) addEdge(u, v string) {
	(*g)[u] = append((*g)[u], v)
}

func toUppercase(s string) string {
	for _, r := range s {
		if unicode.IsLower(r) && unicode.IsLetter(r) {
			return string(unicode.ToUpper(r))
		}
	}
	return s
}
