def main():
    total = 0
    for i in range(1, 2021):
        total += str(i).count('2')
    print(total)


if __name__ == '__main__':
    main()
