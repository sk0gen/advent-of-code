package day1

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func Day1() {
	fmt.Println(Task_a(readLines("input.txt")))
}

func Task_a(lines []string) (res int) {
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

func Task_b(lines []string) (res int) {
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

func readLines(filePath string) []string {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	var lines []string
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}
