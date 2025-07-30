B = 1
KB = 1024 * B
MB = 1024 * KB
GB = 1024 * MB


def main():
    map3_size = 4 * MB
    disk_size = 100 * GB
    num = int(disk_size / map3_size)
    print(num)


if __name__ == '__main__':
    main()
