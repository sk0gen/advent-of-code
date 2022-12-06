package day6

func TaskA(input string) int {
	return findMarker(input, 4)
}

func findMarker(input string, markerLength int) int {
	for i := 0; i < len(input); i++ {
		if unique(input[i : i+markerLength]) {
			return i + markerLength
		}
	}
	return 0
}

func TaskB(input string) int {
	return findMarker(input, 14)
}

func unique(arr string) bool {
	m := make(map[rune]bool)
	for _, i := range arr {
		_, ok := m[i]
		if ok {
			return false
		}

		m[i] = true
	}

	return true
}
