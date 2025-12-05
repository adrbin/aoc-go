package main

import (
	"aoc-go/utils"
	"bufio"
	"fmt"
)

func part1(scanner *bufio.Scanner) string {
	total := 0

	for scanner.Scan() {
		total += getPriority(scanner.Text())
	}

	return fmt.Sprint(total)
}

func part2(scanner *bufio.Scanner) string {
	total := 0

	for scanner.Scan() {
		var rucksacks [3]string
		rucksacks[0] = scanner.Text()
		scanner.Scan()
		rucksacks[1] = scanner.Text()
		scanner.Scan()
		rucksacks[2] = scanner.Text()
		total += getBadge(rucksacks[:])
	}

	return fmt.Sprint(total)
}

func getPriority(line string) int {
	middle := len(line) / 2
	compartment1 := line[:middle]
	compartment2 := utils.BuildSet(line[middle:])

	for _, r := range compartment1 {
		_, hasKey := compartment2[r]
		if hasKey {
			return calculatePriority(r)
		}
	}

	panic("Couldn't find the priority for line " + line)
}

func getBadge(rucksacks []string) int {
	rucksack1 := rucksacks[0]
	rucksack2 := utils.BuildSet(rucksacks[1])
	rucksack3 := utils.BuildSet(rucksacks[2])

	for _, r := range rucksack1 {
		_, hasKey2 := rucksack2[r]
		_, hasKey3 := rucksack3[r]
		if hasKey2 && hasKey3 {
			return calculatePriority(r)
		}
	}

	panic("Couldn't find the badge for line " + rucksacks[0])
}

func calculatePriority(r rune) int {
	if r >= 'a' {
		return int(r-'a') + 1
	}

	return int(r-'A') + 27
}

func main() {
	utils.Run(part1, part2)
}
