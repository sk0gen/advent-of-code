package day5

import (
	"fmt"
	"os"
	"strings"
)

func TaskA(input string) string {
	data := strings.Split(input, "\n\n")
	crates := createCrates(data[0])
	newArrangement := calculateRearrangementTaskA(crates, data[1])

	sb := strings.Builder{}
	for _, crate := range newArrangement {
		sb.WriteString(crate.Pop())
	}
	return sb.String()
}
func calculateRearrangementTaskA(crates []Crate, input string) []Crate {
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		move, from, to := extractMoves(line)
		for i := 0; i < move; i++ {
			element := crates[from].Pop()
			crates[to].Push(element)
		}
	}
	return crates
}
func extractMoves(line string) (int, int, int) {
	var move, from, to int

	_, err := fmt.Sscanf(line, "move %d from %d to %d", &move, &from, &to)
	if err != nil {
	}
	return move, from - 1, to - 1
}

func createCrates(input string) []Crate {
	lines := strings.Split(input, "\n")
	if len(lines) == 0 {
		return nil
	}

	crates := make([]Crate, (len(lines[len(lines)-1])+1)/4)

	for i := len(lines) - 2; i >= 0; i-- {
		for k := 0; k <= len(lines[i])/4; k++ {

			var character = string(lines[i][k*4+1])
			if character == " " {

				continue
			}
			crates[k].Push(character)
		}
	}

	return crates
}

func TaskB(input string) string {
	data := strings.Split(input, "\n\n")
	crates := createCrates(data[0])
	newArrangement := calculateRearrangementTaskB(crates, data[1])

	sb := strings.Builder{}
	for _, crate := range newArrangement {
		sb.WriteString(crate.Pop())
	}
	return sb.String()
}

func calculateRearrangementTaskB(crates []Crate, input string) []Crate {
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		move, from, to := extractMoves(line)
		if move == 1 {
			element := crates[from].Pop()
			crates[to].Push(element)
		} else {
			toMove := make([]string, move)
			for i := 0; i < move; i++ {
				toMove[i] = crates[from].Pop()
			}
			for i := len(toMove) - 1; i >= 0; i-- {
				crates[to].Push(toMove[i])
			}
		}

	}
	return crates
}

func readFile(fileName string) string {
	file, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Print(err)
	}
	return string(file)
}
