package main

import (
	"aoc-go/utils"
	"bufio"
	"fmt"
)

func part1(scanner *bufio.Scanner) string {
	total := 0

	for scanner.Scan() {
	}

	return fmt.Sprint(total)
}

func part2(scanner *bufio.Scanner) string {
	total := 0

	for scanner.Scan() {
	}

	return fmt.Sprint(total)
}

func main() {
	utils.Run(part1, part2)
}
