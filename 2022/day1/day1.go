package day1

import (
	"advent-of-code/utils"
	"fmt"
	"sort"
	"strconv"
)

func Day1() {
	fmt.Println(TaskA(utils.ReadLines("input.txt")))
	fmt.Println(TaskB(utils.ReadLines("input.txt")))
}

func TaskA(lines []string) (res int) {
	var max = 0
	var curr = 0

	for i := 0; i < len(lines); i++ {
		if len(lines[i]) == 0 {
			max = maxValue(max, curr)
			curr = 0
		}

		intVal, _ := strconv.Atoi(lines[i])
		curr += intVal
	}
	max = maxValue(max, curr)

	return max
}

func TaskB(lines []string) (res int) {
	var values []int

	var curr = 0

	for i := 0; i < len(lines); i++ {
		if len(lines[i]) == 0 {
			values = append(values, curr)
			curr = 0
		}

		intVal, _ := strconv.Atoi(lines[i])
		curr += intVal
	}
	values = append(values, curr)

	sort.Slice(values, func(i, j int) bool {
		return values[i] > values[j]
	})
	return values[0] + values[1] + values[2]
}

func maxValue(a, b int) int {
	if a > b {
		return a
	}
	return b
}
