package day3

import (
	"strings"
	"unicode"
)

func TaskA(lines []string) int {
	sum := 0

	for _, line := range lines {

		firstRucksack := line[0 : len(line)/2]
		secondRucksack := line[len(line)/2:]
		commonChars := make(map[rune]bool)

		for _, char := range firstRucksack {
			if strings.ContainsRune(secondRucksack, char) {
				if _, ok := commonChars[char]; !ok {
					commonChars[char] = true
					sum += checkPriorities(char)
				}
			}
		}

	}
	return sum
}

func TaskB(lines []string) int {

	sum := 0
	rucksacks := chunkBy(lines, 3)
	for _, rucksack := range rucksacks {
		sum += checkPriorities(commonCharBetweenRucksacks(rucksack))
	}

	return sum
}

func commonCharBetweenRucksacks(rucksacks []string) rune {
	for _, char := range rucksacks[0] {
		if strings.ContainsRune(rucksacks[1], char) && strings.ContainsRune(rucksacks[2], char) {
			return char
		}
	}
	return 0
}

func chunkBy(items []string, chunkSize int) (chunks [][]string) {
	for chunkSize < len(items) {
		items, chunks = items[chunkSize:], append(chunks, items[0:chunkSize:chunkSize])
	}

	return append(chunks, items)
}

func checkPriorities(c rune) (res int) {
	if unicode.IsLower(c) {
		return int(c) - 'a' + 1
	} else {
		return int(c) - 'A' + 27
	}
}
