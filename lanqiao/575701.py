import sys

N = int(sys.stdin.readline())
for _ in range(N):
    x, y = map(int, sys.stdin.readline().split())
    if x - y > 2 or x - y < 0 or x - y == 1:
        print("No Number")
        continue

    res = x + y
    if x % 2 == 1:
        res -= 1
    print(res)
