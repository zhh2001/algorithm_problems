class CompressString(object):

    def compress(self, string):

        if string is None:
            return None

        if string == '':
            return ''

        res = ''
        cnt = 1
        cur = string[0]
        for char in string[1:]:
            if char == cur:
                cnt += 1
            else:
                res += cur
                if cnt > 1:
                    res += str(cnt)
                    cnt = 1
                cur = char

        res += cur
        if cnt > 1:
            res += str(cnt)

        return res if len(res) < len(string) else string


if __name__ == '__main__':
    compress_string = CompressString()
    print(compress_string.compress(None))
    print(compress_string.compress(''))
    print(compress_string.compress('AAABCCDDDD'))
    print(compress_string.compress('AA'))
    print(compress_string.compress('AB'))
