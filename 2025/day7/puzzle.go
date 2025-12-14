package main

import (
	"aoc-go/utils"
	"bufio"
	"fmt"
	"strings"
)

func part1(scanner *bufio.Scanner) string {
	manifold := []string{}

	for scanner.Scan() {
		text := scanner.Text()
		manifold = append(manifold, text)
	}

	startIndex := strings.Index(manifold[0], "S")
	beams := map[int]bool{startIndex: true}

	count := 0

	for y := 1; y < len(manifold); y++ {
		newBeams := map[int]bool{}
		for beam := range beams {
			char := manifold[y][beam]
			if char == '.' {
				newBeams[beam] = true
				continue
			}

			if char == '^' {
				if beam > 0 {
					newBeams[beam-1] = true
				}

				if beam < len(manifold[y]) {
					newBeams[beam+1] = true
				}

				count++

				continue
			}

			panic(fmt.Sprintf("Incorrect character in the manifold: %c", char))
		}

		beams = newBeams
	}

	return fmt.Sprint(count)
}

func part2(scanner *bufio.Scanner) string {
	manifold := []string{}

	for scanner.Scan() {
		text := scanner.Text()
		manifold = append(manifold, text)
	}

	startIndex := strings.Index(manifold[0], "S")
	beams := map[int]int{startIndex: 1}

	for y := 1; y < len(manifold); y++ {
		newBeams := map[int]int{}
		for beamIndex, beamCount := range beams {
			char := manifold[y][beamIndex]
			if char == '.' {
				newBeams[beamIndex] += beamCount
				continue
			}

			if char == '^' {
				if beamIndex > 0 {
					newBeams[beamIndex-1] += beamCount
				}

				if beamIndex < len(manifold[y]) {
					newBeams[beamIndex+1] += beamCount
				}

				continue
			}

			panic(fmt.Sprintf("Incorrect character in the manifold: %c", char))
		}

		beams = newBeams
	}

	sum := 0
	for _, beam := range beams {
		sum += beam
	}

	return fmt.Sprint(sum)
}

func main() {
	utils.Run(part1, part2)
}
