package main

import "fmt"

func repeatedNTimes(nums []int) int {
	set := map[int]struct{}{}
	for _, n := range nums {
		if _, ok := set[n]; ok {
			return n
		}
		set[n] = struct{}{}
	}
	return -1
}

func main() {
	fmt.Println(repeatedNTimes([]int{1, 2, 3, 3}))
	fmt.Println(repeatedNTimes([]int{2, 1, 2, 5, 3, 2}))
	fmt.Println(repeatedNTimes([]int{5, 1, 5, 2, 5, 3, 5, 4}))
}
