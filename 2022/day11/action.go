package day11

type action func(int) int

func plus(a int) action {
	return func(b int) int {
		return a + b
	}
}

func multiply(a int) action {
	return func(b int) int {
		return a * b
	}
}

func square() action {
	return func(b int) int {
		return b * b
	}
}

func div(a int, t action) action {
	return func(b int) int {
		return int(float64(t(b)) / float64(a))
	}
}

func mod(a int, t action) action {
	return func(b int) int {
		return t(b) % a
	}
}
