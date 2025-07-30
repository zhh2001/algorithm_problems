class UniqueChars(object):

    def has_unique_chars(self, string):
        if string is None:
            return False

        charset = set()
        for char in string:
            if char in charset:
                return False
            charset.add(char)
        return True


if __name__ == '__main__':
    uc = UniqueChars()
    print(uc.has_unique_chars('abc'))
    print(uc.has_unique_chars('abbc'))
    print(uc.has_unique_chars(None))
