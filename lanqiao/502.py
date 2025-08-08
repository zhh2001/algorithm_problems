import sys


def main():
    n = int(sys.stdin.readline())
    count1 = count2 = 0
    for _ in range(n):
        score = int(sys.stdin.readline())
        if score >= 60:
            count1 += 1
            if score >= 85:
                count2 += 1
    print(f'{round(count1 / n)}%')
    print(f'{round(count2 / n)}%')


if __name__ == '__main__':
    main()
