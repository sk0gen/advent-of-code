package day9

import (
	"strconv"
	"strings"
)

var directions = map[string]Direction{
	"R": {x: 0, y: 1},
	"U": {x: -1, y: 0},
	"D": {x: 1, y: 0},
	"L": {x: 0, y: -1},
}

func TaskA(lines []string) (res int) {
	var head = Point{x: 0, y: 0}
	var tail = Point{x: 0, y: 0}
	visited := map[Point]bool{}
	visited[tail] = true

	for _, line := range lines {
		split := strings.Split(line, " ")
		amount, _ := strconv.Atoi(split[1])
		dir := directions[split[0]]

		for i := 0; i < amount; i++ {
			moveA(&head, &tail, dir)
			visited[tail] = true
		}
	}
	return len(visited)
}

func moveA(head *Point, tail *Point, direction Direction) {
	head.Move(direction)
	tail.Follow(*head)
}

func abs(value int) int {
	if value < 0 {
		return -1 * value
	}

	return value
}

func TaskB(lines []string) (res int) {
	knots := make([]Point, 10)
	for i := range knots {
		knots[i] = Point{x: 0, y: 0}
	}

	visited := map[Point]bool{}
	visited[Point{x: 0, y: 0}] = true

	for _, line := range lines {
		split := strings.Split(line, " ")
		amount, _ := strconv.Atoi(split[1])
		dir := directions[split[0]]

		for i := 0; i < amount; i++ {
			moveB(&knots, &dir)
			visited[knots[len(knots)-1]] = true
		}
	}
	return len(visited)
}

func moveB(knots *[]Point, direction *Direction) {
	(*knots)[0].Move(*direction)

	for i := 1; i < len(*knots); i++ {
		(*knots)[i].Follow((*knots)[i-1])
	}
}
