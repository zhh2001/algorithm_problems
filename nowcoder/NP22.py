if __name__ == "__main__":
    string = input().strip()
    name_list = string.split(' ')
    del name_list[0]
    print(name_list)
