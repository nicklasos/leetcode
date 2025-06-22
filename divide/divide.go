package main

import (
	"fmt"
)

func divideString(s string, k int, fill byte) []string {

	l := (len(s) + k - 1) / k
	result := make([]string, l)

	for index, ch := range s {
		result[index/k] += string(ch)
	}

	needed := k - len(result[len(result)-1])

	if needed != k {
		for i := 0; i < needed; i++ {
			result[len(result)-1] += string(fill)
		}
	}

	return result
}

func main() {
	fmt.Println(divideString("abcdefghi", 3, 'x'))
	fmt.Println(divideString("abcdefghij", 3, 'x'))
}
