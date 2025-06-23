package main

import (
	"fmt"
	"strings"
)

func strStr(haystack string, needle string) int {
	return strings.Index(haystack, needle)
}

func main() {
	fmt.Println(strStr("leetcode", "leeto"))
	fmt.Println(strStr("sadbutsad", "sad"))
}
