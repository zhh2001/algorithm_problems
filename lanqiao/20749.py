import sys

t = int(sys.stdin.readline())
for _ in range(t):
    n = int(sys.stdin.readline())
    res = 0
    for x in range(1, int(n ** 0.5) + 1):
        if n % x:
            continue
        y = n // x
        if x % 10 == 0 or y % 10 == 0:
            continue
        print(x + y)
        break
