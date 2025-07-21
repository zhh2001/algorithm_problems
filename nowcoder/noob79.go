package main

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 * 计算两个三维向量的点乘结果
 * @param vector1 int整型一维数组 第一个向量
 * @param vector2 int整型一维数组 第二个向量
 * @return int整型
 */
func dotTime(vector1 []int, vector2 []int) int {
	result := 0
	for i := 0; i < len(vector1); i++ {
		result += vector1[i] * vector2[i]
	}
	return result
}
