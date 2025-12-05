package main

import (
	"aoc-go/utils"
	"bufio"
	"fmt"
	"strings"
)

var oppMap = map[string]int{
	"A": 1,
	"B": 2,
	"C": 3,
}

var myMap = map[string]int{
	"X": 1,
	"Y": 2,
	"Z": 3,
}

const LOSE int = 1
const DRAW int = 2
const WIN int = 3

func part1(scanner *bufio.Scanner) string {
	totalScore := 0
	for scanner.Scan() {
		totalScore += getScore1(scanner.Text())
	}

	return fmt.Sprint(totalScore)
}

func part2(scanner *bufio.Scanner) string {
	totalScore := 0
	for scanner.Scan() {
		totalScore += getScore2(scanner.Text())
	}

	return fmt.Sprint(totalScore)
}

func getScore1(line string) int {
	round := strings.Split(line, " ")
	oppShape := oppMap[round[0]]
	myShape := myMap[round[1]]
	return getRoundScore(oppShape, myShape) + myShape
}

func getScore2(line string) int {
	round := strings.Split(line, " ")
	oppShape := oppMap[round[0]]
	result := myMap[round[1]]
	myShape := getMyShape(oppShape, result)
	return getRoundScore(oppShape, myShape) + myShape
}

func getMyShape(oppShape, result int) int {
	switch result {
	case WIN:
		return oppShape%3 + 1
	case LOSE:
		if oppShape == 1 {
			return 3
		}
		return oppShape - 1
	default:
		return oppShape
	}
}

func getRoundScore(oppShape, myShape int) int {
	switch {
	case myShape == oppShape:
		return 3
	case myShape == oppShape+1 || (myShape == 1 && oppShape == 3):
		return 6
	default:
		return 0
	}
}

func main() {
	utils.Run(part1, part2)
}
