package stacks

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStacks(t *testing.T) {
	s := NewStack[int](5)
	assert.Equal(t, true, s.IsEmpty())
	first, err := s.Peek()
	assert.Equal(t, 0, first)
	assert.EqualError(t, err, "stack is empty")
	s.Push(1)
	s.Push(2)
	assert.Equal(t, 2, s.Size())
	assert.Equal(t, false, s.IsFull())
	s.Push(3)
	s.Push(4)
	s.Push(5)
	assert.Equal(t, true, s.IsFull())
	assert.Equal(t, 5, s.Size())
	err = s.Push(6)
	assert.EqualError(t, err, "stack is full")
	res, err := s.Pop()
	assert.Equal(t, 5, res)
	assert.Equal(t, 4, s.Size())
	assert.Equal(t, nil, err)
	res, err = s.Pop()
	assert.Equal(t, 4, res)
	assert.Equal(t, 3, s.Size())
	assert.Equal(t, nil, err)
	s.Pop()
	s.Pop()
	s.Pop()
	assert.Equal(t, 0, s.Size())
	_, err = s.Pop()
	assert.EqualError(t, err, "stack is empty")
}
