package main

import (
	"fmt"
	"math"
)

func largestSquareArea(bottomLeft [][]int, topRight [][]int) int64 {
	maxSide := 0
	for i, b1 := range bottomLeft {
		t1 := topRight[i]
		for j, b2 := range bottomLeft[:i] {
			t2 := topRight[j]
			side := min(min(t1[1], t2[1])-max(b1[1], b2[1]), min(t1[0], t2[0])-max(b1[0], b2[0]))
			maxSide = max(maxSide, side)
		}
	}
	return int64(math.Pow(float64(maxSide), 2))
}

func main() {
	var bottomLeft, topRight [][]int
	bottomLeft = [][]int{{1, 1}, {2, 2}, {3, 1}}
	topRight = [][]int{{3, 3}, {4, 4}, {6, 6}}
	fmt.Println(largestSquareArea(bottomLeft, topRight))
	bottomLeft = [][]int{{1, 1}, {1, 3}, {1, 5}}
	topRight = [][]int{{5, 5}, {5, 7}, {5, 9}}
	fmt.Println(largestSquareArea(bottomLeft, topRight))
	bottomLeft = [][]int{{1, 1}, {2, 2}, {1, 2}}
	topRight = [][]int{{3, 3}, {4, 4}, {3, 4}}
	fmt.Println(largestSquareArea(bottomLeft, topRight))
	bottomLeft = [][]int{{1, 1}, {3, 3}, {3, 1}}
	topRight = [][]int{{2, 2}, {4, 4}, {4, 2}}
	fmt.Println(largestSquareArea(bottomLeft, topRight))
}
