import math
import sys

D, V = map(int, sys.stdin.readline().split())
d3 = D ** 3 - 6 * V / math.pi
d = d3 ** (1 / 3)
sys.stdout.write(f"{d:.3f}")
