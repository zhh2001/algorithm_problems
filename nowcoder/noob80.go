package main

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 * 计算出这两个向量叉乘的结果
 * @param vector1 int整型一维数组
 * @param vector2 int整型一维数组
 * @return int整型一维数组
 */
func crossTimes(vector1 []int, vector2 []int) []int {
	x := vector1[1]*vector2[2] - vector1[2]*vector2[1]
	y := vector1[2]*vector2[0] - vector1[0]*vector2[2]
	z := vector1[0]*vector2[1] - vector1[1]*vector2[0]
	return []int{x, y, z}
}
