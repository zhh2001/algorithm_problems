import sys

for line in sys.stdin:
    f = float(line.rstrip('\n'))
    print(f'{f:.3f}')
