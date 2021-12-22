package day20

import (
	"fmt"
	"strings"
)

type (
	Grid struct {
		grid                   map[xy]bool
		def                    bool
		minX, maxX, minY, maxY int
	}
)

func NewGrid() *Grid {
	return &Grid{
		grid: map[xy]bool{},
		def:  false,
	}
}

func (g *Grid) Flag(p xy) {
	g.grid[p] = true

	if p.x < g.minX {
		g.minX = p.x
	}

	if p.x > g.maxX {
		g.maxX = p.x
	}

	if p.y < g.minY {
		g.minY = p.y
	}

	if p.y > g.maxY {
		g.maxY = p.y
	}
}

func (g *Grid) String() string {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("def: %v\n", g.def))
	for y := g.minY - 1; y <= g.maxY+1; y++ {
		for x := g.minX - 1; y <= g.maxX+1; x++ {
			if g.grid[xy{x, y}] == !g.def {
				sb.WriteByte('#')
			} else {
				sb.WriteByte('.')
			}
			sb.WriteByte('\n')
		}
	}

	return sb.String()
}

func (g *Grid) ReadIndex(p xy) int {
	var out int
	for y := p.y - 1; y <= p.y+1; y++ {
		for x := p.x - 1; x <= p.x+1; x++ {
			out <<= 1
			if g.grid[xy{x, y}] {
				out |= 1
			}
		}
	}
	if g.def {
		return (^out) & 0b111111111
	}
	return out
}

func (g *Grid) Step(enhance string) *Grid {
	next := NewGrid()

	def := enhance[0]
	if g.def {
		def = enhance[511]
	}

	next.def = def == '#'

	toCheck := map[xy]bool{}
	for p := range g.grid {
		x, y := p.x, p.y
		toCheck[xy{x - 1, y - 1}] = true
		toCheck[xy{x, y - 1}] = true
		toCheck[xy{x + 1, y - 1}] = true
		toCheck[xy{x - 1, y}] = true
		toCheck[xy{x, y}] = true
		toCheck[xy{x + 1, y}] = true
		toCheck[xy{x - 1, y + 1}] = true
		toCheck[xy{x, y + 1}] = true
		toCheck[xy{x + 1, y + 1}] = true
	}

	for p := range toCheck {
		idx := g.ReadIndex(p)
		if enhance[idx] != def {
			next.Flag(p)
		}
	}

	return next
}
