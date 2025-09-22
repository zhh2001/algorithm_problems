import sys

t = int(sys.stdin.readline())
for _ in range(t):
    n = int(sys.stdin.readline())
    s = n
    s += 3 * (n + 1) * n // 2
    s += 3 * n * (n + 1) * (2 * n + 1) // 6
    s %= 998244353
    sys.stdout.write(str(s) + '\n')
