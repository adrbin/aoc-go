package main

import (
	"aoc-go/utils"
	"bufio"
	"fmt"
)

var directions = [][]int{
	{-1, -1},
	{-1, 0},
	{-1, 1},
	{0, -1},
	{0, 1},
	{1, -1},
	{1, 0},
	{1, 1},
}

func part1(scanner *bufio.Scanner) string {
	diagram := [][]rune{}
	for scanner.Scan() {
		text := scanner.Text()
		diagram = append(diagram, []rune(text))
	}

	result := 0

	for y := 0; y < len(diagram); y++ {
		for x := 0; x < len(diagram[0]); x++ {
			if diagram[y][x] != '@' {
				continue
			}

			count := countPaper(diagram, x, y)
			if count < 4 {
				result++
			}
		}
	}

	return fmt.Sprint(result)
}

func countPaper(diagram [][]rune, x int, y int) int {
	count := 0
	for _, direction := range directions {
		dx := direction[0]
		dy := direction[1]

		newX := x + dx
		newY := y + dy

		if newY >= 0 && newY < len(diagram) && newX >= 0 && newX < len(diagram[0]) && diagram[newY][newX] == '@' {
			count++
		}
	}

	return count
}

func part2(scanner *bufio.Scanner) string {
	diagram := [][]rune{}
	for scanner.Scan() {
		text := scanner.Text()
		diagram = append(diagram, []rune(text))
	}

	result := 0
	removed := -1

	for removed != 0 {
		removed = 0
		for y := 0; y < len(diagram); y++ {
			for x := 0; x < len(diagram[0]); x++ {
				if diagram[y][x] != '@' {
					continue
				}

				count := countPaper(diagram, x, y)
				if count < 4 {
					diagram[y][x] = '.'
					result++
					removed++
				}
			}
		}
	}

	return fmt.Sprint(result)
}

func main() {
	utils.Run(part1, part2)
}
