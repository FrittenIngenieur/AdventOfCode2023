// Day 1 - Advent of Code 2023
// https://adventofcode.com/2023/day/1
// =====================================================================================================================

// NOT THE BEST SOLUTION BUT IT WORKS!

package day1

import (
	"fmt"
	"strconv"
	"strings"
)

type Code struct {
	elfCode string
}

func PartOne(lines []string) {
	fmt.Println("Advent of Code 2023 - Day 1 - Part One")

	var totalValue int = sumNumbers(lines)

	fmt.Println(totalValue)
}

func PartTwo(lines []string) {
	fmt.Println("Advent of Code 2023 - Day 1 - Part Two")

	// Push lines in struct
	var newLines []Code
	for _, lineEl := range lines {
		newLines = append(newLines, Code{elfCode: lineEl})
	}

	// Start from the beginning and change written Number to Number
	for lineIndex, line := range newLines {
		// Iterate letter by letter through the line and check if there is a written number
		for i := 2; i <= len(string(line.elfCode)); i++ {
			subString := line.elfCode[0:i]
			restOfLine := line.elfCode[i:len(string(line.elfCode))]

			replacedNumberInCheck := false
			foundNumberInCheck := false

			// Check if there is already a normal number in this subString
			for _, c := range subString {
				if _, err := strconv.Atoi(string(c)); err == nil {
					foundNumberInCheck = true
				}
			}

			if foundNumberInCheck {
				break
			}

			newSubString := replaceNumberInString(subString)

			// If newSubString and subString are different, a written number was changed
			if newSubString != subString {
				subString = newSubString
				replacedNumberInCheck = true
			}

			if replacedNumberInCheck {
				newLines[lineIndex].elfCode = subString + restOfLine
				break
			}
		}
	}

	// Start from the end and change written Number to Number
	for lineIndex, line := range newLines {
		// Iterate letter by letter through the line and check if there is a written number
		for i := 2; i <= len(string(line.elfCode)); i++ {
			subString := line.elfCode[len(string(line.elfCode))-i : len(string(line.elfCode))]
			restOfLine := line.elfCode[0 : len(string(line.elfCode))-i]

			replacedNumberInCheck := false
			foundNumberInCheck := false

			// Check if there is already a normal number in this subString
			for _, c := range subString {
				if _, err := strconv.Atoi(string(c)); err == nil {
					foundNumberInCheck = true
				}
			}

			if foundNumberInCheck {
				break
			}

			newSubString := replaceNumberInString(subString)

			// If newSubString and subString are different, a written number was changed
			if newSubString != subString {
				subString = newSubString
				replacedNumberInCheck = true
			}

			if replacedNumberInCheck {
				newLines[lineIndex].elfCode = restOfLine + subString
				break
			}
		}
	}

	var rawNumberLines []string

	for _, line := range newLines {
		rawNumberLines = append(rawNumberLines, line.elfCode)
	}

	totalValuePart2 := sumNumbers(rawNumberLines)

	fmt.Println(totalValuePart2)
}

func sumNumbers(lines []string) int {
	var numbers []int

	totalValue := 0

	for _, line := range lines {
		var firstNumber int
		var lastNumber int
		firstNumberFound := false
		for _, c := range line {

			if number, err := strconv.Atoi(string(c)); err == nil {

				if firstNumberFound {
					lastNumber = number
				} else {
					firstNumber = number
					lastNumber = number
					firstNumberFound = true
				}
			}
		}
		nums, _ := strconv.Atoi(fmt.Sprintf("%d%d", firstNumber, lastNumber))
		numbers = append(numbers, nums)
		totalValue += nums
	}

	return totalValue
}

func replaceNumberInString(subString string) string {
	changedString := subString
	if strings.Contains(subString, "one") {
		changedString = strings.Replace(subString, "one", "1", -1)
	}
	if strings.Contains(subString, "two") {
		changedString = strings.Replace(subString, "two", "2", -1)
	}
	if strings.Contains(subString, "three") {
		changedString = strings.Replace(subString, "three", "3", -1)
	}
	if strings.Contains(subString, "four") {
		changedString = strings.Replace(subString, "four", "4", -1)
	}
	if strings.Contains(subString, "five") {
		changedString = strings.Replace(subString, "five", "5", -1)
	}
	if strings.Contains(subString, "six") {
		changedString = strings.Replace(subString, "six", "6", -1)
	}
	if strings.Contains(subString, "seven") {
		changedString = strings.Replace(subString, "seven", "7", -1)
	}
	if strings.Contains(subString, "eight") {
		changedString = strings.Replace(subString, "eight", "8", -1)
	}
	if strings.Contains(subString, "nine") {
		changedString = strings.Replace(subString, "nine", "9", -1)
	}

	return changedString
}
