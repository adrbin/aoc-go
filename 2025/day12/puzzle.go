package main

import (
	"aoc-go/utils"
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

func part1(scanner *bufio.Scanner) string {
	shapeSizes := buildShapeSizes(scanner)

	hasText := true
	count := 0
	for hasText {
		text := scanner.Text()
		split := strings.Split(text, ": ")
		sizes := strings.Split(split[0], "x")

		size1, err := strconv.Atoi(sizes[0])
		utils.Check(err)
		size2, err := strconv.Atoi(sizes[1])
		utils.Check(err)

		maxSize := size1 * size2

		size := 0
		for i, count := range strings.Split(split[1], " ") {
			countInt, err := strconv.Atoi(count)
			utils.Check(err)

			size += countInt * shapeSizes[i]
		}

		if size <= maxSize {
			count++
		}

		hasText = scanner.Scan()
	}

	return fmt.Sprint(count)
}

func buildShapeSizes(scanner *bufio.Scanner) map[int]int {
	shapeSizes := map[int]int{}
	for scanner.Scan() {
		text := scanner.Text()
		if strings.ContainsAny(text, "x") {
			break
		}

		index, err := strconv.Atoi(text[:len(text)-1])
		utils.Check(err)

		size := 0
		for scanner.Scan() {
			text := scanner.Text()
			if text == "" {
				break
			}
			size += len(text)
		}

		shapeSizes[index] = size
	}

	return shapeSizes
}

func part2(scanner *bufio.Scanner) string {
	return ""
}

func main() {
	utils.Run(part1, part2)
}
