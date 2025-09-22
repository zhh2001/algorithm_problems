import sys

n, m = map(int, sys.stdin.readline().split())
a = list(map(int, sys.stdin.readline().split()))

r = [0] * (m + 1)
l = [0] * (m + 1)

o = 0
for i in a:
    if i > 0 and abs(i) <= m:
        r[i] += 1
    elif i < 0 and abs(i) <= m:
        l[abs(i)] += 1
    elif i == 0:
        o = 1

for i in range(1, m + 1):
    r[i] += r[i - 1]
    l[i] += l[i - 1]

ans = 0

for j in range(1, m // 2 + 1):
    ans = max(ans, r[j] + l[m - 2 * j], l[j] + r[m - 2 * j])

ans = max(r[m], l[m], ans)

print(ans + o)
