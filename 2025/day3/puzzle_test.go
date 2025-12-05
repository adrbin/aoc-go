package main

import (
	"aoc-go/utils"
	"testing"
)

func TestPart1(t *testing.T) {
	utils.TestPuzzle(t, "17766", part1)
}

func TestPart2(t *testing.T) {
	utils.TestPuzzle(t, "176582889354075", part2)
}
