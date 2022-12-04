package day4

import "fmt"

var startAreaOne = 0
var endAreaOne = 0
var startAreaTwo = 0
var endAreaTwo = 0

func TaskA(lines []string) int {
	sum := 0

	for _, line := range lines {

		_, err := fmt.Sscanf(line, "%d-%d,%d-%d", &startAreaOne, &endAreaOne, &startAreaTwo, &endAreaTwo)
		if err != nil {
			return 0
		}
		if startAreaOne <= startAreaTwo && endAreaOne >= endAreaTwo || startAreaTwo <= startAreaOne && endAreaTwo >= endAreaOne {
			sum++
		}

	}
	return sum
}

func TaskB(lines []string) int {
	sum := 0

	for _, line := range lines {

		_, err := fmt.Sscanf(line, "%d-%d,%d-%d", &startAreaOne, &endAreaOne, &startAreaTwo, &endAreaTwo)
		if err != nil {
			return 0
		}
		if startAreaOne <= endAreaTwo && endAreaOne >= startAreaTwo {
			sum++
		}

	}
	return sum
}
