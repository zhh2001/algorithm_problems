class Solution:
    def isPalindrome(self, x: int) -> bool:
        x = str(x)
        n = len(x) // 2
        return x[n - 1::-1] == x[-n:]
