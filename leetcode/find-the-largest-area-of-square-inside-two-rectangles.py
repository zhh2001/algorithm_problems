class Solution:
    def largestSquareArea(self, bottomLeft: List[List[int]], topRight: List[List[int]]) -> int:
        return int(math.pow(max(min(max(0, min(topRight[i][0], topRight[j][0]) - max(b1[0], b2[0])),
                                    max(0, min(topRight[i][1], topRight[j][1]) - max(b1[1], b2[1])))
                            for i, b1 in enumerate(bottomLeft)
                            for j, b2 in enumerate(bottomLeft[:i])), 2))
