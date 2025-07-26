package main

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 *
 * @param nums int整型一维数组
 * @param target int整型
 * @return int整型一维数组
 */
func twoSum(nums []int, target int) []int {
	mapping := make(map[int]int)
	for index, num := range nums {
		key := target - num
		value, ok := mapping[key]
		if !ok {
			mapping[num] = index
			continue
		}
		return []int{value + 1, index + 1}
	}
	return nil
}
