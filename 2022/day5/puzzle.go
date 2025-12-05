package main

import (
	"aoc-go/utils"
	"bufio"
	"regexp"
	"strconv"
	"strings"
)

func buildStacks() []utils.Stack[string] {
	return []utils.Stack[string]{
		{S: strings.Split("BZT", "")},
		{S: strings.Split("VHTDN", "")},
		{S: strings.Split("BFMD", "")},
		{S: strings.Split("TJGWVQL", "")},
		{S: strings.Split("WDGPVFQM", "")},
		{S: strings.Split("VZQGHFS", "")},
		{S: strings.Split("ZSNRLTCW", "")},
		{S: strings.Split("ZHWDJNRM", "")},
		{S: strings.Split("MQLFDS", "")},
	}
}

func part1(scanner *bufio.Scanner) string {
	exp := regexp.MustCompile(`move (\d+) from (\d+) to (\d+)`)
	stacks := buildStacks()

	for scanner.Scan() {
		if !strings.Contains(scanner.Text(), "move") {
			continue
		}

		match := exp.FindSubmatch([]byte(scanner.Text()))
		count, err := strconv.Atoi(string(match[1]))
		utils.Check(err)
		from, err := strconv.Atoi(string(match[2]))
		utils.Check(err)
		to, err := strconv.Atoi(string(match[3]))
		utils.Check(err)

		for i := 0; i < count; i++ {
			value := stacks[from-1].Pop()
			stacks[to-1].Push(value)
		}
	}

	builder := strings.Builder{}

	for _, stack := range stacks {
		builder.WriteString(stack.Peek())
	}

	return builder.String()
}

func part2(scanner *bufio.Scanner) string {
	exp := regexp.MustCompile(`move (\d+) from (\d+) to (\d+)`)
	stacks := buildStacks()

	for scanner.Scan() {
		if !strings.Contains(scanner.Text(), "move") {
			continue
		}

		match := exp.FindSubmatch([]byte(scanner.Text()))
		count, err := strconv.Atoi(string(match[1]))
		utils.Check(err)
		from, err := strconv.Atoi(string(match[2]))
		utils.Check(err)
		to, err := strconv.Atoi(string(match[3]))
		utils.Check(err)

		tempStack := utils.Stack[string]{}

		for i := 0; i < count; i++ {
			value := stacks[from-1].Pop()
			tempStack.Push(value)
		}

		for i := 0; i < count; i++ {
			value := tempStack.Pop()
			stacks[to-1].Push(value)
		}
	}

	builder := strings.Builder{}

	for _, stack := range stacks {
		builder.WriteString(stack.Peek())
	}

	return builder.String()
}

func main() {
	utils.Run(part1, part2)
}
