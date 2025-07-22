package main

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 * 计算出旺仔哥哥最后会站在哪位小朋友旁边
 * @param a int整型一维数组 第 i 个小朋友的数字是 a_i
 * @param m int整型 表示旺仔哥哥的移动次数
 * @return int整型
 */
func stopAtWho(a []int, m int) int {
	cur := 0
	for i := 0; i < m; i++ {
		cur -= a[cur]
		cur %= len(a)
		if cur < 0 {
			cur += len(a)
		}
	}
	return cur + 1
}
