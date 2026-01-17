package main

import "fmt"

func maximalRectangle(matrix [][]byte) int {
	ans := 0
	if len(matrix) == 0 {
		return ans
	}
	left := make([][]int, len(matrix))
	for i, row := range matrix {
		left[i] = make([]int, len(matrix[0]))
		for j, v := range row {
			if v == '0' {
				continue
			}
			if j == 0 {
				left[i][j] = 1
			} else {
				left[i][j] = left[i][j-1] + 1
			}
		}
	}
	for i, row := range matrix {
		for j, v := range row {
			if v == '0' {
				continue
			}
			width := left[i][j]
			area := width
			for k := i - 1; k >= 0; k-- {
				width = min(width, left[k][j])
				area = max(area, (i-k+1)*width)
			}
			ans = max(ans, area)
		}
	}
	return ans
}

func main() {
	matrix1 := [][]byte{
		{'1', '0', '1', '0', '0'},
		{'1', '0', '1', '1', '1'},
		{'1', '1', '1', '1', '1'},
		{'1', '0', '0', '1', '0'},
	}
	fmt.Println(maximalRectangle(matrix1))

	matrix2 := [][]byte{
		{'0'},
	}
	fmt.Println(maximalRectangle(matrix2))

	matrix3 := [][]byte{
		{'1'},
	}
	fmt.Println(maximalRectangle(matrix3))
}
