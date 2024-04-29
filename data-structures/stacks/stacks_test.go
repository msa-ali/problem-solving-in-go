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
	assert.Equal(t, []int{1, 2, 3, 4, 5}, s.AsList())
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

func TestCalculator(t *testing.T) {
	result := Calculator("4+(15-12)+100")
	assert.Equal(t, 107, result)
}

func TestRemoveDuplicates(t *testing.T) {
	assert.Equal(t, "ay", removeDuplicates("azxxzy"))
	assert.Equal(t, "abcde", removeDuplicates("abcde"))
	assert.Equal(t, "", removeDuplicates("aabbccdd"))
}

func TestMinRemoveParentheses(t *testing.T) {
	assert.Equal(t, "(abc)(to)(q)()", minRemoveParentheses("(((abc)(to)((q)()("))
}

func TestQueueFromStack(t *testing.T) {
	q := NewQueue(10)
	q.push(10)
	q.push(20)
	assert.Equal(t, 10, q.peek())
	assert.Equal(t, 10, q.pop())
	assert.Equal(t, 20, q.pop())
	assert.Equal(t, true, q.empty())
}

func TestIsValidParenthesis(t *testing.T) {
	assert.Equal(t, true, isValidParenthesis("(){}[]"))
	assert.Equal(t, false, isValidParenthesis("{}[]{}[{}])"))
	assert.Equal(t, true, isValidParenthesis("(){[{()}]}"))
	assert.Equal(t, false, isValidParenthesis("))){{}}}]]"))
	assert.Equal(t, false, isValidParenthesis(")))))"))
	assert.Equal(t, true, isValidParenthesis("()"))
}
