package utils

import (
	"bufio"
	"fmt"
	"os"
	"testing"
)

type PuzzleFunc func(s *bufio.Scanner) string

func Check(e error) {
	if e != nil {
		panic(e)
	}
}

func Run(part1, part2 PuzzleFunc) {
	fileName := "input.txt"
	if len(os.Args) > 1 {
		fileName = os.Args[1]
	}

	f, err := os.Open(fileName)
	Check(err)
	scanner := bufio.NewScanner(f)
	result1 := part1(scanner)
	fmt.Println(result1)
	f.Seek(0, 0)
	scanner = bufio.NewScanner(f)
	result2 := part2(scanner)
	fmt.Println(result2)
}

func TestPuzzle(t *testing.T, expected string, puzzleFunc PuzzleFunc) {
	f, err := os.Open("input.txt")
	Check(err)
	scanner := bufio.NewScanner(f)
	result := puzzleFunc(scanner)

	if result != expected {
		t.Fatalf(`Expected %q, instead of %q`, expected, result)
	}
}

func BenchmarkPuzzle(b *testing.B, expected string, puzzleFunc PuzzleFunc) {
	f, err := os.Open("input.txt")
	Check(err)
	scanner := bufio.NewScanner(f)
	result := puzzleFunc(scanner)

	if result != expected {
		b.Fatalf(`Expected %q, instead of %q`, expected, result)
	}
}

func BuildSet(text string) map[rune]struct{} {
	exists := struct{}{}
	set := make(map[rune]struct{})
	for _, r := range text {
		set[r] = exists
	}

	return set
}

func SumChan(length int, ch chan int) int {
	total := 0
	for range length {
		total += <-ch
	}

	return total
}

type Stack[T any] struct {
	S []T
}

func (stack *Stack[T]) Push(element T) {
	stack.S = append(stack.S, element)
}

func (stack *Stack[T]) Pop() (element T) {
	element = stack.Peek()
	stack.S = stack.S[:len(stack.S)-1]
	return
}

func (stack *Stack[T]) Peek() (element T) {
	element = stack.S[len(stack.S)-1]
	return
}

func Mod(a, b int) int {
	return (a%b + b) % b
}

func Min(a, b int) int {
	if b < a {
		return b
	}

	return a
}

func Max(a, b int) int {
	if b > a {
		return b
	}

	return a
}

func Abs(number int) int {
	if number < 0 {
		return -number
	}

	return number
}
