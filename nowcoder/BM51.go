package main

import "fmt"

/**
 *
 * @param numbers int整型一维数组
 * @return int整型
 */
func MoreThanHalfNum_Solution(numbers []int) int {
	count := make(map[int]int)
	for _, v := range numbers {
		count[v]++
	}

	halfLen := len(numbers) / 2
	var result int
	for k, v := range count {
		if v > halfLen {
			result = k
			break
		}
	}
	return result
}

func test1() {
	numbers := []int{1, 2, 3, 2, 2, 2, 5, 4, 2}
	fmt.Println(MoreThanHalfNum_Solution(numbers))
}

func test2() {
	numbers := []int{3, 3, 3, 3, 2, 2, 2}
	fmt.Println(MoreThanHalfNum_Solution(numbers))
}

func test3() {
	numbers := []int{1}
	fmt.Println(MoreThanHalfNum_Solution(numbers))
}

func main() {
	test1()
	test2()
	test3()
}
