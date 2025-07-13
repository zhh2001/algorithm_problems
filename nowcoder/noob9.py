import sys

for line in sys.stdin:
    n = int(line.strip())
    print(abs(n) % 10)
