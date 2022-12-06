package day5

// Crate is a stack of strings.
type Crate []string

// Push adds a string to the top of the stack.
func (s *Crate) Push(str string) {
	*s = append(*s, str)
}

// Pop removes the top string from the stack and returns it.
// If the stack is empty, it returns an empty string.
func (s *Crate) Pop() string {
	if len(*s) == 0 {
		return ""
	}
	str := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return str
}
