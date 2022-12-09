package day8

import (
	"strconv"
)

// TaskA
/*
30373
25512
65332
33549
35390
*/

func TaskA(input []string) int {
	trees := fillTrees(input)
	return calculateVisibleTrees(&trees)
}

func fillTrees(input []string) [][]int {
	trees := make([][]int, len(input))
	for i := range trees {
		cols := len(input[0])
		trees[i] = make([]int, cols)
		for j := 0; j < cols; j++ {
			trees[i][j], _ = strconv.Atoi(string(input[i][j]))
		}
	}

	return trees
}

func transpose(slice [][]int) *[][]int {
	xl := len(slice[0])
	yl := len(slice)
	result := make([][]int, xl)
	for i := range result {
		result[i] = make([]int, yl)
	}
	for i := 0; i < xl; i++ {
		for j := 0; j < yl; j++ {
			result[i][j] = slice[j][i]
		}
	}
	return &result
}

func calculateVisibleTrees(trees *[][]int) int {
	visible := 0
	rows := len(*trees)
	columns := len((*trees)[0])
	transposed := transpose(*trees)
	for i := 0; i < rows; i++ {
		for j := 0; j < columns; j++ {
			currentTreeHeight := (*trees)[i][j]

			if j == 0 || highestTree((*trees)[i][:j]) < currentTreeHeight {
				visible++
			} else if (j == columns-1) || highestTree((*trees)[i][j+1:]) < currentTreeHeight {
				visible++
			} else if i == 0 || highestTree((*transposed)[j][:i]) < currentTreeHeight {
				visible++
			} else if i == rows-1 || highestTree((*transposed)[j][i+1:]) < currentTreeHeight {
				visible++
			}
		}
	}

	return visible
}

func highestTree(trees []int) int {
	var highest = 0
	for i := 0; i < len(trees); i++ {
		highest = maxInt(highest, (trees)[i])
	}

	return highest
}

func maxInt(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func TaskB(input []string) int {
	max := 0
	trees := fillTrees(input)
	for row, treeRow := range trees {
		for col := range treeRow {
			max = maxInt(max, calculateScenicScore(&trees, Point{row: row, column: col}))
		}
	}
	return max
}

func calculateScenicScore(trees *[][]int, startingPoint Point) int {
	var score = 1
	var currentPosition = startingPoint
	rows := len(*trees)
	columns := len((*trees)[0])

	var directions = []Point{
		{1, 0},
		{0, 1},
		{-1, 0},
		{0, -1}}

	for _, direction := range directions {
		currentDirectionScore := 0

		currentPosition.Move(direction)
		for currentPosition.WithinRange(rows, columns) {
			currentDirectionScore++
			if (*trees)[startingPoint.row][startingPoint.column] <= (*trees)[currentPosition.row][currentPosition.column] {
				break
			}
			currentPosition.Move(direction)

		}
		score = score * (currentDirectionScore)
		currentPosition = startingPoint
	}

	return score
}
