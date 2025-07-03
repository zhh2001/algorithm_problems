if __name__ == "__main__":
    string = input().strip()
    name_list = string.split(' ')

    name_list.pop()
    name_list.pop()
    name_list.pop()

    print(name_list)
