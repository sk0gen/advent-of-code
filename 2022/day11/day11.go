package day11

import (
	"fmt"
	"sort"
)

func TaskA(monkeys []*Monkey) int {

	for i := 0; i < 20; i++ {
		for _, monkey := range monkeys {
			monkey.step(monkeys)
		}
	}
	work := make([]int, len(monkeys))
	for i, m := range monkeys {
		work[i] = m.done
	}
	sort.Ints(work)
	return work[len(work)-1] * work[len(work)-2]
}

func TaskB(monkeys []*Monkey) int {
	for i := 0; i < 1000; i++ {
		for _, monkey := range monkeys {
			monkey.step(monkeys)
		}
	}
	work := make([]int, len(monkeys))
	for i, m := range monkeys {
		work[i] = m.done
	}
	sort.Ints(work)
	fmt.Println(work[len(work)-1])
	fmt.Println(work[len(work)-2])
	return work[len(work)-1] * work[len(work)-2]
}

func maxValue(a, b int) int {
	if a > b {
		return a
	}
	return b
}
