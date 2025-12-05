package main

import (
	"aoc-go/utils"
	"bufio"
	"fmt"
	"math"
	"strconv"
	"strings"
)

func part1(scanner *bufio.Scanner) string {
	dial := 50
	password := 0
	for scanner.Scan() {
		text := scanner.Text()
		direction := text[0]
		value, err := strconv.Atoi(strings.TrimSpace(text[1:]))
		utils.Check(err)

		if direction == 'R' {
			dial = utils.Mod(dial+value, 100)
		} else {
			dial = utils.Mod(dial-value, 100)
		}

		if dial == 0 {
			password++
		}
	}

	return fmt.Sprint(password)
}

func part2(scanner *bufio.Scanner) string {
	dial := 50
	password := 0
	for scanner.Scan() {
		text := scanner.Text()
		direction := text[0]
		value, err := strconv.Atoi(strings.TrimSpace(text[1:]))
		utils.Check(err)

		zeroesPassed := 0

		if direction == 'R' {
			zeroesPassed = int(math.Abs(float64(dial+value)) / 100)
			dial = utils.Mod(dial+value, 100)
		} else {
			inverted := 0
			if dial > 0 {
				inverted = (100 - dial)
			}
			zeroesPassed = int(math.Abs(float64(inverted+value)) / 100)
			dial = utils.Mod(dial-value, 100)
		}

		password += zeroesPassed
	}

	return fmt.Sprint(password)
}

func main() {
	utils.Run(part1, part2)
}
