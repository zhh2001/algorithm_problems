package main

import (
	"fmt"
	"strconv"
)

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 *
 * @param a string字符串
 * @param b string字符串
 * @return string字符串
 */
func sum(a string, b string) string {
	// write code here
	aInt, err := strconv.ParseInt(a, 10, 64)
	if err != nil {
		panic(err)
	}
	bInt, err := strconv.ParseInt(b, 10, 64)
	if err != nil {
		panic(err)
	}
	return strconv.Itoa(int(aInt + bInt))
}

func main() {
	fmt.Println(sum("12", "34"))
}
