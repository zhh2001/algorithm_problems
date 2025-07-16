import sys


def main():
    a = sys.stdin.readline().strip()
    a = int(a)
    print(a)

    b = sys.stdin.readline().strip()
    b = int(b)
    print(b)

    c = sys.stdin.readline().strip()
    c = float(c)
    print(f"{c:.1f}")

    d = sys.stdin.readline().strip()
    print(d)

    e = sys.stdin.readline().strip()
    print(e)


if __name__ == "__main__":
    main()
