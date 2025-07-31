class ReverseString(object):

    def reverse(self, chars):
        if chars is None or len(chars) == 0:
            return chars

        i = 0
        j = len(chars) - 1
        while i < j:
            chars[i], chars[j] = chars[j], chars[i]
            i += 1
            j -= 1

        return chars


def main():
    rs = ReverseString()
    print(rs.reverse(['h', 'e', 'l', 'l', 'o']))
    print(rs.reverse(None))
    print(rs.reverse([]))


if __name__ == '__main__':
    main()
