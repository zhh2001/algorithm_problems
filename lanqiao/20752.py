import sys

n = int(sys.stdin.read())
sys.stdout.write(["", "Q1", "Q1", "Q1", "Q2", "Q2", "Q2", "Q3", "Q3", "Q3", "Q4", "Q4", "Q4"][n] if 1 <= n <= 12 else "??")
