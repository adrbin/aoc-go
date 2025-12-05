package main

import (
	"aoc-go/utils"
	"bufio"
	"strconv"
)

func part1(scanner *bufio.Scanner) string {
	sum := 0
	for scanner.Scan() {
		text := scanner.Text()
		sum += findTheBiggestNumber(text, 2)
	}

	return strconv.Itoa(sum)
}

func part2(scanner *bufio.Scanner) string {
	sum := 0
	for scanner.Scan() {
		text := scanner.Text()
		sum += findTheBiggestNumber(text, 12)
	}

	return strconv.Itoa(sum)
}

func findTheBiggestNumber(candidates string, length int) int {
	chosen := ""

	for i := 0; i < length; i++ {
		max := 0
		index := 0

		for j := 0; j <= len(candidates)+len(chosen)-length; j++ {
			number, err := strconv.Atoi(string(candidates[j]))
			utils.Check(err)
			if number > max {
				max = number
				index = j
			}
		}

		chosen += strconv.Itoa(max)
		candidates = candidates[index+1:]
	}

	number, err := strconv.Atoi(chosen)
	utils.Check(err)

	return number
}

// // Initial version passing both parts
// // but we don't need to take all combinations of max indexes,
// // the first one is enough because it is always the biggest
// func findTheBiggestNumber(candidates string, length int) int {
// 	max := 0

// 	var findRecursively func(candidates string, chosen string)

// 	findRecursively = func(candidates string, chosen string) {
// 		if len(chosen) == length {
// 			number, err := strconv.Atoi(chosen)
// 			utils.Check(err)
// 			if number > max {
// 				max = number
// 			}

// 			return
// 		}

// 		max := 0
// 		indexes := []int{}

// 		for i := 0; i <= len(candidates)+len(chosen)-length; i++ {
// 			number, err := strconv.Atoi(string(candidates[i]))
// 			utils.Check(err)
// 			if number == max {
// 				indexes = append(indexes, i)
// 			} else if number > max {
// 				max = number
// 				indexes = []int{i}
// 			}
// 		}

// 		for _, i := range indexes {
// 			newChosen := chosen + string(candidates[i])
// 			findRecursively(candidates[i+1:], newChosen)
// 		}
// 	}

// 	findRecursively(candidates, "")

// 	return max
// }

// // Version only for part 1
// func findTheBiggestNumber(s string) int {
// 	max := 0
// 	for i := 0; i < len(s)-1; i++ {
// 		for j := i + 1; j < len(s); j++ {
// 			numberText := string(s[i]) + string(s[j])
// 			number, err := strconv.Atoi(numberText)
// 			utils.Check(err)
// 			if number > max {
// 				max = number
// 			}
// 		}
// 	}

// 	return max
// }

// Does not finish in reasonable time
// func findTheBiggestNumber2(candidates string, length int) int {
// 	max := 0

// 	var findRecursively func(candidates string, chosen string)

// 	findRecursively = func(candidates string, chosen string) {
// 		if len(chosen) == length {
// 			number, err := strconv.Atoi(chosen)
// 			utils.Check(err)
// 			if number > max {
// 				max = number
// 			}

// 			return
// 		}

// 		for i := 0; i <= len(candidates)+len(chosen)-length; i++ {
// 			newChosen := chosen + string(candidates[i])
// 			findRecursively(candidates[i+1:], newChosen)
// 		}
// 	}

// 	findRecursively(candidates, "")

// 	print(max)
// 	return max
// }

func main() {
	utils.Run(part1, part2)
}
