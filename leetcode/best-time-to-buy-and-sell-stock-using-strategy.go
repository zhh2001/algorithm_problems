package main

import "fmt"

func maxProfit(prices []int, strategy []int, k int) int64 {
	profitSum := make([]int64, len(prices)+1)
	priceSum := make([]int64, len(prices)+1)
	for i := range prices {
		profitSum[i+1] = profitSum[i] + int64(prices[i])*int64(strategy[i])
		priceSum[i+1] = priceSum[i] + int64(prices[i])
	}
	res := profitSum[len(prices)]
	for i := k - 1; i < len(prices); i++ {
		leftProfit := profitSum[i-k+1]
		rightProfit := profitSum[len(prices)] - profitSum[i+1]
		changeProfit := priceSum[i+1] - priceSum[i-k/2+1]
		res = max(res, leftProfit+changeProfit+rightProfit)
	}
	return res
}

func main() {
	res1 := maxProfit([]int{4, 2, 8}, []int{-1, 0, 1}, 2)
	fmt.Println(res1)
	res2 := maxProfit([]int{5, 4, 3}, []int{1, 1, 0}, 2)
	fmt.Println(res2)
}
