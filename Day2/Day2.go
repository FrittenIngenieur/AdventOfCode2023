// Day 2 - Advent of Code 2023
// https://adventofcode.com/2023/day/2
// =====================================================================================================================

// Game 77: 17 red, 6 green; 7 green, 12 red, 4 blue; 4 red, 7 green, 14 blue

package day2

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type Cubes struct {
	amount_blue  int
	amount_red   int
	amount_green int
	isPlausible  bool
}

type Game struct {
	gameID      int
	cubes       []Cubes
	isPlausible bool
}

const MAX_RED_CUBES = 12
const MAX_GREEN_CUBES = 13
const MAX_BLUE_CUBES = 14

func PartOne(lines []string) {
	fmt.Println("Advent of Code 2023 - Day 2 - Part One")

	var games []Game

	games = pushGamesToStruct(lines)

	fmt.Println(sumGameIds(games))
}

func PartTwo(lines []string) {
	fmt.Println("Advent of Code 2023 - Day 2 - Part Two")
	var games []Game

	games = pushGamesToStruct(lines)

	fmt.Println(calculatePowerOfGames(games))
}

func pushGamesToStruct(lines []string) []Game {
	var games []Game

	for index, line := range lines {
		lineWithoutGame := strings.Replace(string(line), fmt.Sprintf("Game %d: ", index+1), "", 1)

		lineWithoutGame = strings.Replace(lineWithoutGame, " ", "", -1)

		cubePairs := strings.Split(lineWithoutGame, ";")

		amountBlue, amountRed, amountGreen := 0, 0, 0

		gameIsPlausible, blueIsPlausible, greenIsPlausible, redIsPlausible := true, true, true, true

		var cubesOfGame []Cubes

		for _, pair := range cubePairs {
			colorKeys := strings.Split(pair, ",")

			for _, color := range colorKeys {
				patternAmount := regexp.MustCompile("[0-9]*")
				matchAmount := patternAmount.FindString(color)

				if strings.Contains(color, "blue") {
					amountBlue, _ = strconv.Atoi(matchAmount)
					if amountBlue > MAX_BLUE_CUBES {
						blueIsPlausible = false
					}
				}
				if strings.Contains(color, "red") {
					amountRed, _ = strconv.Atoi(matchAmount)
					if amountRed > MAX_RED_CUBES {
						redIsPlausible = false
					}
				}
				if strings.Contains(color, "green") {
					amountGreen, _ = strconv.Atoi(matchAmount)
					if amountGreen > MAX_GREEN_CUBES {
						greenIsPlausible = false
					}
				}
			}

			pairIsPlausible := false

			if blueIsPlausible == true && greenIsPlausible == true && redIsPlausible == true {
				pairIsPlausible = true
			} else {
				pairIsPlausible = false
			}

			blueIsPlausible = true
			redIsPlausible = true
			greenIsPlausible = true

			cubes := Cubes{amount_blue: amountBlue, amount_red: amountRed, amount_green: amountGreen, isPlausible: pairIsPlausible}
			cubesOfGame = append(cubesOfGame, cubes)
		}

		gameIsPlausible = true
		for _, c := range cubesOfGame {
			if c.isPlausible == false {
				gameIsPlausible = false
				break
			}
		}

		games = append(games, Game{gameID: index + 1, cubes: cubesOfGame, isPlausible: gameIsPlausible})

	}
	return games
}

func sumGameIds(g []Game) int {
	var sum int = 0
	for _, game := range g {
		if game.isPlausible {
			sum += game.gameID
		}
	}
	return sum
}

func calculatePowerOfGames(g []Game) int {
	powerOfAllGames := 0
	for _, game := range g {
		minRedCubes, minBlueCubes, minGreenCubes := 0, 0, 0

		for _, cubes := range game.cubes {
			if cubes.amount_blue > minBlueCubes {
				minBlueCubes = cubes.amount_blue
			}
			if cubes.amount_green > minGreenCubes {
				minGreenCubes = cubes.amount_green
			}
			if cubes.amount_red > minRedCubes {
				minRedCubes = cubes.amount_red
			}
		}

		powerOfAllGames = powerOfAllGames + (minBlueCubes * minGreenCubes * minRedCubes)

	}

	return powerOfAllGames
}
