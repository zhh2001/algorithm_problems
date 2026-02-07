from typing import Any, List


class Solution:
    def constructTransformedArray(self, nums: List[int]) -> List[int]:
        return [nums[(i + n) % len(nums)] for i, n in enumerate[Any](nums)]


if __name__ == "__main__":
    solution = Solution()
    print(solution.constructTransformedArray([3, -2, 1, 1]))
    print(solution.constructTransformedArray([-1, 4, -1]))
