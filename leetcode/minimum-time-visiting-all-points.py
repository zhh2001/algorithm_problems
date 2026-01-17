from typing import List


class Solution:
    def minTimeToVisitAllPoints(self, points: List[List[int]]) -> int:
        return sum(
            max(
                abs(points[i - 1][0] - points[i][0]),
                abs(points[i - 1][1] - points[i][1]),
            )
            for i in range(1, len(points))
        )


if __name__ == "__main__":
    solution = Solution()
    points = [[1, 1], [3, 4], [-1, 0]]
    print("输入: point =", points)
    print("输出:", solution.minTimeToVisitAllPoints(points))
