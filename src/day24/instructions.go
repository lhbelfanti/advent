package day24

import (
	"strconv"
	"strings"
)

type (
	State struct {
		Regs  []int // length 4 slice
		input []int
	}

	Instructions interface {
		Exec(st *State)
	}

	Reg int

	Inp struct {
		reg Reg
	}

	Add struct {
		dst Reg
		src Reg
		val int
	}

	Mul struct {
		dst Reg
		src Reg
		val int
	}

	Div struct {
		dst Reg
		src Reg
		val int
	}

	Mod struct {
		dst Reg
		src Reg
		val int
	}

	Eql struct {
		dst Reg
		src Reg
		val int
	}
)

const (
	RegW Reg = iota
	RegX
	RegY
	RegZ
	RegNone = -1
)

func (i Inp) Exec(st *State) {
	st.Regs[i.reg] = st.input[0]
	st.input = st.input[1:]
}

func (i Add) Exec(st *State) {
	if i.src == RegNone {
		st.Regs[i.dst] += i.val
	} else {
		st.Regs[i.dst] += st.Regs[i.src]
	}
}

func (i Mul) Exec(st *State) {
	if i.src == RegNone {
		st.Regs[i.dst] *= i.val
	} else {
		st.Regs[i.dst] *= st.Regs[i.src]
	}
}

func (i Div) Exec(st *State) {
	if i.src == RegNone {
		st.Regs[i.dst] /= i.val
	} else {
		st.Regs[i.dst] /= st.Regs[i.src]
	}
}

func (i Mod) Exec(st *State) {
	if i.src == RegNone {
		st.Regs[i.dst] %= i.val
	} else {
		st.Regs[i.dst] %= st.Regs[i.src]
	}
}

func (i Eql) Exec(st *State) {
	var b bool
	if i.src == RegNone {
		b = st.Regs[i.dst] == i.val
	} else {
		b = st.Regs[i.dst] == st.Regs[i.src]
	}

	st.Regs[i.dst] = 0
	if b {
		st.Regs[i.dst] = 1
	}
}

func parseInstructions(data []string) []Instructions {
	var instructions = make([]Instructions, 0, len(data))
	for _, s := range data {
		sp := strings.Split(s, " ")
		switch sp[0] {
		case "inp":
			i := Inp{}
			i.reg, _ = parseRegOrNumber(sp[1])
			instructions = append(instructions, i)
		case "add":
			i := Add{}
			i.dst, _ = parseRegOrNumber(sp[1])
			i.src, i.val = parseRegOrNumber(sp[2])
			instructions = append(instructions, i)
		case "mul":
			i := Mul{}
			i.dst, _ = parseRegOrNumber(sp[1])
			i.src, i.val = parseRegOrNumber(sp[2])
			instructions = append(instructions, i)
		case "div":
			i := Div{}
			i.dst, _ = parseRegOrNumber(sp[1])
			i.src, i.val = parseRegOrNumber(sp[2])
			instructions = append(instructions, i)
		case "mod":
			i := Mod{}
			i.dst, _ = parseRegOrNumber(sp[1])
			i.src, i.val = parseRegOrNumber(sp[2])
			instructions = append(instructions, i)
		case "eql":
			i := Eql{}
			i.dst, _ = parseRegOrNumber(sp[1])
			i.src, i.val = parseRegOrNumber(sp[2])
			instructions = append(instructions, i)
		}
	}

	return instructions
}

func parseRegOrNumber(s string) (Reg, int) {
	switch s {
	case "w":
		return RegW, 0
	case "x":
		return RegX, 0
	case "y":
		return RegY, 0
	case "z":
		return RegZ, 0
	default:
		i, _ := strconv.Atoi(s)
		return RegNone, i
	}
}
