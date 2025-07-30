class Rotation(object):

    def is_substring(self, s1, s2):
        if s1 is None or s2 is None:
            return False

        return s1.find(s2) != -1

    def is_rotation(self, s1, s2):
        if s1 is None or s2 is None:
            return False

        if len(s1) != len(s2):
            return False

        return self.is_substring(s1 + s1, s2)


if __name__ == '__main__':
    rotation = Rotation()
    print(rotation.is_rotation('abcd', 'dcab'))
    print(rotation.is_rotation('abcd', 'bcda'))
    print(rotation.is_rotation('abcd', 'abcd'))
    print(rotation.is_rotation('abcd', None))
    print(rotation.is_rotation('', 'a'))
    print(rotation.is_rotation('cd', 'abcd'))
