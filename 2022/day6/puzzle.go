package main

import (
	"aoc-go/utils"
	"bufio"
	"container/list"
	"fmt"
)

func part1(scanner *bufio.Scanner) string {
	scanner.Scan()
	text := scanner.Text()
	l := list.New()

	for i, r := range text {
		if l.Len() >= 4 {
			if checkList(l) {
				return fmt.Sprint(i)
			}

			l.Remove(l.Front())
		}

		l.PushBack(r)
	}

	panic("Didn't find the index")
}

func checkList(l *list.List) bool {
	for e1 := l.Front(); e1 != nil; e1 = e1.Next() {
		for e2 := e1.Next(); e2 != nil; e2 = e2.Next() {
			if e1.Value == e2.Value {
				return false
			}
		}
	}

	return true
}

func part2(scanner *bufio.Scanner) string {
	scanner.Scan()
	text := scanner.Text()
	l := list.New()

	for i, r := range text {
		if l.Len() >= 14 {
			if checkList(l) {
				return fmt.Sprint(i)
			}

			l.Remove(l.Front())
		}

		l.PushBack(r)
	}

	panic("Didn't find the index")
}

func main() {
	utils.Run(part1, part2)
}
