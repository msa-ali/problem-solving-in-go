package stacks

import "fmt"

type Stack[T any] struct {
	capacity  int
	container []T
}

func NewStack[T any](size int) Stack[T] {
	return Stack[T]{
		capacity:  size,
		container: make([]T, 0, size),
	}
}

func (s *Stack[T]) Size() int {
	return len(s.container)
}

func (s *Stack[T]) IsEmpty() bool {
	return s.Size() == 0
}

func (s *Stack[T]) IsFull() bool {
	return s.Size() == s.capacity
}

func (s *Stack[T]) Peek() (T, error) {
	if s.IsEmpty() {
		var res T
		return res, fmt.Errorf("stack is empty")
	}
	return s.container[s.Size()-1], nil
}

func (s *Stack[T]) Push(value T) error {
	if s.IsFull() {
		return fmt.Errorf("stack is full")
	}
	s.container = append(s.container, value)
	return nil
}

func (s *Stack[T]) Pop() (T, error) {
	var result T
	if s.IsEmpty() {
		return result, fmt.Errorf("stack is empty")
	}
	result, _ = s.Peek()
	s.container = s.container[:s.Size()-1]
	return result, nil
}

func (s *Stack[T]) AsList() (res []T) {
	res = make([]T, s.Size())
	copy(res, s.container)
	return
}

func (s *Stack[T]) Print() {
	fmt.Println(s.container)
}
