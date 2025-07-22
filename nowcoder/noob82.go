package main

import (
	"strconv"
	"strings"
)

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 * 整理出一个将序列中的数字以逗号隔开从而得到的字符串
 * @param a int整型一维数组 需要整理的序列 a
 * @return string字符串
 */
func commaTransformer(a []int) string {
	builder := strings.Builder{}
	for i, v := range a {
		builder.WriteString(strconv.Itoa(v))
		if i != len(a)-1 {
			builder.WriteString(",")
		}
	}
	return builder.String()
}
