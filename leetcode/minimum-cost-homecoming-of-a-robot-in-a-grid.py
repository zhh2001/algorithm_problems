from typing import List


class Solution:
    def minCost(
        self,
        startPos: List[int],
        homePos: List[int],
        rowCosts: List[int],
        colCosts: List[int],
    ) -> int:
        return (
            sum(
                rowCosts[
                    min(startPos[0], homePos[0]) : max(startPos[0], homePos[0]) + 1
                ]
            )
            - rowCosts[startPos[0]]
        ) + (
            sum(
                colCosts[
                    min(startPos[1], homePos[1]) : max(startPos[1], homePos[1]) + 1
                ]
            )
            - colCosts[startPos[1]]
        )


if __name__ == "__main__":
    solution = Solution()
    print(solution.minCost([1, 0], [2, 3], [5, 4, 3], [8, 2, 6, 7]))
    print(solution.minCost([0, 0], [0, 0], [5], [26]))
