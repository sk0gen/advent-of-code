package day12

import (
	"advent-of-code/utils"
	"testing"
)

func TestTaskATestInput(t *testing.T) {

	lines := utils.ReadLines("input_test.txt")
	result := TaskA(lines)

	expected := 31
	if result != expected {
		t.Errorf("got %d, wanted %d", result, expected)
	}
}

func TestTaskATaskInput(t *testing.T) {

	lines := utils.ReadLines("input.txt")
	result := TaskA(lines)

	expected := 361

	if result != expected {
		t.Errorf("got %d, wanted %d", result, expected)
	}
}

func TestTaskBTestInput(t *testing.T) {
	lines := utils.ReadLines("input_test.txt")
	result := TaskB(lines)

	expected := 29
	if result != expected {
		t.Errorf("got %d, wanted %d", result, expected)
	}
}

func TestTaskBTaskInput(t *testing.T) {
	lines := utils.ReadLines("input.txt")
	result := TaskB(lines)

	expected := 354
	if result != expected {
		t.Errorf("got %d, wanted %d", result, expected)
	}
}
