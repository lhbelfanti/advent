package day18

import (
	"fmt"
)

type (
	Day18 struct{}

	Pair struct {
		parent *Pair
		depth  int
		left   interface{}
		right  interface{}
	}

	Num struct {
		parent *Pair
		depth  int
		value  int
	}

	Stream struct {
		s   string
		pos int
	}
)

func (p *Pair) String() string {
	// return fmt.Sprintf("(%d)[%v,%v]", p.depth, p.left, p.right)
	return fmt.Sprintf("[%v,%v]", p.left, p.right)
}

func (n *Num) String() string {
	// return fmt.Sprintf("(%d)%d", n.depth, n.value)
	return fmt.Sprintf("%d", n.value)
}

func (d Day18) Part1() {
	pairs := readFile()
	p := newPair(pairs[0])
	for i := 1; i < len(pairs); i++ {
		p = add(p, newPair(pairs[i]))
	}

	mag := magnitude(p)

	fmt.Printf("The answer is: %d\n", mag)
}

func (d Day18) Part2() {
	pairs := readFile()

	maxMag := 0
	for i := 0; i < len(pairs); i++ {
		for j := i + 1; j < len(pairs); j++ {
			s1 := add(newPair(pairs[i]), newPair(pairs[j]))
			m1 := magnitude(s1)
			if m1 > maxMag {
				maxMag = m1
			}

			s2 := add(newPair(pairs[j]), newPair(pairs[i]))
			m2 := magnitude(s2)
			if m2 > maxMag {
				maxMag = m2
			}
		}
	}

	fmt.Printf("The answer is: %d\n", maxMag)
}
