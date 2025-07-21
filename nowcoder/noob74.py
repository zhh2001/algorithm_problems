import math


class Point:
    def __init__(self, A, B):
        self.x = A
        self.y = B


class Line:
    def __init__(self, A, B):
        self.point_A = A
        self.point_B = B


class Circle:
    def __init__(self, A, B):
        self.O = A
        self.r = B


def point_to_line_distance(point, line):
    # 计算点到直线的距离
    x1, y1 = line.point_A.x, line.point_A.y
    x2, y2 = line.point_B.x, line.point_B.y
    x0, y0 = point.x, point.y

    # 直线方程 Ax + By + C = 0
    A = y2 - y1
    B = x1 - x2
    C = x2 * y1 - x1 * y2

    distance = abs(A * x0 + B * y0 + C) / math.sqrt(A ** 2 + B ** 2)
    return distance


def getDistance(circle, l):
    x1, y1 = l.point_A.x, l.point_A.y
    x2, y2 = l.point_B.x, l.point_B.y
    x0, y0 = circle.O.x, circle.O.y

    A = y2 - y1
    B = x1 - x2
    C = x2 * y1 - x1 * y2

    distance = abs(A * x0 + B * y0 + C) / math.sqrt(A ** 2 + B ** 2)
    return 2 * math.sqrt(circle.r ** 2 - distance ** 2) if distance < circle.r else 0


def main():
    ox, oy, r = map(float, input().split())
    x1, y1, x2, y2 = map(float, input().split())

    center = Point(ox, oy)
    circle = Circle(center, int(r))

    p1 = Point(x1, y1)
    p2 = Point(x2, y2)
    l = Line(p1, p2)

    result = getDistance(circle, l)
    print("{:.6f}".format(result))


if __name__ == "__main__":
    main()
