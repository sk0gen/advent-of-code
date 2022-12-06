package day6

func TaskA(input string) int {
	return findMarker(input, 4)
}

func findMarker(input string, markerLength int) int {
	for i := 0; i < len(input); i++ {
		markerLetters := input[i : i+markerLength]
		lettersMap := make(map[rune]bool)
		for _, letter := range markerLetters {
			lettersMap[letter] = true
		}

		if len(lettersMap) == markerLength {
			return i + markerLength
		}
	}
	return 0
}

func TaskB(input string) int {
	return findMarker(input, 14)
}
