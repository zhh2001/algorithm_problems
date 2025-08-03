package main

import (
	"fmt"
	"sort"
)

/**
 *
 * @param nums1 int整型一维数组
 * @param nums2 int整型一维数组
 * @return int整型一维数组
 */
func intersection(nums1 []int, nums2 []int) []int {
	map1 := map[int]struct{}{}
	map2 := map[int]struct{}{}

	for _, num := range nums1 {
		map1[num] = struct{}{}
	}

	for _, num := range nums2 {
		map2[num] = struct{}{}
	}

	res := make([]int, 0, len(map1))
	for k := range map1 {
		if _, ok := map2[k]; ok {
			res = append(res, k)
		}
	}

	sort.Ints(res)
	return res
}

func main() {
	nums1 := []int{1, 1, 4}
	nums2 := []int{5, 1, 4}
	res := intersection(nums1, nums2)
	fmt.Println(res)
}
