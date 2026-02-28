class Solution:
    def minPartitions(self, n: str) -> int:
        return max(int(c) for c in n)


if __name__ == "__main__":
    solution = Solution()
    print(solution.minPartitions("32"))
    print(solution.minPartitions("82734"))
    print(solution.minPartitions("27346209830709182346"))
