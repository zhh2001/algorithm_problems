import math
import sys

while True:
    D, V = map(int, sys.stdin.readline().split())
    if D == 0 and V == 0:
        break
    d3 = D ** 3 - 6 * V / math.pi
    d = d3 ** (1 / 3)
    sys.stdout.write(f"{d:.3f}\n")
