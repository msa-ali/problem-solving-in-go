package stacks

import "strings"

type pair struct {
	char  byte
	index int
}

// (((abc)(to)((q)()( -> (abc)(to)(q)()
func minRemoveParentheses(s string) string {
	stack := NewStack[pair](len(s))
	sList := []byte(s)
	// var res string
	for i, c := range sList {
		top, _ := stack.Peek()
		if c == ')' && !stack.IsEmpty() && top.char == '(' {
			stack.Pop()
			continue
		}
		if c == '(' || c == ')' {
			stack.Push(pair{char: c, index: i})
		}
	}

	for !stack.IsEmpty() {
		pair, _ := stack.Pop()
		sList[pair.index] = ' '
	}
	var res strings.Builder
	for _, c := range sList {
		if c != ' ' {
			res.WriteByte(c)
		}
	}
	return res.String()
}
