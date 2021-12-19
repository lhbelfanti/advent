package day18

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func readFile() []string {
	file, err := os.Open("src/day18/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	pairs := make([]string, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		pairs = append(pairs, scanner.Text())
	}

	return pairs
}

// Example: "[2, 1]"
func newPair(s string) *Pair {
	p := parseValue(newStream(s), nil, 0)
	switch v := p.(type) {
	case *Pair:
		return v
	default:
		panic(s)
	}
}

func newStream(s string) *Stream {
	return &Stream{
		s:   s,
		pos: 0,
	}
}

func parseValue(s *Stream, parent *Pair, depth int) interface{} {
	next := s.next()
	switch next {
	case "[":
		p := &Pair{parent: parent, depth: depth}
		p.left = parseValue(s, p, depth+1)
		s.next() // ","
		p.right = parseValue(s, p, depth+1)
		s.next() // "]"
		return p
	default: // Number
		i, _ := strconv.Atoi(next)
		return &Num{
			parent: parent,
			depth:  depth,
			value:  i,
		}
	}
}

func (s *Stream) next() string {
	out := string(s.s[s.pos])
	s.pos++
	return out
}
