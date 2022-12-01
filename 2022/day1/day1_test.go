package day1

import (
	"advent-of-code/utils"
	"testing"
)

func TestATestInput(t *testing.T) {

	lines := utils.ReadLines("input_test.txt")
	result := TaskA(lines)

	expected := 24000
	if result != expected {
		t.Errorf("got %d, wanted %d", result, expected)
	}
}

func TestATaskInput(t *testing.T) {

	lines := utils.ReadLines("input.txt")
	result := TaskA(lines)

	expected := 70764

	if result != expected {
		t.Errorf("got %d, wanted %d", result, expected)
	}
}

func TestBTestInput(t *testing.T) {
	lines := utils.ReadLines("input_test.txt")
	result := TaskB(lines)

	expected := 45000
	if result != expected {
		t.Errorf("got %d, wanted %d", result, expected)
	}
}

func TestBTaskInput(t *testing.T) {
	lines := utils.ReadLines("input.txt")
	result := TaskB(lines)

	expected := 203905
	if result != expected {
		t.Errorf("got %d, wanted %d", result, expected)
	}
}
