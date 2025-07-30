class Permutations(object):

    def is_permutation(self, str1, str2):
        if str1 is None or str2 is None:
            return False

        if str1 == '' or str2 == '':
            return False

        chars = {}
        for char in str1:
            chars[char] = chars.get(char, 0) + 1

        for char in str2:
            if char not in chars:
                return False
            chars[char] -= 1
            if chars[char] == 0:
                del chars[char]

        return not bool(chars)


if __name__ == '__main__':
    permutations = Permutations()
    print(permutations.is_permutation('abc', None))
    print(permutations.is_permutation('', 'ac'))
    print(permutations.is_permutation('abc', 'ac'))
    print(permutations.is_permutation('abc', 'bac'))
    print(permutations.is_permutation('abc', 'bdc'))
