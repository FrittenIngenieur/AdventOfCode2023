package main

import (
	day1 "AdventOfCode2023/Day1"
	helperFiles "AdventOfCode2023/HelperFiles"
)

func main() {
	callDay1()
}

func callDay1() {
	// Day 1
	var lines []string
	lines = helperFiles.ReadInput(".\\day1\\input_Part1.txt")
	day1.PartOne(lines)

	day1.PartTwo(lines)
}
