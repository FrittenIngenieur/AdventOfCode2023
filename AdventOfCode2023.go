package main

import (
	day1 "AdventOfCode2023/Day1"
	day2 "AdventOfCode2023/Day2"
	helperFiles "AdventOfCode2023/HelperFiles"
)

const DAYTOCOMPILE = 2

func main() {
	switch DAYTOCOMPILE {
	case 1:
		callDay1()
	case 2:
		callDay2()
	}
}

func callDay1() {
	// Day 1
	var lines []string
	lines = helperFiles.ReadInput(".\\day1\\input_Part1.txt")
	day1.PartOne(lines)
	day1.PartTwo(lines)
}

func callDay2() {
	// Day 2
	var lines []string
	lines = helperFiles.ReadInput(".\\day2\\input_Part1.txt")
	day2.PartOne(lines)
	day2.PartTwo(lines)
}
