from heapq import nsmallest
from typing import List


class Solution:
    def minimumCost(self, nums: List[int]) -> int:
        return sum(nsmallest(2, nums[1:]) + nums[:1])


if __name__ == "__main__":
    solution = Solution()
    print(solution.minimumCost([1, 2, 3, 12]))
    print(solution.minimumCost([5, 4, 3]))
    print(solution.minimumCost([10, 3, 1, 1]))
