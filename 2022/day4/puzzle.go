package main

import (
	"aoc-go/utils"
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

type Assignment struct {
	min int
	max int
}

type OverlapFunc func(assignment1, assignment2 Assignment) bool

func part1(scanner *bufio.Scanner) string {
	result := make(chan int)

	i := 0
	for scanner.Scan() {
		go calculateLine(scanner.Text(), isAssignmentInside, result)
		i++
	}

	total := utils.SumChan(i, result)

	return fmt.Sprint(total)
}

func part2(scanner *bufio.Scanner) string {
	result := make(chan int)

	i := 0
	for scanner.Scan() {
		go calculateLine(scanner.Text(), isOverlapping, result)
		i++
	}

	total := utils.SumChan(i, result)

	return fmt.Sprint(total)
}

func calculateLine(line string, overlapFunc OverlapFunc, result chan int) {
	assignments := strings.Split(line, ",")
	assignment1 := parseAssignment(assignments[0])
	assignment2 := parseAssignment(assignments[1])

	if overlapFunc(assignment1, assignment2) || overlapFunc(assignment2, assignment1) {
		result <- 1
		return
	}

	result <- 0
}

func parseAssignment(text string) Assignment {
	ranges := strings.Split(text, "-")
	min, err := strconv.Atoi(ranges[0])
	utils.Check(err)
	max, err := strconv.Atoi(ranges[1])
	utils.Check(err)
	return Assignment{
		min: min,
		max: max,
	}
}

func isAssignmentInside(assignment1, assignment2 Assignment) bool {
	return assignment1.min >= assignment2.min && assignment1.max <= assignment2.max
}

func isOverlapping(assignment1, assignment2 Assignment) bool {
	return (assignment1.min >= assignment2.min && assignment1.min <= assignment2.max) || (assignment1.max >= assignment2.min && assignment1.max <= assignment2.max)
}

func main() {
	utils.Run(part1, part2)
}
