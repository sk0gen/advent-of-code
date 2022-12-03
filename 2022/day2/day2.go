package day2

func TaskA(lines []string) (res int) {
	pointsMap := map[string]int{
		"A X": 4,
		"A Y": 8,
		"A Z": 3,
		"B X": 1,
		"B Y": 5,
		"B Z": 9,
		"C X": 7,
		"C Y": 2,
		"C Z": 6,
	}

	var points = 0

	for _, line := range lines {
		points += pointsMap[line]
	}

	return points
}

func TaskB(lines []string) (res int) {
	pointsMap := map[string]int{
		"A X": 3,
		"A Y": 4,
		"A Z": 8,
		"B X": 1,
		"B Y": 5,
		"B Z": 9,
		"C X": 2,
		"C Y": 6,
		"C Z": 7,
	}

	var points = 0

	for _, line := range lines {
		points += pointsMap[line]
	}

	return points
}
