package stacks

import "strings"

func removeDuplicates(input string) string {
	stack := NewStack[rune](len(input))
	for _, c := range input {
		top, _ := stack.Peek()
		if top == c {
			stack.Pop()
			continue
		} else {
			stack.Push(c)
		}
	}
	var sb strings.Builder
	for _, c := range stack.AsList() {
		sb.WriteRune(c)
	}
	return sb.String()
}
