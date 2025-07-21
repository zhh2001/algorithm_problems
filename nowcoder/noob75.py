class Point:
    def __init__(self, x=0.0, y=0.0):
        self.x = x
        self.y = y


class Line:
    def __init__(self, point_a=None, point_b=None):
        self.point_A = point_a if point_a is not None else Point()
        self.point_B = point_b if point_b is not None else Point()


def find_meeting_point(line_A, line_B):
    if all((line_A.point_A.x == line_A.point_B.x, line_B.point_A.x == line_B.point_B.x)):
        return Point(-1, -1)
    elif all((line_A.point_A.y == line_A.point_B.y, line_B.point_A.y == line_B.point_B.y)) :
        return Point(-1, -1)
    elif line_A.point_A.x == line_A.point_B.x:
        k_b = (line_B.point_A.y - line_B.point_B.y) / (line_B.point_A.x - line_B.point_B.x)
        y = k_b * (line_A.point_A.x - line_B.point_A.x) + line_A.point_A.y
        return Point(line_A.point_A.x, y)
    elif line_B.point_A.x == line_B.point_B.x:
        k_a = (line_A.point_A.y - line_B.point_B.y) / (line_A.point_A.x - line_B.point_B.x)
        y = k_a * (line_B.point_A.x - line_A.point_A.x) + line_B.point_A.y
        return Point(line_B.point_A.x, y)
    else:
        k_a = (line_A.point_A.y - line_A.point_B.y) / (line_A.point_A.x - line_A.point_B.x)
        k_b = (line_B.point_A.y - line_B.point_B.y) / (line_B.point_A.x - line_B.point_B.x)
        b1 = line_A.point_A.y - k_a * line_A.point_A.x
        b2 = line_B.point_A.y - k_b * line_B.point_A.x
        x = (b2 - b1) / (k_a - k_b)
        y = k_a * x + b1
        return Point(x, y)


def main():
    data = list(map(float, input().split()))
    data.extend(list(map(float, input().split())))
    A = Point(data[0], data[1])
    B = Point(data[2], data[3])
    C = Point(data[4], data[5])
    D = Point(data[6], data[7])
    AB = Line(A, B)
    CD = Line(C, D)
    ans = find_meeting_point(AB, CD)
    print(f"{ans.x} {ans.y}")


if __name__ == "__main__":
    main()
