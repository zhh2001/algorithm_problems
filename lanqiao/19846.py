import sys

n = int(sys.stdin.readline())
s = set(map(int, sys.stdin.readline().split()))
for i in range(1, n + 2):
    if i not in s:
        sys.stdout.write(str(i))
        break
