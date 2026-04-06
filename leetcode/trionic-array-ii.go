package main

import (
	"fmt"
	"math"
)

func maxSumTrionic(nums []int) int64 {
	var n = len(nums)
	// 最小有效长度：递增 2 + 递减 1 + 递增 1 = 4
	if n < 4 {
		return int64(math.MinInt64)
	}

	maxTriSum := int64(math.MinInt64)

	for i := 0; i < n; i++ {
		// 1. 寻找第一段：严格递增区间 [i, incEnd]
		incEnd := findIncreasingEnd(nums, i)
		if incEnd == i {
			continue
		}

		// 2. 寻找第二段：严格递减区间 [incEnd, decEnd]
		decEnd := findDecreasingEnd(nums, incEnd)
		if !isValidThirdSegment(nums, decEnd) {
			i = decEnd
			continue
		}

		// 3. 计算三段式子数组的总元素和
		currentSum := calculateTriSum(nums, i, incEnd, decEnd)
		if currentSum > maxTriSum {
			maxTriSum = currentSum
		}

		i = decEnd - 1
	}

	return maxTriSum
}

// findIncreasingEnd 从 start 开始，找严格递增的最后一个索引
func findIncreasingEnd(nums []int, start int) int {
	var j int
	var n = len(nums)
	for j = start + 1; j < n; j++ {
		if nums[j-1] >= nums[j] {
			break
		}
	}
	return j - 1
}

// findDecreasingEnd 从 start 开始，找严格递减的最后一个索引
func findDecreasingEnd(nums []int, start int) int {
	var j int
	var n = len(nums)
	for j = start + 1; j < n; j++ {
		if nums[j-1] <= nums[j] {
			break
		}
	}
	return j - 1
}

// isValidThirdSegment 验证是否存在有效的第三段（递增）
func isValidThirdSegment(nums []int, decEnd int) bool {
	n := len(nums)
	return decEnd != n-1 && nums[decEnd+1] > nums[decEnd]
}

// calculateTriSum 计算 递增->递减->递增 三段式的总元素和
func calculateTriSum(nums []int, start, incEnd, decEnd int) int64 {
	var sum int64

	// 累加第一段（递增）：末尾 2 个元素 + 前缀最大累加和
	sum += int64(nums[incEnd] + nums[incEnd-1])
	sum += maxIncreasingPrefixSum(nums, start, incEnd-2)

	// 累加第二段（递减）：所有元素
	for k := incEnd + 1; k <= decEnd; k++ {
		sum += int64(nums[k])
	}

	// 累加第三段（递增）：第一个元素 + 后缀最大累加和
	sum += int64(nums[decEnd+1])
	sum += maxIncreasingSuffixSum(nums, decEnd+2)

	return sum
}

// maxIncreasingPrefixSum 计算 [start, end] 连续递增的最大前缀累加和
func maxIncreasingPrefixSum(nums []int, start, end int) int64 {
	maxSum, currentSum := int64(0), int64(0)
	for k := end; k >= start; k-- {
		currentSum += int64(nums[k])
		if currentSum > maxSum {
			maxSum = currentSum
		}
	}
	return maxSum
}

// maxIncreasingSuffixSum 计算从 start 开始连续递增的最大后缀累加和
func maxIncreasingSuffixSum(nums []int, start int) int64 {
	var maxSum, currentSum int64
	var n = len(nums)
	for k := start; k < n; k++ {
		if nums[k] <= nums[k-1] {
			break
		}

		currentSum += int64(nums[k])
		if currentSum > maxSum {
			maxSum = currentSum
		}
	}
	return maxSum
}

func main() {
	var nums1 = []int{0, -2, -1, -3, 0, 2, -1}
	var maxTriSum1 = maxSumTrionic(nums1)
	fmt.Println(maxTriSum1)

	var nums2 = []int{1, 4, 2, 7}
	var maxTriSum2 = maxSumTrionic(nums2)
	fmt.Println(maxTriSum2)
}
