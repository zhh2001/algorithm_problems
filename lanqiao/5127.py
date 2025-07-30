def is_three_and_one(arr):
    counter = {}
    for n in arr:
        counter[n] = counter.get(n, 0) + 1
    flag = False
    for value in counter.values():
        if value == 3:
            flag = True
    return flag


def main():
    t = int(input())
    for _ in range(t):
        arr = input()
        print('Yes' if is_three_and_one(arr) else 'No')


if __name__ == '__main__':
    main()
