package main

import (
	"aoc-go/utils"
	"bufio"
	"bytes"
	"fmt"
	"os/exec"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

const MaxJoltagesLength = 10

type Machine struct {
	TargetDiagram  []rune
	Buttons        [][]int
	TargetJoltages [MaxJoltagesLength]int
	JoltagesLength int
}

func part1(scanner *bufio.Scanner) string {
	machines := createMachines(scanner)

	sum := 0
	for _, machine := range machines {
		diagramLength := len(machine.TargetDiagram)
		initialDiagram := slices.Repeat([]rune{'.'}, diagramLength)
		diagrams := [][]rune{initialDiagram}
		for level := 1; ; level++ {
			newDiagrams := [][]rune{}
			for _, diagram := range diagrams {
				for _, button := range machine.Buttons {
					newDiagram := make([]rune, diagramLength)
					copy(newDiagram, diagram)

					for _, index := range button {
						newDiagram[index] = switchIndicator(newDiagram[index])
					}

					if slices.Equal(newDiagram, machine.TargetDiagram) {
						sum += level
						goto out
					}

					newDiagrams = append(newDiagrams, newDiagram)
				}
			}
			diagrams = newDiagrams
		}
	out:
	}

	return fmt.Sprint(sum)
}

func part2(scanner *bufio.Scanner) string {
	machines := createMachines(scanner)

	sum := 0
	for i, machine := range machines {
		fmt.Printf("Processing machine %d: %d buttons, joltages length=%d\n", i, len(machine.Buttons), machine.JoltagesLength)

		// prefer z3 solver if available
		if _, err := exec.LookPath("z3"); err == nil {
			if level, ok := solveWithZ3(machine); ok {
				fmt.Printf(" machine %d solved by z3 with level=%d\n", i, level)
				sum += level
				continue
			}
			fmt.Printf(" machine %d: z3 found no solution, falling back\n", i)
		} else {
			fmt.Printf(" z3 not found; using fallback solver for machine %d\n", i)
		}

		// fallback: BFS with pruning (original approach), with logging
		if level, ok := fallbackSolveMachine(i, machine); ok {
			sum += level
		}
	}

	return fmt.Sprint(sum)
}

func createMachines(scanner *bufio.Scanner) []Machine {
	machines := []Machine{}
	for scanner.Scan() {
		text := scanner.Text()
		split1 := strings.Split(text, "] ")
		diagram := []rune(split1[0][1:])

		split2 := strings.Split(split1[1], " {")
		buttons := strings.Split(split2[0], " ")
		buttonsArr := [][]int{}
		for _, button := range buttons {
			split3 := strings.Split(button[1:len(button)-1], ",")
			intButton := []int{}
			for _, s := range split3 {
				v, err := strconv.Atoi(s)
				utils.Check(err)
				intButton = append(intButton, v)
			}
			buttonsArr = append(buttonsArr, intButton)
		}

		joltages := split2[1]
		joltagesArr := strings.Split(joltages[:len(joltages)-1], ",")
		joltagesLength := len(joltagesArr)
		if joltagesLength > MaxJoltagesLength {
			panic("Too large joltages length in the input data")
		}
		joltagesNumArr := [MaxJoltagesLength]int{}
		for i := 0; i < joltagesLength; i++ {
			joltage, err := strconv.Atoi(joltagesArr[i])
			utils.Check(err)
			joltagesNumArr[i] = joltage
		}

		machine := Machine{diagram, buttonsArr, joltagesNumArr, joltagesLength}
		machines = append(machines, machine)
	}
	return machines
}

func switchIndicator(r rune) rune {
	if r == '.' {
		return '#'
	}

	if r == '#' {
		return '.'
	}

	panic("Incorrect indicator in the diagram")
}

func checkJoltages(joltages, targetJoltages [MaxJoltagesLength]int, length int) bool {
	for i := 0; i < length; i++ {
		if joltages[i] > targetJoltages[i] {
			return false
		}
	}

	return true
}

// solveWithZ3 tries to find a non-negative integer combination of buttons
// that reaches the target joltages while minimizing total button presses.
// It uses the `z3` CLI with an Optimize objective. Returns minimal level and true if found.
func solveWithZ3(machine Machine) (int, bool) {
	nButtons := len(machine.Buttons)
	if nButtons == 0 {
		return 0, false
	}

	var sb strings.Builder
	sb.WriteString("(set-option :produce-models true)\n")
	for i := range nButtons {
		sb.WriteString(fmt.Sprintf("(declare-const b%d Int)\n", i))
		sb.WriteString(fmt.Sprintf("(assert (>= b%d 0))\n", i))
	}

	// equality constraints per jolt index
	for j := range machine.JoltagesLength {
		sb.WriteString("(assert (= ")
		// build sum term: for each joltage index j, sum over all buttons that affect it
		sb.WriteString("(+")
		for i := range nButtons {
			// check if button i affects joltage index j
			affectsJ := slices.Contains(machine.Buttons[i], j)
			// coefficient is 1 if button i affects j, else 0 (omit from sum)
			if affectsJ {
				sb.WriteString(fmt.Sprintf(" b%d", i))
			}
		}
		sb.WriteString(fmt.Sprintf(" ) %d))\n", machine.TargetJoltages[j]))
	}

	// minimize total presses
	sb.WriteString("(minimize (+")
	for i := range nButtons {
		sb.WriteString(fmt.Sprintf(" b%d", i))
	}
	sb.WriteString("))\n(check-sat)\n(get-model)\n")

	cmd := exec.Command("z3", "-in", "-smt2")
	cmd.Stdin = strings.NewReader(sb.String())
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &out
	if err := cmd.Run(); err != nil {
		fmt.Printf(" z3 run error: %v\n", err)
		return 0, false
	}

	outStr := out.String()
	if !strings.HasPrefix(outStr, "sat") && !strings.Contains(outStr, "sat") {
		return 0, false
	}

	// parse model lines like: (define-fun b0 () Int 3)
	re := regexp.MustCompile(`(?m)\(define-fun\s+(b[0-9]+)\s+\(\)\s+Int\s+(-?[0-9]+)\s*\)`)
	matches := re.FindAllStringSubmatch(outStr, -1)
	if matches == nil {
		return 0, false
	}
	total := 0
	for _, m := range matches {
		val, _ := strconv.Atoi(m[2])
		total += val
	}
	return total, true
}

// fallbackSolveMachine runs the original BFS + Pareto pruning with logging.
func fallbackSolveMachine(index int, machine Machine) (int, bool) {
	fmt.Printf(" fallback solver: machine %d starting BFS\n", index)
	initialJoltages := [MaxJoltagesLength]int{}
	candidateJoltages := map[[MaxJoltagesLength]int]bool{initialJoltages: true}
	for level := 1; ; level++ {
		newCandidateJoltages := map[[MaxJoltagesLength]int]bool{}
		for joltages := range candidateJoltages {
			for _, button := range machine.Buttons {
				newJoltages := joltages
				for _, idx := range button {
					newJoltages[idx]++
				}

				if !checkJoltages(newJoltages, machine.TargetJoltages, machine.JoltagesLength) {
					continue
				}

				if newJoltages == machine.TargetJoltages {
					fmt.Printf(" fallback solver: machine %d solved at level %d\n", index, level)
					return level, true
				}

				newCandidateJoltages[newJoltages] = true
			}
		}

		candidateJoltages = pruneCandidates(newCandidateJoltages, machine.JoltagesLength)
		fmt.Printf(" machine %d level %d candidate count after prune: %d\n", index, level, len(candidateJoltages))
		if len(candidateJoltages) == 0 {
			fmt.Printf(" fallback solver: machine %d no candidates left by level %d\n", index, level)
			return 0, false
		}
		// safety guard: if level grows too large, bail out
		if level > 1000 {
			fmt.Printf(" fallback solver: machine %d giving up after level %d\n", index, level)
			return 0, false
		}
	}
}

// pruneCandidates removes any candidate that is componentwise dominated by another
func pruneCandidates(candidates map[[MaxJoltagesLength]int]bool, length int) map[[MaxJoltagesLength]int]bool {
	keys := make([][MaxJoltagesLength]int, 0, len(candidates))
	for k := range candidates {
		keys = append(keys, k)
	}

	keep := map[[MaxJoltagesLength]int]bool{}
	for i := 0; i < len(keys); i++ {
		a := keys[i]
		dominated := false
		for j := 0; j < len(keys); j++ {
			if i == j {
				continue
			}
			b := keys[j]
			// if a is componentwise <= b (and not equal), then a is dominated
			le := true
			eq := true
			for d := range length {
				if a[d] > b[d] {
					le = false
					break
				}
				if a[d] != b[d] {
					eq = false
				}
			}
			if le && !eq {
				dominated = true
				break
			}
		}
		if !dominated {
			keep[a] = true
		}
	}

	return keep
}

func main() {
	utils.Run(part1, part2)
}
