from typing import List


class Solution:
    def findMaxConsecutiveOnes(self, nums: List[int]) -> int:
        return max(len(s) for s in "".join(map(str, nums)).split("0"))


if __name__ == '__main__':
    solution = Solution()
    print(solution.findMaxConsecutiveOnes([1, 1, 0, 1, 1, 1]))
    print(solution.findMaxConsecutiveOnes([1, 0, 1, 1, 0, 1]))
