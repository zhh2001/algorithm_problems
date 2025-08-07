import sys


def main():
    word = sys.stdin.readline().rstrip()
    char_counter = {}
    for char in word:
        char_counter[char] = char_counter.get(char, 0) + 1

    count = 0
    char = ''
    for key, value in char_counter.items():
        if value > count:
            count = value
            char = key
        elif value == count:
            if key < char:
                count = value
                char = key
    print(char)
    print(count)


if __name__ == '__main__':
    main()
