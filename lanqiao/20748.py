import sys

cost = 0
n, x = map(int, sys.stdin.readline().split())
ns = sorted(map(int, sys.stdin.readline().split()), reverse=True)
for i in range(n):
    value = ns.pop() - x * i
    if value <= 0:
        continue
    cost += value
print(cost)
