package main

import (
	"aoc-go/utils"
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

func part1(scanner *bufio.Scanner) string {
	return sumInvalidIds(scanner, isInvalidId)
}

func part2(scanner *bufio.Scanner) string {
	return sumInvalidIds(scanner, isInvalidId2)
}

func sumInvalidIds(scanner *bufio.Scanner, f func(string) bool) string {
	sum := 0
	for scanner.Scan() {
		text := scanner.Text()
		for ranges := range strings.SplitSeq(text, ",") {
			if ranges == "" {
				continue
			}

			split := strings.Split(ranges, "-")
			min, err := strconv.Atoi(split[0])
			utils.Check(err)
			max, err := strconv.Atoi(split[1])
			utils.Check(err)

			for i := min; i <= max; i++ {
				if f(strconv.Itoa(i)) {
					sum += i
				}
			}
		}

	}

	return fmt.Sprint(sum)
}

func isInvalidId(s string) bool {
	if len(s) < 2 || len(s)%2 != 0 {
		return false
	}

	halfIndex := len(s) / 2

	for i := 0; i < halfIndex; i++ {
		if s[i] != s[halfIndex+i] {
			return false
		}
	}

	return true
}

func isInvalidId2(s string) bool {
	if len(s) < 2 {
		return false
	}

	halfIndex := (len(s)) / 2

	for i := 1; i <= halfIndex; i++ {
		repeatNumber := len(s) / i
		if strings.Repeat(s[:i], repeatNumber) == s {
			return true
		}
	}

	return false
}

func main() {
	utils.Run(part1, part2)
}
