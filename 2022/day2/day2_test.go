package day2

import (
	"advent-of-code/utils"
	"testing"
)

func TestATestInput(t *testing.T) {

	lines := utils.ReadLines("input_test.txt")
	result := TaskA(lines)

	expected := 15
	if result != expected {
		t.Errorf("got %d, wanted %d", result, expected)
	}
}

func TestATaskInput(t *testing.T) {

	lines := utils.ReadLines("input.txt")
	result := TaskA(lines)

	expected := 8933

	if result != expected {
		t.Errorf("got %d, wanted %d", result, expected)
	}
}

func TestBTestInput(t *testing.T) {
	lines := utils.ReadLines("input_test.txt")
	result := TaskB(lines)

	expected := 12
	if result != expected {
		t.Errorf("got %d, wanted %d", result, expected)
	}
}

func TestBTaskInput(t *testing.T) {
	lines := utils.ReadLines("input.txt")
	result := TaskB(lines)

	expected := 11998
	if result != expected {
		t.Errorf("got %d, wanted %d", result, expected)
	}
}
