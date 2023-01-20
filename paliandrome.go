package main

import (
	"fmt"
	"strconv"
)

func isPalindrome1(x int) bool {
	if x < 0 {
		return false
	}

	conv := strconv.Itoa(x)
	for i, j := 0, len(conv)-1; i < j; i, j = i+1, j-1 {
		fmt.Println("i,j", i, j)
		if conv[i] != conv[j] {
			return false
		}
	}
	return true

}

func isPalindrome(x int) bool {

	if x < 0 {
		return false
	} else if x >= 0 {
		conv := strconv.Itoa(x)
		if Reverse(conv) == conv {
			return true
		} else if len(conv) == 1 {
			return true

		} else {
			return false
		}

	}
	return false

}
func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func main() {

	var angka int = 142241

	fmt.Println(isPalindrome1(angka))

}
