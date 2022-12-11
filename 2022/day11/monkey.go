package day11

type Monkey struct {
	items     []int
	operation action
	testValue int
	destTrue  int
	destFalse int
	done      int
}

func (m *Monkey) step(monkey []*Monkey) {
	for i := 0; i < len(m.items); i++ {
		newValue := m.operation(m.items[i])
		if newValue%m.testValue == 0 {
			monkey[m.destTrue].items = append(monkey[m.destTrue].items, newValue)
		} else {
			monkey[m.destFalse].items = append(monkey[m.destFalse].items, newValue)
		}
		m.done++
	}
	m.items = m.items[:0]
}
