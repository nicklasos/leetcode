package main

import "fmt"

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	m := len(strs[0])

	for _, s := range strs {
		m = min(m, len(s))
	}

	prefix := ""

	for i := 0; i < m; i++ {
		prev := string(strs[0][i])

		for j := 1; j < len(strs); j++ {
			if prev != string(strs[j][i]) {
				prev = ""
				break
			}
		}

		if prev == "" {
			break
		}

		prefix += prev
	}

	return prefix
}

func main() {
	fmt.Println(longestCommonPrefix([]string{"flower", "flow", "flight"}))
	fmt.Println(longestCommonPrefix([]string{"dog", "racecar", "car"}))
	fmt.Println(longestCommonPrefix([]string{"cir", "car"}))
}
