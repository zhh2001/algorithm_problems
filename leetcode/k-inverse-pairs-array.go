package main

import "fmt"

const MOD = 1e9 + 7

func kInversePairs(n int, k int) int {
	var dp = make([][]int, n+1)
	var preSum = make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, k+1)
		preSum[i] = make([]int, k+1)
	}

	dp[1][0] = 1
	for j := 0; j <= k; j++ {
		preSum[1][j] = 1
	}

	for i := 2; i <= n; i++ {
		for j := 0; j <= k; j++ {
			if j < i {
				dp[i][j] = preSum[i-1][j]
			} else {
				dp[i][j] = (preSum[i-1][j] - preSum[i-1][j-i] + MOD) % MOD
			}

			if j == 0 {
				preSum[i][j] = dp[i][j]
			} else {
				preSum[i][j] = (preSum[i][j-1] + dp[i][j]) % MOD
			}
		}
	}

	return dp[n][k]
}

func main() {
	var n1 = 3
	var k1 = 0
	var ans1 = kInversePairs(n1, k1)
	fmt.Println(ans1)

	var n2 = 3
	var k2 = 1
	var ans2 = kInversePairs(n2, k2)
	fmt.Println(ans2)
}
