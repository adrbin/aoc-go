package main

import (
	"aoc-go/utils"
	"bufio"
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Point struct {
	X int
	Y int
}

type Distance struct {
	Min int
	Max int
}

func part1(scanner *bufio.Scanner) string {
	points := []Point{}

	for scanner.Scan() {
		text := scanner.Text()
		split := strings.Split(text, ",")
		x, err := strconv.Atoi(split[0])
		utils.Check(err)
		y, err := strconv.Atoi(split[1])
		utils.Check(err)

		point := Point{x, y}
		points = append(points, point)
	}

	maxArea := 0

	for _, p1 := range points {
		for _, p2 := range points {
			if p1 != p2 {
				area := calculateArea(p1, p2)
				if area > maxArea {
					maxArea = area
				}
			}
		}
	}

	return fmt.Sprint(maxArea)
}

func part2(scanner *bufio.Scanner) string {
	redPoints := map[Point]bool{}
	redPointInX := map[int][]Point{}
	redPointInY := map[int][]Point{}

	for scanner.Scan() {
		text := scanner.Text()
		split := strings.Split(text, ",")
		x, err := strconv.Atoi(split[0])
		utils.Check(err)
		y, err := strconv.Atoi(split[1])
		utils.Check(err)

		point := Point{x, y}
		redPoints[point] = true
		redPointInX[x] = append(redPointInX[x], point)
		redPointInY[y] = append(redPointInY[y], point)
	}

	markedPoints := map[Point]bool{}
	markedPointInY := map[int][]Point{}

	for p1 := range redPoints {
		for _, p2 := range redPointInX[p1.X] {
			if p1 != p2 {
				min := utils.Min(p1.Y, p2.Y)
				max := utils.Max(p1.Y, p2.Y)
				for y := min; y <= max; y++ {
					point := Point{p1.X, y}
					markedPoints[point] = true
					markedPointInY[y] = append(markedPointInY[y], point)
				}
			}
		}

		for _, p2 := range redPointInY[p1.Y] {
			if p1 != p2 {
				min := utils.Min(p1.X, p2.X)
				max := utils.Max(p1.X, p2.X)
				for x := min; x <= max; x++ {
					point := Point{x, p1.Y}
					markedPoints[point] = true
					markedPointInY[p1.Y] = append(markedPointInY[p1.Y], point)
				}
			}
		}
	}

	// visualize(redPoints, markedPoints)

	distances := map[int]Distance{}

	for y, points := range markedPointInY {
		if len(points) < 2 {
			continue
		}

		minX := math.MaxInt
		maxX := 0
		for _, point := range points {
			if point.X < minX {
				minX = point.X
			}

			if point.X > maxX {
				maxX = point.X
			}
		}

		distances[y] = Distance{minX, maxX}
	}

	// visualize(redPoints, markedPoints)

	maxArea := 0

	for p1 := range redPoints {
		for p2 := range redPoints {
			if p1 != p2 {
				area := calculateArea(p1, p2)
				if area > maxArea && checkArea(p1, p2, distances) {
					maxArea = area
				}
			}
		}
	}

	return fmt.Sprint(maxArea)
}

func visualize(redPoints map[Point]bool, markedPoints map[Point]bool) {
	for y := range 20 {
		for x := range 20 {
			if redPoints[Point{x, y}] {
				print("#")
			} else if markedPoints[Point{x, y}] {
				print("X")
			} else {
				print(".")
			}
		}
		println()
	}
	println()
}

func calculateArea(p1, p2 Point) int {
	dx := utils.Abs(p2.X-p1.X) + 1
	dy := utils.Abs(p2.Y-p1.Y) + 1

	return dx * dy
}

func checkArea(p1, p2 Point, distances map[int]Distance) bool {
	minX := utils.Min(p1.X, p2.X)
	minY := utils.Min(p1.Y, p2.Y)

	maxX := utils.Max(p1.X, p2.X)
	maxY := utils.Max(p1.Y, p2.Y)

	for y := minY; y <= maxY; y++ {
		distance := distances[y]
		if minX < distance.Min || maxX > distance.Max {
			return false
		}
	}

	return true
}

func main() {
	utils.Run(part1, part2)
}
