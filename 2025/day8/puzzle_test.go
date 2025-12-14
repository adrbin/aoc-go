package main

import (
	"aoc-go/utils"
	"testing"
)

func TestPart1(t *testing.T) {
	utils.TestPuzzle(t, "140008", part1)
}

func TestPart2(t *testing.T) {
	utils.TestPuzzle(t, "9253260633", part2)
}

func BenchmarkPart2(b *testing.B) {
	for b.Loop() {
		utils.BenchmarkPuzzle(b, "9253260633", part2)
	}
}
