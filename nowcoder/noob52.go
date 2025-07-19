package main

import (
	"fmt"
	"strconv"
)

func main() {
	var n, m int
	_, err := fmt.Scanf("%d %d\n", &n, &m)
	if err != nil {
		panic(err)
	}

	matrix := make([][]byte, n)
	for i := 0; i < n; i++ {
		var str string
		_, err = fmt.Scanf("%s\n", &str)
		if err != nil {
			panic(err)
		}
		matrix[i] = []byte(str)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if matrix[i][j] == '.' {
				cnt := 0
				for x := i - 1; x <= i+1; x++ {
					for y := j - 1; y <= j+1; y++ {
						if x == i && y == j {
							continue
						}
						if x < 0 || x >= n || y < 0 || y >= m {
							continue
						}
						if matrix[x][y] == '*' {
							cnt++
						}
					}
				}
				matrix[i][j] = strconv.Itoa(cnt)[0]
			}
			fmt.Printf("%c", matrix[i][j])
		}
		fmt.Println()
	}
}
