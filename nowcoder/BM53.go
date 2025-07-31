package main

import (
	"fmt"
)

/**
 *
 * @param nums int整型一维数组
 * @return int整型
 */
func minNumberDisappeared(nums []int) int {
	set := make(map[int]struct{})
	for _, v := range nums {
		set[v] = struct{}{}
	}

	var n int
	for i := 1; ; i++ {
		if _, ok := set[i]; !ok {
			n = i
			break
		}
	}
	return n
}

func test1() {
	nums := []int{1, 0, 2}
	fmt.Println("=====================")
	fmt.Println("Input:", nums)
	fmt.Println("Output:", minNumberDisappeared(nums))
}

func test2() {
	nums := []int{-2, 3, 4, 1, 5}
	fmt.Println("=====================")
	fmt.Println("Input:", nums)
	fmt.Println("Output:", minNumberDisappeared(nums))
}

func test3() {
	nums := []int{4, 5, 6, 8, 9}
	fmt.Println("=====================")
	fmt.Println("Input:", nums)
	fmt.Println("Output:", minNumberDisappeared(nums))
}

func main() {
	test1()
	test2()
	test3()
}
