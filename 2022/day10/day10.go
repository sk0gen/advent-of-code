package day10

import (
	"fmt"
	"strconv"
	"strings"
)

func TaskA(lines []string) (res int) {

	cycles := []int{20, 60, 100, 140, 180, 220}
	programResult := runProgram(lines)

	output := 0
	for _, cycle := range cycles {
		fmt.Println(programResult[cycle-1] * cycle)
		output += programResult[cycle-1] * cycle
	}

	return output
}

func runProgram(program []string) []int {
	var result []int

	register := 1
	for i := 0; i < len(program); i++ {
		var command = strings.Split(program[i], " ")
		if command[0] == "noop" {
			result = append(result, register)
		} else if command[0] == "addx" {
			result = append(result, register)
			result = append(result, register)
			value, _ := strconv.Atoi(command[1])
			register += value
		}
	}

	return result
}
func TaskB(lines []string) []string {
	sb := strings.Builder{}
	for i, line := range runProgram(lines) {
		var pixel = i % 40
		if pixel == line-1 || pixel == line || pixel == line+1 {
			sb.WriteString("#")
		} else {
			sb.WriteString(".")
		}
	}
	resultString := sb.String()

	var output []string
	for i := 0; i < len(resultString)/40; i++ {
		output = append(output, resultString[i*40:i*40+40])
		fmt.Println(output[i])
	}
	return output
}
