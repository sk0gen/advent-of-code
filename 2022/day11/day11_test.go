package day11

import (
	"testing"
)

func TestTaskATestInput(t *testing.T) {

	result := TaskA(part1Modifier(GenTestMonkeys()))

	expected := 10605
	if result != expected {
		t.Errorf("got %d, wanted %d", result, expected)
	}
}

func TestTaskATaskInput(t *testing.T) {

	result := TaskA(part1Modifier(GenTaskMonkeys()))

	expected := 50172

	if result != expected {
		t.Errorf("got %d, wanted %d", result, expected)
	}
}

func TestTaskBTestInput(t *testing.T) {
	result := TaskB(part2Modifier(GenTestMonkeys()))

	expected := 45000
	if result != expected {
		t.Errorf("got %d, wanted %d", result, expected)
	}
}

func TestTaskBTaskInput(t *testing.T) {
	result := TaskB(GenTaskMonkeys())

	expected := 203905
	if result != expected {
		t.Errorf("got %d, wanted %d", result, expected)
	}
}
func part1Modifier(monkeys []*Monkey) []*Monkey {
	for _, monkey := range monkeys {
		monkey.operation = div(3, monkey.operation)
	}
	return monkeys
}

func part2Modifier(monkeys []*Monkey) []*Monkey {
	modifier := 17 * 2 * 5 * 3 * 7 * 13 * 19 * 11
	for _, monkey := range monkeys {
		monkey.operation = mod(modifier, monkey.operation)
	}
	return monkeys
}

func GenTestMonkeys() []*Monkey {
	var monkeys = []*Monkey{
		{items: []int{79, 98}, operation: multiply(19), testValue: 23, destTrue: 2, destFalse: 3},
		{items: []int{54, 65, 75, 74}, operation: plus(6), testValue: 19, destTrue: 2},
		{items: []int{79, 60, 97}, operation: square(), testValue: 13, destTrue: 1, destFalse: 3},
		{items: []int{74}, operation: plus(3), testValue: 17, destFalse: 1},
	}
	return monkeys
}

func GenTaskMonkeys() []*Monkey {
	monkeys := []*Monkey{
		{items: []int{54, 61, 97, 63, 74}, operation: multiply(7), testValue: 17, destTrue: 5, destFalse: 3},
		{items: []int{61, 70, 97, 64, 99, 83, 52, 87}, operation: plus(8), testValue: 2, destTrue: 7, destFalse: 6},
		{items: []int{60, 67, 80, 65}, operation: multiply(13), testValue: 5, destTrue: 1, destFalse: 6},
		{items: []int{61, 70, 76, 69, 82, 56}, operation: plus(7), testValue: 3, destTrue: 5, destFalse: 2},
		{items: []int{79, 98}, operation: plus(2), testValue: 7, destTrue: 0, destFalse: 3},
		{items: []int{72, 79, 55}, operation: plus(1), testValue: 13, destTrue: 2, destFalse: 1},
		{items: []int{63}, operation: plus(4), testValue: 19, destTrue: 7, destFalse: 4},
		{items: []int{72, 51, 93, 63, 80, 86, 81}, operation: square(), testValue: 11, destTrue: 0, destFalse: 4},
	}

	return monkeys
}
