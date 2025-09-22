import sys

sys.stdout.write(str(sum(line[0] != line[1] and line[1] == line[2] for line in sys.stdin.readlines()[1:])))
