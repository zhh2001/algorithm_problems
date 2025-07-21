package main

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 * 求序列a中的峰、谷点的个数
 * @param a int整型一维数组 序列a
 * @return int整型
 */
func countPeakPoint(a []int) int {
	count := 0
	for i := 1; i < len(a)-1; i++ {
		if a[i] > a[i+1] && a[i] > a[i-1] {
			count++
		} else if a[i] < a[i+1] && a[i] < a[i-1] {
			count++
		}
	}
	return count
}
