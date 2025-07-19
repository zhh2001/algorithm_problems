package main

import (
	"fmt"
)

func main() {
	var n, m int
	var err error
	_, err = fmt.Scanf("%d %d\n", &n, &m)
	if err != nil {
		panic(err)
	}

	matrix := make([][]uint, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]uint, m)
		for j := 0; j < m; j++ {
			if i == 0 || j == 0 {
				matrix[i][j] = 1
				continue
			}
			matrix[i][j] = (matrix[i-1][j] + matrix[i][j-1]) % (1e9 + 7)
		}
	}
	fmt.Println(matrix[n-1][m-1])
}
