import math


class Point:
    def __init__(self, a=0, b=0):
        self.x = a
        self.y = b


# 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
#
# 计算A点与B点之间的距离
# @param point_A Point类 A点
# @param point_B Point类 B点
# @return double浮点型
#
class Solution:
    def calculateDistance(self, point_A: Point, point_B: Point) -> float:
        delta_x = point_A.x - point_B.x
        delta_y = point_A.y - point_B.y
        distance = math.sqrt(math.pow(delta_x, 2) + math.pow(delta_y, 2))
        return distance


if __name__ == '__main__':
    Point_A = Point(1, 1)
    Point_B = Point(1, 8)
    distance = Solution().calculateDistance(Point_A, Point_B)
    print(distance)
