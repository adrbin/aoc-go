package main

import (
	"aoc-go/utils"
	"testing"
)

func TestPart1(t *testing.T) {
	utils.TestPuzzle(t, "552", part1)
}

func TestPart2(t *testing.T) {
	utils.TestPuzzle(t, "307608674109300", part2)
}
