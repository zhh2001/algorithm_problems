import sys

cnt = {}
res = 0
x = int(sys.stdin.readline().split()[1])
for a in map(int, sys.stdin.readline().split()):
    cnt[a] = cnt.get(a, 0) + 1
    if cnt[a] >= x:
        cnt[a] = x
    total = cnt[a] * a
    if total > res:
        res = total
sys.stdout.write(str(res))
