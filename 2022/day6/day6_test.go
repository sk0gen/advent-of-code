package day6

import (
	"os"
	"testing"
)

func TestTaskATestInput(t *testing.T) {
	file, _ := os.ReadFile("input_test.txt")
	result := TaskA(string(file))

	expected := 7
	if result != expected {
		t.Errorf("got %d, wanted %d", result, expected)
	}
}

func TestTaskATaskInput(t *testing.T) {

	file, _ := os.ReadFile("input.txt")
	result := TaskA(string(file))

	expected := 1640

	if result != expected {
		t.Errorf("got %d, wanted %d", result, expected)
	}
}

func TestTaskBTestInput(t *testing.T) {
	file, _ := os.ReadFile("input_test.txt")
	result := TaskB(string(file))

	expected := 19

	if result != expected {
		t.Errorf("got %d, wanted %d", result, expected)
	}
}

func TestTaskBTaskInput(t *testing.T) {
	file, _ := os.ReadFile("input.txt")
	result := TaskB(string(file))

	expected := 3613

	if result != expected {
		t.Errorf("got %d, wanted %d", result, expected)
	}
}
