import math


class Point:
    def __init__(self, x, y):
        self.x = x
        self.y = y


class Line:
    def __init__(self, point_a, point_b):
        self.point_a = point_a
        self.point_b = point_b


def get_distance(P, L):
    delta_x = L.point_a.x - L.point_b.x
    delta_y = L.point_a.y - L.point_b.y

    if delta_x == 0:
        return abs(L.point_a.x - P.x)

    k = delta_y / delta_x
    c = L.point_a.y - k * L.point_a.x

    distance = abs(k * P.x - P.y + c) / math.sqrt(pow(k, 2) + pow(-1, 2))
    return round(distance, 2)


def main():
    a, b = map(int, input().split())
    sx, sy, tx, ty = map(int, input().split())

    point_a = Point(sx, sy)
    point_b = Point(tx, ty)
    point_c = Point(a, b)
    line = Line(point_a, point_b)

    print("{:.2f}".format(get_distance(point_c, line)))


if __name__ == "__main__":
    main()
