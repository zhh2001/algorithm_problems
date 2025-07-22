def main():
    s = input()
    chars = list(s)
    for index, char in enumerate(chars):
        if char == '5':
            chars[index] = '*'
    print(''.join(chars))


if __name__ == '__main__':
    main()
