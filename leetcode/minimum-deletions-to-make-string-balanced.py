import itertools


class Solution:
    def minimumDeletions(self, s: str) -> int:
        return min(
            itertools.accumulate(
                (-1 if c == "a" else 1 for c in s), initial=s.count("a")
            )
        )


if __name__ == "__main__":
    solution = Solution()
    print(solution.minimumDeletions("aababbab"))
    print(solution.minimumDeletions("bbaaaaabb"))
