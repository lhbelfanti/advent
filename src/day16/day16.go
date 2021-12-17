package day16

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
)

type (
	Day16 struct{}

	Literal struct {
		version int64
		typeID  int64
		value   int64
	}

	Operator struct {
		version  int64
		typeID   int64
		lengthID int64
		length   int64
		packets  []interface{}
	}
)

func (d Day16) Part1() {
	hex := readFile()
	bits := hexToBitArray(hex)
	packet, _ := readPacket(bits, 0)
	sum := sumVersions(packet)

	fmt.Printf("The answer is: %d\n", sum)
}

func (d Day16) Part2() {
	hex := readFile()
	bits := hexToBitArray(hex)
	packet, _ := readPacket(bits, 0)
	a := eval(packet)

	fmt.Printf("The answer is: %d\n", a)
}

func readFile() string {
	file, err := os.Open("src/day16/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var hex string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		hex = scanner.Text()
	}

	return hex
}

func hexToBitArray(s string) []byte {
	out := make([]byte, 0, 4*len(s))
	for _, c := range s {
		switch c {
		case '0':
			out = append(out, []byte{0, 0, 0, 0}...)
		case '1':
			out = append(out, []byte{0, 0, 0, 1}...)
		case '2':
			out = append(out, []byte{0, 0, 1, 0}...)
		case '3':
			out = append(out, []byte{0, 0, 1, 1}...)
		case '4':
			out = append(out, []byte{0, 1, 0, 0}...)
		case '5':
			out = append(out, []byte{0, 1, 0, 1}...)
		case '6':
			out = append(out, []byte{0, 1, 1, 0}...)
		case '7':
			out = append(out, []byte{0, 1, 1, 1}...)
		case '8':
			out = append(out, []byte{1, 0, 0, 0}...)
		case '9':
			out = append(out, []byte{1, 0, 0, 1}...)
		case 'A':
			out = append(out, []byte{1, 0, 1, 0}...)
		case 'B':
			out = append(out, []byte{1, 0, 1, 1}...)
		case 'C':
			out = append(out, []byte{1, 1, 0, 0}...)
		case 'D':
			out = append(out, []byte{1, 1, 0, 1}...)
		case 'E':
			out = append(out, []byte{1, 1, 1, 0}...)
		case 'F':
			out = append(out, []byte{1, 1, 1, 1}...)
		}
	}

	return out
}

func readPacket(data []byte, startPos int) (l interface{}, c int) {
	n := startPos

	version, count := readBits(data, n, 3)
	n += count

	typeID, count := readBits(data, n, 3)
	n += count

	switch typeID {
	case 4:
		value, count := readNumber(data, n)
		n += count

		return Literal{
			version: version,
			typeID:  typeID,
			value:   value,
		}, n - startPos
	default:
		lengthID, count := readBits(data, n, 1)
		n += count

		if lengthID == 0 {
			length, count := readBits(data, n, 15)
			n += count

			op := Operator{
				version:  version,
				typeID:   typeID,
				lengthID: lengthID,
				length:   length,
				packets:  nil,
			}

			subPacketStart := n

			for int64(n-subPacketStart) < length {
				packet, count := readPacket(data, n)
				n += count

				op.packets = append(op.packets, packet)
			}

			return op, n - startPos
		} else {
			length, count := readBits(data, n, 11)
			n += count

			op := Operator{
				version:  version,
				typeID:   typeID,
				lengthID: lengthID,
				length:   length,
				packets:  nil,
			}

			for i := int64(0); i < length; i++ {
				packet, count := readPacket(data, n)
				n += count

				op.packets = append(op.packets, packet)
			}

			return op, n - startPos
		}
	}
}

func readNumber(data []byte, startPos int) (out int64, count int) {
	for {
		part, _ := readBits(data, startPos, 5)
		out <<= 4
		out |= int64(part & 0x0f)
		count += 5
		startPos += 5
		if part&0x10 == 0 {
			break
		}
	}

	return out, count
}

func readBits(data []byte, startPos, count int) (out int64, c int) {
	for _, b := range data[startPos : startPos+count] {
		out <<= 1
		out |= int64(b)
	}

	return out, count
}

func sumVersions(packet interface{}) int64 {
	var sum int64
	switch p := packet.(type) {
	case Literal:
		sum += p.version
	case Operator:
		sum += p.version
		for _, p2 := range p.packets {
			sum += sumVersions(p2)
		}
	}

	return sum
}

func eval(packet interface{}) int64 {
	switch p := packet.(type) {
	case Literal:
		return p.value
	case Operator:
		switch p.typeID {
		case 0:
			var sum int64 = 0
			for _, p2 := range p.packets {
				sum += eval(p2)
			}
			return sum
		case 1:
			var prod int64 = 1
			for _, p2 := range p.packets {
				prod *= eval(p2)
			}
			return prod
		case 2:
			var min int64 = math.MaxInt64
			for _, p2 := range p.packets {
				v := eval(p2)
				if v < min {
					min = v
				}
			}
			return min
		case 3:
			var max int64 = 0
			for _, p2 := range p.packets {
				v := eval(p2)
				if v > max {
					max = v
				}
			}
			return max
		case 5:
			v1 := eval(p.packets[0])
			v2 := eval(p.packets[1])
			if v1 > v2 {
				return 1
			} else {
				return 0
			}
		case 6:
			v1 := eval(p.packets[0])
			v2 := eval(p.packets[1])
			if v1 < v2 {
				return 1
			} else {
				return 0
			}
		case 7:
			v1 := eval(p.packets[0])
			v2 := eval(p.packets[1])
			if v1 == v2 {
				return 1
			} else {
				return 0
			}
		}
	}

	return -1
}
