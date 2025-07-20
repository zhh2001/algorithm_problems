package main

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 * 判断二元一次方程组是否有解
 * @param a int整型 二次项系数
 * @param b int整型 一次项系数
 * @param c int整型 常数项
 * @return bool布尔型
 */
func judgeSolutions(a int, b int, c int) bool {
	delta := b*b - a*c*4
	return delta >= 0
}
