class Solution(object):

    def find_diff(self, str1, str2):
        if str1 is None or str2 is None:
            raise TypeError('传入的字符串参数为 None')

        char_counter = {}
        for char in str1:
            char_counter[char] = char_counter.get(char, 0) + 1

        for char in str2:
            if char not in char_counter:
                return char

            char_counter[char] -= 1
            if char_counter[char] == 0:
                del char_counter[char]

        return list(char_counter.keys())[0]


def main():
    s = Solution()
    print(s.find_diff('aad', 'ad'))
    print(s.find_diff('aaabccdd', 'abdcacade'))
    print(s.find_diff('aaabccdd', None))


if __name__ == '__main__':
    main()
