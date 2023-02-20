package main

import "fmt"

func isValid(s string) bool {
	if len(s) <= 1 {
		return false
	}
	m := map[rune]rune{
		'(': ')',
		'[': ']',
		'{': '}',
	}
	var stack []rune

	for _, v := range s {
		if v == '(' || v == '[' || v == '{' {

			stack = append(stack, v)
		} else {

			lastIndex := len(stack) - 1
			if lastIndex == -1 {
				return false
			}
			lastOpenParenthese := stack[lastIndex]
			if v == m[lastOpenParenthese] {
				stack = stack[:lastIndex]
				fmt.Println("stack", stack)
			} else {
				return false
			}
		}

	}

	return len(stack) == 0
}

func main() {

	fmt.Println(isValid("({){}}"))

}
