package main

import "fmt"

type Stack struct {
	items []int
}

func (s *Stack) Push(value int) {
	s.items = append(s.items, value)
}

func (s *Stack) Pop() (int, bool) {
	if len(s.items) == 0 {
		return 0, false
	}

	last := len(s.items) - 1
	value := s.items[last]
	s.items = s.items[:last]
	return value, true
}

func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0
}

func main() {
	stack := &Stack{}
	stack.Push(10)
	stack.Push(20)
	stack.Push(30)

	for !stack.IsEmpty() {
		value, _ := stack.Pop()
		fmt.Println(value)
	}
}
