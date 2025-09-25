import sys

n = int(sys.stdin.read())
r = n // 2
c = n - r
sys.stdout.write(str((r + 1) * (c + 1)))
