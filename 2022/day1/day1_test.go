package day1

import "testing"

func TestATestInput(t *testing.T) {

	lines := readLines("input_test.txt")
	result := Task_a(lines)

	expected := 24000
	if result != expected {
		t.Errorf("got %d, wanted %d", result, expected)
	}
}

func TestATaskInput(t *testing.T) {

	lines := readLines("input.txt")
	result := Task_a(lines)

	expected := 70764

	if result != expected {
		t.Errorf("got %d, wanted %d", result, expected)
	}
}

func TestBTestInput(t *testing.T) {
	lines := readLines("input_test.txt")
	result := Task_b(lines)

	expected := 45000
	if result != expected {
		t.Errorf("got %d, wanted %d", result, expected)
	}
}

func TestBTaskInput(t *testing.T) {
	lines := readLines("input.txt")
	result := Task_b(lines)

	expected := 203905
	if result != expected {
		t.Errorf("got %d, wanted %d", result, expected)
	}
}
