from typing import List


class Solution:
    def minBitwiseArray(self, nums: List[int]) -> List[int]:
        return [n ^ ((n + 1) & ~n) >> 1 if n != 2 else -1 for n in nums]


if __name__ == "__main__":
    solution = Solution()
    print(solution.minBitwiseArray([2, 3, 5, 7]))
    print(solution.minBitwiseArray([11, 13, 31]))
