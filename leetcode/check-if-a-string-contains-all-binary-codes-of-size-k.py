class Solution:
    def hasAllCodes(self, s: str, k: int) -> bool:
        return len({s[i : i + k] for i in range(len(s) - k + 1)}) == 2**k


if __name__ == "__main__":
    solution = Solution()
    print(solution.hasAllCodes("00110110", 2))
    print(solution.hasAllCodes("0110", 1))
    print(solution.hasAllCodes("0110", 2))
