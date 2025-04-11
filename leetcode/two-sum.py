class Solution:
    def twoSum(self, nums: List[int], target: int) -> List[int]:
        mapping: Dict[int, int] = {}
        for i, n in enumerate(nums):
            m = target - n
            if m in mapping:
                return mapping[m], i
            mapping[n] = i
