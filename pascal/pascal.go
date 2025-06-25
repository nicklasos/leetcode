package main

import "fmt"

func generate(numRows int) [][]int {
	rows := make([][]int, numRows)

	for i := 0; i < numRows; i++ {
		for j := 0; j <= i; j++ {

			n := 1
			if i-1 >= 0 && j >= 0 && j < len(rows[i-1]) {
				if i-1 >= 0 && j-1 >= 0 && j-1 < len(rows[i-1]) {
					n = rows[i-1][j] + rows[i-1][j-1]
				}
			}

			rows[i] = append(rows[i], n)
		}
	}

	return rows
}

func main() {

	fmt.Println(generate(5))
	fmt.Println(generate(1))

}
