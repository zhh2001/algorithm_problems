if __name__ == "__main__":
    string = input().strip()
    name_list = string.split(' ')

    name = input().strip()
    name_list.remove(name)

    print(name_list)
