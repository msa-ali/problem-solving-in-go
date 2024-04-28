package stacks

import (
	"strconv"
	"unicode"
)

func Calculator(expression string) int {
	result, current, sign := 0, 0, 1
	stack := NewStack[int](2)
	for _, c := range expression {
		if unicode.IsDigit(c) {
			currentDigit, _ := strconv.Atoi(string(c))
			current = current*10 + currentDigit
			continue
		}
		if c == '+' || c == '-' {
			result += current * sign
			current = 0
			if c == '+' {
				sign = 1
			} else {
				sign = -1
			}
			continue
		}
		if c == '(' {
			stack.Push(result)
			stack.Push(sign)
			result = 0
			sign = 1
			continue
		} else if c == ')' {
			result += current * sign
			prevSign, _ := stack.Pop()
			prevResult, _ := stack.Pop()
			result += prevResult * prevSign
			current = 0
			continue
		}
	}
	return result + (current * sign)
}
