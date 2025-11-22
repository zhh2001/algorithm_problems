package main

import (
	"fmt"
	"strconv"
)

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 *
 * @param x int整型
 * @return bool布尔型
 */
func isPalindrome(x int) bool {
	// write code here
	s := strconv.Itoa(x)
	i, j := 0, len(s)-1
	for i < j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}

func main() {
	fmt.Println(isPalindrome(121))
	fmt.Println(isPalindrome(1221))
}
