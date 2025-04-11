func twoSum(nums []int, target int) []int {
    mapping := map[int]int{}
    for i, n := range nums {
        if p, ok := mapping[target-n]; ok {
            return []int{p, i}
        }
        mapping[n] = i
    }
    return nil
}
