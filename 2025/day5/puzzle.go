package main

import (
	"aoc-go/utils"
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

type Range struct {
	Min int
	Max int
}

func part1(scanner *bufio.Scanner) string {
	ranges := []Range{}
	for scanner.Scan() {
		text := strings.TrimSpace(scanner.Text())
		if text == "" {
			break
		}

		r := getRange(text)
		ranges = append(ranges, r)
	}

	count := 0

	for scanner.Scan() {
		text := scanner.Text()

		value, err := strconv.Atoi(text)
		utils.Check(err)

		for _, r := range ranges {
			if value >= r.Min && value <= r.Max {
				count++
				break
			}
		}
	}

	return fmt.Sprint(count)
}

func part2(scanner *bufio.Scanner) string {
	ranges := []*Range{}
	for scanner.Scan() {
		text := strings.TrimSpace(scanner.Text())
		if text == "" {
			break
		}

		r := getRange(text)
		shouldAdd := true

		newRanges := []*Range{}

		for _, other := range ranges {
			if r.Min > other.Min && r.Max < other.Max {
				shouldAdd = false
			}

			if r.Min <= other.Min && r.Max >= other.Max {
				continue
			}

			if r.Min <= other.Max && r.Min >= other.Min {
				r.Min = other.Max + 1
			} else if r.Max >= other.Min && r.Max <= other.Max {
				r.Max = other.Min - 1
			}

			newRanges = append(newRanges, other)
		}

		if shouldAdd && r.Max >= r.Min {
			newRanges = append(newRanges, &r)
		}

		ranges = newRanges
	}

	count := 0

	for _, r := range ranges {
		count += r.Max - r.Min + 1
	}

	return fmt.Sprint(count)
}

func getRange(s string) Range {
	split := strings.Split(s, "-")

	min, err := strconv.Atoi(split[0])
	utils.Check(err)

	max, err := strconv.Atoi(split[1])
	utils.Check(err)

	return Range{
		min,
		max,
	}
}

func main() {
	utils.Run(part1, part2)
}
