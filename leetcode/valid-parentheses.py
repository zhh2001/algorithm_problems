wrapper = {
    ")": "(",
    "}": "{",
    "]": "["
}


class Solution:
    def isValid(self, s: str) -> bool:
        stack = []
        for ch in s:
            if not stack or ch not in wrapper:
                stack.append(ch)
                continue
            if wrapper[ch] == stack[-1]:
                stack.pop()
            else:
                stack.append(ch)
        return not stack
