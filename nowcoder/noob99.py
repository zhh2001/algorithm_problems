import bisect
from typing import Final

data: Final[list] = []


def insertValue(x):
    idx = bisect.bisect_left(data, x)
    if idx == len(data) or data[idx] != x:
        data.insert(idx, x)


def eraseValue(x):
    idx = bisect.bisect_left(data, x)
    if idx < len(data) and data[idx] == x:
        data.pop(idx)


def xInSet(x):
    idx = bisect.bisect_left(data, x)
    return idx < len(data) and data[idx] == x


def sizeOfSet():
    return len(data)


def getPre(x):
    idx = bisect.bisect_left(data, x)
    return data[idx - 1] if idx > 0 else -1


def getBack(x):
    idx = bisect.bisect_right(data, x)
    return data[idx] if idx < len(data) else -1


def main():
    q = int(input())
    for _ in range(q):
        line = map(int, input().split())
        cnt, op, x = 0, 0, 0
        for i in line:
            if cnt == 0:
                op = i
            else:
                x = i
            cnt += 1

        if op == 1:
            insertValue(x)
        elif op == 2:
            eraseValue(x)
        elif op == 3:
            print('YES' if xInSet(x) else 'NO')
        elif op == 4:
            print(sizeOfSet())
        elif op == 5:
            print(getPre(x))
        elif op == 6:
            print(getBack(x))


if __name__ == '__main__':
    main()
