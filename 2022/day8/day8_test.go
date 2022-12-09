package day8

import (
	"advent-of-code/utils"
	"testing"
)

func TestTaskATestInput(t *testing.T) {
	lines := utils.ReadLines("input_test.txt")
	result := TaskA(lines)

	expected := 21
	if result != expected {
		t.Errorf("got %d, wanted %d", result, expected)
	}
}

func TestTaskATaskInput(t *testing.T) {

	lines := utils.ReadLines("input.txt")
	result := TaskA(lines)

	expected := 1779
	if result != expected {
		t.Errorf("got %d, wanted %d", result, expected)
	}
}

func TestTaskBTestInput(t *testing.T) {
	lines := utils.ReadLines("input_test.txt")
	result := TaskB(lines)

	expected := 8
	if result != expected {
		t.Errorf("got %d, wanted %d", result, expected)
	}
}

func TestTaskBTaskInput(t *testing.T) {
	lines := utils.ReadLines("input.txt")
	result := TaskB(lines)

	expected := 172224
	if result != expected {
		t.Errorf("got %d, wanted %d", result, expected)
	}
}
