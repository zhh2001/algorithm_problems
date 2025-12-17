package main

import "fmt"

func maximumProfit(prices []int, k int) int64 {
	dp := make([][][]int64, len(prices))
	for i := range dp {
		dp[i] = make([][]int64, k+1)
		for j := range dp[i] {
			dp[i][j] = make([]int64, 3)
		}
	}

	for j := 1; j <= k; j++ {
		dp[0][j][1] = int64(-prices[0])
		dp[0][j][2] = int64(prices[0])
	}

	for i := 1; i < len(prices); i++ {
		for j := 1; j <= k; j++ {
			dp[i][j][0] = max(dp[i-1][j][0], max(dp[i-1][j][1]+int64(prices[i]), dp[i-1][j][2]-int64(prices[i])))
			dp[i][j][1] = max(dp[i-1][j][1], dp[i-1][j-1][0]-int64(prices[i]))
			dp[i][j][2] = max(dp[i-1][j][2], dp[i-1][j-1][0]+int64(prices[i]))
		}
	}

	return dp[len(prices)-1][k][0]
}

func main() {
	res1 := maximumProfit([]int{1, 7, 9, 8, 2}, 2)
	fmt.Println(res1)

	res2 := maximumProfit([]int{12, 16, 19, 19, 8, 1, 19, 13, 9}, 3)
	fmt.Println(res2)
}
