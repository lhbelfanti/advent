package main

import (
	"fmt"

	"advent2024/src/day1"
	"advent2024/src/day2"
	"advent2024/src/day3"
	"advent2024/src/day4"
	"advent2024/src/day5"
	"advent2024/src/day6"
)

func main() {
	fmt.Println("Day 1")
	var d1 day1.Day1
	d1.Part1()
	d1.Part2()

	fmt.Println("Day 2")
	var d2 day2.Day2
	d2.Part1()
	d2.Part2()

	fmt.Println("Day 3")
	var d3 day3.Day3
	d3.Part1()
	d3.Part2()

	fmt.Println("Day 4")
	var d4 day4.Day4
	d4.Part1()
	d4.Part2()

	fmt.Println("Day 5")
	var d5 day5.Day5
	d5.Part1()
	d5.Part2()

	fmt.Println("Day 6")
	var d6 day6.Day6
	d6.Part1()
	d6.Part2()
}
