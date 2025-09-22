import sys

for line in sys.stdin.readlines()[1::2]:
    true = set()
    fake = set()
    for ticket in map(int, line.split()):
        if ticket not in true:
            true.add(ticket)
        else:
            fake.add(ticket)
    print(len(fake))
