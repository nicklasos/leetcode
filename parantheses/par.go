package main

import (
	"fmt"
	"slices"
)

func isClosingBracket(b byte) bool {
	closing := []byte{']', '}', ')'}

	return slices.Contains(closing, b)
}

func isSame(a, b byte) bool {
	if (a == '[' || a == ']') && (b == '[' || b == ']') {
		return true
	}

	if (a == '(' || a == ')') && (b == '(' || b == ')') {
		return true
	}

	if (a == '{' || a == '}') && (b == '{' || b == '}') {
		return true
	}

	return false
}

func isValid(s string) bool {
	if len(s)%2 != 0 {
		return false
	}

	if !isClosingBracket(s[len(s)-1]) {
		return false
	}

	if isClosingBracket(s[0]) {
		return false
	}

	stack := []byte{}

	for _, ch := range s {

		if isClosingBracket(byte(ch)) {
			if len(stack) == 0 {
				return false
			}

			prev := stack[len(stack)-1]

			if isClosingBracket(prev) || !isSame(prev, byte(ch)) {
				return false
			}

			stack = stack[:len(stack)-1]

		} else {
			stack = append(stack, byte(ch))
		}
	}

	if len(stack) != 0 {
		return false
	}

	return true
}

func main() {
	fmt.Println(isValid("()))"))
	fmt.Println(isValid("[[[]"))
	fmt.Println(isValid("([]"))
	fmt.Println(isValid("()"))
	fmt.Println(isValid("(]"))
	fmt.Println(isValid("()[]{}"))
	fmt.Println(isValid("([])"))
}
