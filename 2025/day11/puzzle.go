package main

import (
	"aoc-go/utils"
	"bufio"
	"fmt"
	"strings"
)

type Node struct {
	Code     string
	OutCodes []string
}

func part1(scanner *bufio.Scanner) string {
	nodes := buildNodes(scanner)
	memo := make(map[string]int)
	count := countPaths("you", "out", nodes, memo)

	return fmt.Sprint(count)
}

func part2(scanner *bufio.Scanner) string {
	nodes := buildNodes(scanner)

	// Case 1: svr -> dac -> fft -> out
	case1 := countPathsWithMemo("svr", "dac", nodes) *
		countPathsWithMemo("dac", "fft", nodes) *
		countPathsWithMemo("fft", "out", nodes)

	// Case 2: svr -> fft -> dac -> out
	case2 := countPathsWithMemo("svr", "fft", nodes) *
		countPathsWithMemo("fft", "dac", nodes) *
		countPathsWithMemo("dac", "out", nodes)

	count := case1 + case2

	return fmt.Sprint(count)
}

func buildNodes(scanner *bufio.Scanner) map[string]Node {
	nodes := map[string]Node{}
	for scanner.Scan() {
		text := scanner.Text()
		split := strings.Split(text, ": ")
		code := split[0]

		outCodes := strings.Split(split[1], " ")
		nodes[code] = Node{code, outCodes}
	}

	return nodes
}

func countPathsWithMemo(currentCode, destinationNode string, nodes map[string]Node) int {
	memo := make(map[string]int)
	return countPaths(currentCode, destinationNode, nodes, memo)
}

func countPaths(currentCode, destinationNode string, nodes map[string]Node, memo map[string]int) int {
	// Check if we've already calculated this
	if count, exists := memo[currentCode]; exists {
		return count
	}

	// Base case: reached the end node
	if currentCode == destinationNode {
		return 1
	}

	// Get the current node
	node, exists := nodes[currentCode]
	if !exists {
		return 0
	}

	// Count paths from all next nodes
	totalPaths := 0
	for _, nextCode := range node.OutCodes {
		totalPaths += countPaths(nextCode, destinationNode, nodes, memo)
	}

	// Store in memo before returning
	memo[currentCode] = totalPaths
	return totalPaths
}

func main() {
	utils.Run(part1, part2)
}
