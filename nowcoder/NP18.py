if __name__ == "__main__":
    string = input().strip()
    num_list = string.split(' ')
    num_list = [int(num) for num in num_list]
    print(num_list)
