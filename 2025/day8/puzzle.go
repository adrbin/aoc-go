package main

import (
	"aoc-go/utils"
	"bufio"
	"cmp"
	"fmt"
	"maps"
	"math"
	"slices"
	"strconv"
	"strings"
)

type Point struct {
	X float64
	Y float64
	Z float64
}

type PointPair struct {
	P1 Point
	P2 Point
}

type Distance struct {
	PointPair PointPair
	Dist      float64
}

func part1(scanner *bufio.Scanner) string {
	points := []Point{}
	circuits := []map[Point]bool{}

	for scanner.Scan() {
		text := scanner.Text()
		split := strings.Split(text, ",")
		x, err := strconv.Atoi(split[0])
		utils.Check(err)
		y, err := strconv.Atoi(split[1])
		utils.Check(err)
		z, err := strconv.Atoi(split[2])
		utils.Check(err)

		point := Point{float64(x), float64(y), float64(z)}
		circuit := map[Point]bool{point: true}
		circuits = append(circuits, circuit)
		points = append(points, point)
	}

	distances := []Distance{}
	visited := map[PointPair]bool{}

	for _, p1 := range points {
		for _, p2 := range points {
			if p1 != p2 && !visited[PointPair{p1, p2}] && !visited[PointPair{p2, p1}] {
				distance := Distance{PointPair{p1, p2}, calculateDistance(p1, p2)}
				distances = append(distances, distance)
				visited[PointPair{p1, p2}] = true
			}
		}
	}

	slices.SortFunc(distances, func(a, b Distance) int {
		return cmp.Compare(a.Dist, b.Dist)
	})

	round := 0
	for i, distance := range distances {
		if i == 1000 {
			break
		}

		c1Index := -1
		c2Index := -1
		for j, circuit := range circuits {
			if circuit[distance.PointPair.P1] {
				c1Index = j
			}

			if circuit[distance.PointPair.P2] {
				c2Index = j
			}
		}

		round++

		if c1Index == c2Index {
			continue
		}

		maps.Copy(circuits[c1Index], circuits[c2Index])
		circuits = slices.Delete(circuits, c2Index, c2Index+1)
	}

	slices.SortFunc(circuits, func(a, b map[Point]bool) int {
		return cmp.Compare(len(b), len(a))
	})

	sum := len(circuits[0]) * len(circuits[1]) * len(circuits[2])

	return fmt.Sprint(sum)
}

// func part1(scanner *bufio.Scanner) string {
// 	points := []Point{}

// 	for scanner.Scan() {
// 		text := scanner.Text()
// 		split := strings.Split(text, ",")
// 		x, err := strconv.Atoi(split[0])
// 		utils.Check(err)
// 		y, err := strconv.Atoi(split[1])
// 		utils.Check(err)
// 		z, err := strconv.Atoi(split[2])
// 		utils.Check(err)

// 		point := Point{float64(x), float64(y), float64(z)}
// 		points = append(points, point)
// 	}

// 	distances := []Distance{}

// 	for _, p1 := range points {
// 		for _, p2 := range points {
// 			if p1 != p2 {
// 				distance := Distance{p1, p2, calculateDistance(p1, p2)}
// 				distances = append(distances, distance)
// 			}
// 		}
// 	}

// 	slices.SortFunc(distances, func(a, b Distance) int {
// 		return cmp.Compare(a.Dist, b.Dist)
// 	})

// 	circuits := []map[Point]bool{}
// 	count := 0
// 	for _, distance := range distances {
// 		if count == 10 {
// 			break
// 		}

// 		circuit1Index := -1
// 		circuit2Index := -1
// 		alreadyAdded := false
// 		for i, circuit := range circuits {
// 			p1AlreadyInCircuit := circuit[distance.P1]
// 			p2AlreadyInCircuit := circuit[distance.P2]
// 			if p1AlreadyInCircuit && p2AlreadyInCircuit {
// 				alreadyAdded = true
// 				continue
// 			}

// 			if p1AlreadyInCircuit {
// 				circuit[distance.P1] = true
// 				circuit[distance.P2] = true
// 				circuit1Index = i
// 			}
// 			if p2AlreadyInCircuit {
// 				circuit[distance.P1] = true
// 				circuit[distance.P2] = true
// 				circuit2Index = i
// 			}
// 		}

// 		if alreadyAdded {
// 			continue
// 		}

// 		count++

// 		if circuit1Index == -1 && circuit2Index == -1 {
// 			circuits = append(circuits, map[Point]bool{distance.P1: true, distance.P2: true})
// 			continue
// 		}

// 		if circuit1Index == -1 || circuit2Index == -1 || circuit1Index == circuit2Index {
// 			continue
// 		}

// 		maps.Copy(circuits[circuit1Index], circuits[circuit2Index])
// 		circuits = slices.Delete(circuits, circuit2Index, circuit2Index+1)
// 	}

// 	slices.SortFunc(circuits, func(a, b map[Point]bool) int {
// 		return cmp.Compare(len(b), len(a))
// 	})

// 	sum := len(circuits[0]) * len(circuits[1]) * len(circuits[2])

// 	return fmt.Sprint(sum)
// }

// func part1(scanner *bufio.Scanner) string {
// 	points := []Point{}

// 	for scanner.Scan() {
// 		text := scanner.Text()
// 		split := strings.Split(text, ",")
// 		x, err := strconv.Atoi(split[0])
// 		utils.Check(err)
// 		y, err := strconv.Atoi(split[1])
// 		utils.Check(err)
// 		z, err := strconv.Atoi(split[2])
// 		utils.Check(err)

// 		points = append(points, Point{float64(x), float64(y), float64(z)})
// 	}

// 	circuits := []map[Point]bool{}
// 	for _, point := range points {
// 		minPoint := Point{}
// 		minDistance := math.Inf(1)
// 		for _, other := range points {
// 			distance := calculateDistance(other, point)
// 			if point != other && distance < minDistance {
// 				minPoint = other
// 				minDistance = distance
// 			}
// 		}

// 		added := false
// 		for _, circuit := range circuits {
// 			if _, ok := circuit[minPoint]; ok {
// 				circuit[point] = true
// 				added = true
// 			}
// 		}

// 		if !added {
// 			circuits = append(circuits, map[Point]bool{point: true})
// 		}
// 	}

// 	slices.SortFunc(circuits, func(a, b map[Point]bool) int {
// 		return cmp.Compare(len(b), len(a))
// 	})

// 	sum := len(circuits[0]) * len(circuits[1]) * len(circuits[2])

// 	return fmt.Sprint(sum)
// }

func calculateDistance(p1, p2 Point) float64 {
	distanceX := math.Pow(p2.X-p1.X, 2)
	distanceY := math.Pow(p2.Y-p1.Y, 2)
	distanceZ := math.Pow(p2.Z-p1.Z, 2)

	return math.Sqrt(distanceX + distanceY + distanceZ)
}

func part2(scanner *bufio.Scanner) string {
	points := []Point{}
	circuits := []map[Point]bool{}

	for scanner.Scan() {
		text := scanner.Text()
		split := strings.Split(text, ",")
		x, err := strconv.Atoi(split[0])
		utils.Check(err)
		y, err := strconv.Atoi(split[1])
		utils.Check(err)
		z, err := strconv.Atoi(split[2])
		utils.Check(err)

		point := Point{float64(x), float64(y), float64(z)}
		circuit := map[Point]bool{point: true}
		circuits = append(circuits, circuit)
		points = append(points, point)
	}

	distances := []Distance{}
	visited := map[PointPair]bool{}

	for _, p1 := range points {
		for _, p2 := range points {
			if p1 != p2 && !visited[PointPair{p1, p2}] && !visited[PointPair{p2, p1}] {
				distance := Distance{PointPair{p1, p2}, calculateDistance(p1, p2)}
				distances = append(distances, distance)
				visited[PointPair{p1, p2}] = true
			}
		}
	}

	slices.SortFunc(distances, func(a, b Distance) int {
		return cmp.Compare(a.Dist, b.Dist)
	})

	round := 0
	for _, distance := range distances {
		c1Index := -1
		c2Index := -1
		for j, circuit := range circuits {
			if circuit[distance.PointPair.P1] {
				c1Index = j
			}

			if circuit[distance.PointPair.P2] {
				c2Index = j
			}
		}

		round++

		if c1Index == c2Index {
			continue
		}

		maps.Copy(circuits[c1Index], circuits[c2Index])
		circuits = slices.Delete(circuits, c2Index, c2Index+1)

		if len(circuits) == 1 {
			result := int(distance.PointPair.P1.X * distance.PointPair.P2.X)
			return fmt.Sprint(result)
		}
	}

	return "Did not found"
}

func main() {
	utils.Run(part1, part2)
}
