from bisect import bisect_right
from typing import List


class Solution:
    def nextGreatestLetter(self, letters: List[str], target: str) -> str:
        return letters[bisect_right(letters, target) % len(letters)]


if __name__ == "__main__":
    solution = Solution()
    print(solution.nextGreatestLetter(["c", "f", "j"], "a"))
    print(solution.nextGreatestLetter(["c", "f", "j"], "c"))
    print(solution.nextGreatestLetter(["x", "x", "y", "y"], "z"))
