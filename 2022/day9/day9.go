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
	var hx, hy, tx, ty = 0, 0, 0, 0
	visited := map[Point]bool{}
	visited[Point{x: tx, y: ty}] = true

	for _, line := range lines {
		split := strings.Split(line, " ")
		amount, _ := strconv.Atoi(split[1])
		dir := directions[split[0]]

		for i := 0; i < amount; i++ {
			moveA(&hx, &hy, &tx, &ty, dir)
			visited[Point{tx, ty}] = true
		}
	}
	return len(visited)
}

func moveA(hx *int, hy *int, tx *int, ty *int, direction Direction) {
	*hx += direction.x
	*hy += direction.y

	if !touching(hx, hy, tx, ty) {
		if *hx != *tx {
			*tx += (*hx - *tx) / abs(*hx-*tx)
		}
		if *hy != *ty {
			*ty += (*hy - *ty) / abs(*hy-*ty)
		}

	}
}

func abs(value int) int {
	if value < 0 {
		return -1 * value
	}

	return value
}
func touching(x1 *int, y1 *int, x2 *int, y2 *int) bool {
	return abs(*x1-*x2) <= 1 && abs(*y1-*y2) <= 1
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
