package day10

import (
	"advent-of-code/utils"
	"fmt"
	"reflect"
	"strings"
	"testing"
)

func TestTaskATestInput(t *testing.T) {

	lines := utils.ReadLines("input_test.txt")
	result := TaskA(lines)

	expected := 13141
	if result != expected {
		t.Errorf("got %d, wanted %d", result, expected)
	}
}

func TestTaskATaskInput(t *testing.T) {

	lines := utils.ReadLines("input.txt")
	result := TaskA(lines)

	expected := 70764

	if result != expected {
		t.Errorf("got %d, wanted %d", result, expected)
	}
}

func TestTaskBTestInput(t *testing.T) {
	lines := utils.ReadLines("input_test.txt")
	result := TaskB(lines)

	expected := strings.Split("##..##..##..##..##..##..##..##..##..##..\n###...###...###...###...###...###...###.\n####....####....####....####....####....\n#####.....#####.....#####.....#####.....\n######......######......######......####\n#######.......#######.......#######.....", "\n")

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Slices are not the same")
		fmt.Println(result)
		fmt.Println(expected)
	}
}

func TestTaskBTaskInput(t *testing.T) {
	lines := utils.ReadLines("input.txt")
	result := TaskB(lines)

	expected := strings.Split("###..###....##..##..####..##...##..###..\n#..#.#..#....#.#..#....#.#..#.#..#.#..#.\n###..#..#....#.#..#...#..#....#..#.#..#.\n#..#.###.....#.####..#...#.##.####.###..\n#..#.#....#..#.#..#.#....#..#.#..#.#....\n###..#.....##..#..#.####..###.#..#.#....", "\n")

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Slices are not the same")
		fmt.Println(result)
		fmt.Println(expected)
	}
}
