import sys

t = int(sys.stdin.readline())
for _ in range(t):
    n, x = map(int, sys.stdin.readline().split(' '))
    sys.stdout.write("YES" if x >= n / 2 else "NO")
    sys.stdout.write('\n')
