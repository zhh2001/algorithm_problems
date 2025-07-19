package main

import (
	"fmt"
)

// Pascal 生成杨辉三角
func Pascal(n int) [][]int {
	triangle := make([][]int, n)
	for i := 0; i < n; i++ {
		triangle[i] = make([]int, i+1)
		for j := 0; j <= i; j++ {
			if j == 0 || i == j {
				triangle[i][j] = 1
			} else {
				triangle[i][j] = triangle[i-1][j-1] + triangle[i-1][j]
			}
		}
	}
	return triangle
}

func main() {
	var n int
	_, err := fmt.Scanf("%d\n", &n)
	if err != nil {
		panic(err)
	}

	triangle := Pascal(n)
	for i := 0; i < n; i++ {
		for j := 0; j <= i; j++ {
			fmt.Printf("%d ", triangle[i][j])
		}
		fmt.Println()
	}
}
