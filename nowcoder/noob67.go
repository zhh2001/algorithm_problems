package main

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 * 计算 n 的阶乘
 * @param n int整型
 * @return int整型
 */
func factorialOfN(n int) int {
	result := 1
	for i := 1; i <= n; i++ {
		result *= i
		result %= 1e9 + 7
	}
	return result
}
