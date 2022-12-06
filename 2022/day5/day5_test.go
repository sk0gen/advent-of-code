package day5

import (
	"testing"
)

func TestTaskATestInput(t *testing.T) {
	file := readFile("input_test.txt")
	result := TaskA(file)

	expected := "CMZ"
	if result != expected {
		t.Errorf("got %s, wanted %s", result, expected)
	}
}

func TestTaskATaskInput(t *testing.T) {
	file := readFile("input.txt")
	result := TaskA(file)

	expected := "NTWZZWHFV"
	if result != expected {
		t.Errorf("got %s, wanted %s", result, expected)
	}
}

func TestTaskBTestInput(t *testing.T) {
	file := readFile("input_test.txt")
	result := TaskB(file)

	expected := "MCD"
	if result != expected {
		t.Errorf("got %s, wanted %s", result, expected)
	}
}

func TestTaskBTaskInput(t *testing.T) {
	file := readFile("input.txt")
	result := TaskB(file)

	expected := "BRZGFVBTJ"
	if result != expected {
		t.Errorf("got %s, wanted %s", result, expected)
	}
}
