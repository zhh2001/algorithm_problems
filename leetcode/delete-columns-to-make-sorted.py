from typing import List


class Solution:
    def minDeletionSize(self, strs: List[str]) -> int:
        return sum(s != tuple(sorted(s)) for s in zip(*strs))


if __name__ == '__main__':
    solution = Solution()
    print(solution.minDeletionSize(["cba", "daf", "ghi"]))
    print(solution.minDeletionSize(["a", "b"]))
    print(solution.minDeletionSize(["zyx", "wvu", "tsr"]))
