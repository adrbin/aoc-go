package main

import (
	"aoc-go/utils"
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

func part1(scanner *bufio.Scanner) string {
	array := [][]string{}
	for scanner.Scan() {
		text := scanner.Text()
		array = append(array, strings.Fields(text))
	}

	sum := 0
	for i := 0; i < len(array[0]); i++ {
		sub, err := strconv.Atoi(array[0][i])
		utils.Check(err)
		operator := array[len(array)-1][i]
		for j := 1; j < len(array)-1; j++ {
			value, err := strconv.Atoi(array[j][i])
			utils.Check(err)
			if operator == "+" {
				sub += value
			} else {
				sub *= value
			}
		}

		sum += sub
	}

	return fmt.Sprint(sum)
}

func part2(scanner *bufio.Scanner) string {
	array := []string{}
	for scanner.Scan() {
		text := scanner.Text()
		array = append(array, text)
	}

	indexes := []int{}
	for i := 1; i < len(array[len(array)-1]); i++ {
		operator := array[len(array)-1][i]
		if operator == '+' || operator == '*' {
			indexes = append(indexes, i)
		}
	}

	indexes = append(indexes, len(array[len(array)-1]) + 1)

	sum := 0
	currentIndex := 0
	for _, nextIndex := range indexes {
		operator := array[len(array)-1][currentIndex]
		sub := 0
		if operator == '*' {
			sub = 1
		}

		for i := currentIndex; i < nextIndex-1; i++ {
			var number strings.Builder
			for j := 0; j < len(array)-1; j++ {
				digit := string(array[j][i])
				if digit != " " {
					number.WriteString(string(array[j][i]))
				}
			}
			value, err := strconv.Atoi(number.String())
			utils.Check(err)

			if operator == '+' {
				sub += value
			} else {
				sub *= value
			}
		}

		sum += sub
		currentIndex = nextIndex
	}

	return fmt.Sprint(sum)
}

func main() {
	utils.Run(part1, part2)
}
