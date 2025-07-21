class Point:
    def __init__(self, x, y):
        self.x = x
        self.y = y


class Triangle:
    def __init__(self, a, b, c):
        self.a = a
        self.b = b
        self.c = c


def get_area(T):
    vector1 = T.a.x - T.b.x, T.a.y - T.b.y
    vector2 = T.a.x - T.c.x, T.a.y - T.c.y
    area = abs(vector1[0] * vector2[1] - vector1[1] * vector2[0]) / 2
    return area


def main():
    x, y = map(int, input().split())
    a = Point(x, y)

    x, y = map(int, input().split())
    b = Point(x, y)

    x, y = map(int, input().split())
    c = Point(x, y)

    T = Triangle(a, b, c)
    print("{:.2f}".format(get_area(T)))


if __name__ == "__main__":
    main()
