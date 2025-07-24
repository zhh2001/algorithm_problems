package main

import (
	"fmt"
	"strconv"
	"strings"
)

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 * 给定一个后缀表达式，返回它的结果
 * @param str string字符串
 * @return long长整型
 */
func legalExp(str string) int64 {
	chars := make([]int64, 0, len(str))
	for i := 0; i < len(str); i++ {
		if str[i] >= '0' && str[i] <= '9' {
			var builder strings.Builder
			for ; str[i] != '#'; i++ {
				builder.WriteByte(str[i])
			}
			digit := builder.String()
			n, err := strconv.Atoi(digit)
			if err != nil {
				panic(err)
			}
			chars = append(chars, int64(n))
			continue
		}

		var res int64
		switch str[i] {
		case '*':
			res = chars[len(chars)-2] * chars[len(chars)-1]
		case '+':
			res = chars[len(chars)-2] + chars[len(chars)-1]
		case '-':
			res = chars[len(chars)-2] - chars[len(chars)-1]
		}
		fmt.Println(res)
		chars = append(chars[:len(chars)-2], res)
	}
	return chars[0]
}
