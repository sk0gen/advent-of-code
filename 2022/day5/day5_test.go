package day5

import (
	"testing"
)

func TestATestInput(t *testing.T) {
	file := readFile("input_test.txt")
	result := TaskA(file)

	expected := "CMZ"
	if result != expected {
		t.Errorf("got %s, wanted %s", result, expected)
	}
}

func TestATaskInput(t *testing.T) {
	file := readFile("input.txt")
	result := TaskA(file)

	expected := "NTWZZWHFV"
	if result != expected {
		t.Errorf("got %s, wanted %s", result, expected)
	}
}

func TestBTestInput(t *testing.T) {
	file := readFile("input_test.txt")
	result := TaskB(file)

	expected := "MCD"
	if result != expected {
		t.Errorf("got %s, wanted %s", result, expected)
	}
}

func TestBTaskInput(t *testing.T) {
	file := readFile("input.txt")
	result := TaskB(file)

	expected := "BRZGFVBTJ"
	if result != expected {
		t.Errorf("got %s, wanted %s", result, expected)
	}
}
