package stacks

// Given a string that may consist of opening and closing parentheses, your task is to check whether or not the string contains valid parenthesization.

// The conditions to validate are as follows:

// Every opening parenthesis should be closed by the same kind of parenthesis. Therefore, {)and [(]) strings are invalid.

// Every opening parenthesis must be closed in the correct order. Therefore, )( and ()(() are invalid.

func isValidParenthesis(input string) bool {
	s := NewStack[byte](len(input))
	inputByte := []byte(input)
	for _, b := range inputByte {
		if b == '(' || b == '[' || b == '{' {
			s.Push(b)
		} else {
			top, _ := s.Peek()
			if top == '(' && b == ')' {
				s.Pop()
			} else if top == '{' && b == '}' {
				s.Pop()
			} else if top == '[' && b == ']' {
				s.Pop()
			} else {
				return false
			}
		}
	}
	return s.IsEmpty()
}
