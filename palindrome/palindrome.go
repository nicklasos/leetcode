package main

import (
	"fmt"
	"strconv"
)

func isPalindrome(x int) bool {

	strX := strconv.Itoa(x)

	r := []rune(strX)

	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		if r[i] != r[j] {
			return false
		}
	}

	return true
}

func main() {
	n := -121

	r := isPalindrome(n)

	fmt.Println(r)
}
