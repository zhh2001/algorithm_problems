class Solution(object):

    def two_sum(self, nums, val):
        if nums is None or val is None:
            raise TypeError('参数不能为 None')

        if isinstance(nums, list) and nums == []:
            raise ValueError('第一个参数不能为空数组')

        target_index_mapping = {}
        for index, num in enumerate(nums):
            if num in target_index_mapping:
                return [target_index_mapping[num], index]
            target_index_mapping[val - num] = index
        return None


def main():
    s = Solution()
    nums = [1, 2, 3, -2, 5, 7]
    print(s.two_sum(nums, 7))
    print(s.two_sum(nums, 2))
    print(s.two_sum(nums, None))
    print(s.two_sum(None, 7))
    print(s.two_sum([], 7))


if __name__ == '__main__':
    main()
