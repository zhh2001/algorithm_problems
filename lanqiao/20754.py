import sys

n = sys.stdin.read().rstrip('\n')
sys.stdout.write("YES" if sum(int(i) for i in n[:-3]) == sum(int(i) for i in n[-3:]) else "NO")
