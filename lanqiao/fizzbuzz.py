class Solution(object):

    def fizz_buzz(self, num):
        if num is None:
            raise TypeError('参数为 None')
        if num < 1:
            raise ValueError('参数小于 1')

        result = []
        for n in range(1, num + 1):
            if n % 3 == 0 and n % 5 == 0:
                result.append('FizzBuzz')
                continue

            if n % 3 == 0:
                result.append('Fizz')
                continue

            if n % 5 == 0:
                result.append('Buzz')
                continue

            result.append(str(n))

        return result


def main():
    s = Solution()
    print(s.fizz_buzz(9))
    print(s.fizz_buzz(23))
    print(s.fizz_buzz(0))
    print(s.fizz_buzz(-1))
    print(s.fizz_buzz(None))


if __name__ == '__main__':
    main()
