package day7

import (
	"os"
	"strings"
	"testing"
)

func TestTaskATestInput(t *testing.T) {
	result := TaskA(read("input_test.txt"))

	expected := 95437

	if result != expected {
		t.Errorf("got %d, wanted %d", result, expected)
	}
}

func TestTaskATaskInput(t *testing.T) {

	result := TaskA(read("input.txt"))

	expected := 1500

	if result != expected {
		t.Errorf("got %d, wanted %d", result, expected)
	}
}

func TestTaskBTestInput(t *testing.T) {
	result := TaskB(read("input_test.txt"))

	expected := 24933642

	if result != expected {
		t.Errorf("got %d, wanted %d", result, expected)
	}
}

func TestTaskBTaskInput(t *testing.T) {
	result := TaskB(read("input.txt"))

	expected := 7421137

	if result != expected {
		t.Errorf("got %d, wanted %d", result, expected)
	}
}

func read(fileName string) []string {
	file, _ := os.ReadFile(fileName)

	result := strings.Split(string(file), "\n$ ")
	result[0] = strings.Replace(result[0], "$ ", "", -1)

	return result
}
