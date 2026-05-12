package main

import "fmt"

type Stack struct {
	items []string
}

func (s *Stack) Push(value string) {
	s.items = append(s.items, value)
}

func (s *Stack) Pop() (string, bool) {
	if len(s.items) == 0 {
		return "", false
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
	browserHistory := &Stack{}

	browserHistory.Push("google.com")
	browserHistory.Push("github.com")
	browserHistory.Push("stackoverflow.com")

	fmt.Println("User presses back button:")
	for !browserHistory.IsEmpty() {
		page, _ := browserHistory.Pop()
		fmt.Println("Back to:", page)
	}
}
