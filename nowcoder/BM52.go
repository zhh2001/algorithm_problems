package main

import (
	"fmt"
	"sort"
)

/**
 *
 * @param nums int整型一维数组
 * @return int整型一维数组
 */
func FindNumsAppearOnce(nums []int) []int {
	set := make(map[int]struct{})
	for _, v := range nums {
		_, ok := set[v]
		if ok {
			delete(set, v)
		} else {
			set[v] = struct{}{}
		}
	}

	slice := make([]int, 0, len(set))
	for k := range set {
		slice = append(slice, k)
	}
	sort.Ints(slice)
	return slice
}

func test1() {
	numbers := []int{1, 4, 1, 6}
	fmt.Println("=====================")
	fmt.Println("Input:", numbers)
	fmt.Println("Output:", FindNumsAppearOnce(numbers))
}

func test2() {
	numbers := []int{1, 2, 3, 3, 2, 9}
	fmt.Println("=====================")
	fmt.Println("Input:", numbers)
	fmt.Println("Output:", FindNumsAppearOnce(numbers))
}

func main() {
	test1()
	test2()
}
