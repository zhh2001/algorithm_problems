package main

import "fmt"

func minCost(startPos, homePos, rowCosts, colCosts []int) int {
	sx, sy, ex, ey := startPos[0], startPos[1], homePos[0], homePos[1]

	// 计算行总代价 + 列总代价
	return calcCost(sx, ex, rowCosts) + calcCost(sy, ey, colCosts)
}

// 起点到终点的代价计算
func calcCost(start, end int, costs []int) int {
	sum := 0
	for _, v := range costs[min(start, end) : max(start, end)+1] {
		sum += v
	}
	return sum - costs[start] // 减去起点费用
}

func main() {
	startPos := []int{1, 0}
	homePos := []int{2, 3}
	rowCosts := []int{5, 4, 3}
	colCosts := []int{8, 2, 6, 7}
	cost := minCost(startPos, homePos, rowCosts, colCosts)
	fmt.Println(cost)

	startPos = []int{0, 0}
	homePos = []int{0, 0}
	rowCosts = []int{5}
	colCosts = []int{26}
	cost = minCost(startPos, homePos, rowCosts, colCosts)
	fmt.Println(cost)
}
