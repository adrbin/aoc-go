package main

import (
	"aoc-go/utils"
	"bufio"
	"fmt"
	"sort"
	"strconv"
)

func part1(scanner *bufio.Scanner) string {
	max := 0
	sum := 0
	for scanner.Scan() {
		text := scanner.Text()

		if text == "" {
			if sum > max {
				max = sum
			}
			sum = 0
			continue
		}

		num, err := strconv.Atoi(text)
		utils.Check(err)
		sum += num
	}

	return fmt.Sprint(max)
}

func part2(scanner *bufio.Scanner) string {
	sums := make([]int, 100)
	curSum := 0
	for scanner.Scan() {
		text := scanner.Text()

		if text == "" {
			sums = append(sums, curSum)
			curSum = 0
			continue
		}

		num, err := strconv.Atoi(text)
		utils.Check(err)
		curSum += num
	}

	sort.Ints(sums)

	result := sums[len(sums)-1] + sums[len(sums)-2] + sums[len(sums)-3]

	return fmt.Sprint(result)
}

func main() {
	utils.Run(part1, part2)
}
