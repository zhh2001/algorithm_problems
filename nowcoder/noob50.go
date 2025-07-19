package main

import (
	"fmt"
)

func Transpose(matrix [][]int) [][]int {
	n := len(matrix)
	m := len(matrix[0])
	newMatrix := make([][]int, m)
	for i := 0; i < m; i++ {
		newMatrix[i] = make([]int, n)
		for j := 0; j < n; j++ {
			newMatrix[i][j] = matrix[j][i]
		}
	}
	return newMatrix
}

func main() {
	var n, m int
	_, err := fmt.Scanf("%d %d\n", &n, &m)
	if err != nil {
		panic(err)
	}

	matrix := make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, m)
		for j := 0; j < m; j++ {
			_, err = fmt.Scan(&matrix[i][j])
			if err != nil {
				panic(err)
			}
		}
	}

	newMatrix := Transpose(matrix)
	for _, row := range newMatrix {
		for _, v := range row {
			fmt.Printf("%d ", v)
		}
		fmt.Println()
	}
}
