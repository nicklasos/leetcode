package main

import (
	"strings"
)

func lengthOfLastWord(s string) int {
	words := strings.Fields(s)
	last := words[len(words)-1]

	return len(last)
}

func main() {
	s := "   fly me   to   the moon  "

	r := lengthOfLastWord(s)

	println(r)
}
