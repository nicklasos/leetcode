package main

import (
	"fmt"
)

func romanToInt(s string) int {
	nums := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,

		"IV": 3,
		"IX": 8,
		"XL": 30,
		"XC": 80,
		"CD": 300,
		"CM": 800,
	}

	// Exceptions
	// IV - 4 , IX - 9
	// XL - 40 , XC - 90
	// CD - 400 , CM - 900

	sum := 0

	for i, n := range s {
		key := string(n)

		if i != 0 && nums[string(s[i-1])+string(n)] != 0 {
			key = string(s[i-1]) + string(n)
		}

		sum += nums[key]
	}

	return sum
}

func main() {
	fmt.Println(romanToInt("IV"))
	fmt.Println(romanToInt("III"))
	fmt.Println(romanToInt("LVIII"))
	fmt.Println(romanToInt("MCMXCIV"))
}
