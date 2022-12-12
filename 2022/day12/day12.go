package day12

import (
	"image"
)

func TaskA(lines []string) int {
	var start, end image.Point
	grid := map[image.Point]rune{}
	for x, s := range lines {
		for y, r := range s {
			grid[image.Point{X: x, Y: y}] = r
			if r == 'S' {
				start = image.Point{X: x, Y: y}
			} else if r == 'E' {
				end = image.Point{X: x, Y: y}
			}
		}
	}
	grid[start], grid[end] = 'a', 'z'

	dist := map[image.Point]int{end: 0}
	queue := []image.Point{end}
	var shortest *image.Point

	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]

		if grid[cur] == 'a' && shortest == nil {
			shortest = &cur
		}

		for _, d := range []image.Point{{0, -1}, {1, 0}, {0, 1}, {-1, 0}} {
			next := cur.Add(d)
			_, seen := dist[next]
			_, valid := grid[next]

			if !seen && valid && grid[cur] <= grid[next]+1 {
				dist[next] = dist[cur] + 1
				queue = append(queue, next)
			}
		}
	}

	return dist[start]
}

func TaskB(lines []string) int {

	var start, end image.Point
	grid := map[image.Point]rune{}
	for x, s := range lines {
		for y, r := range s {
			grid[image.Point{X: x, Y: y}] = r
			if r == 'S' {
				start = image.Point{X: x, Y: y}
			} else if r == 'E' {
				end = image.Point{X: x, Y: y}
			}
		}
	}
	grid[start], grid[end] = 'a', 'z'

	dist := map[image.Point]int{end: 0}
	queue := []image.Point{end}
	var shortest *image.Point

	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]

		if grid[cur] == 'a' && shortest == nil {
			shortest = &cur
		}

		for _, d := range []image.Point{{0, -1}, {1, 0}, {0, 1}, {-1, 0}} {
			next := cur.Add(d)
			_, seen := dist[next]
			_, valid := grid[next]

			if !seen && valid && grid[cur] <= grid[next]+1 {
				dist[next] = dist[cur] + 1
				queue = append(queue, next)
			}
		}
	}
	return dist[*shortest]
}
