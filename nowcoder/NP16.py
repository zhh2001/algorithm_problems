if __name__ == "__main__":
    offer_list = ["Allen", "Tom"]
    for name in offer_list:
        print(
            f"{name}, you have passed our interview and will soon become a member of our company."
        )
    for index, name in enumerate(offer_list):
        if name == "Tom":
            offer_list[index] = "Andy"
            break
    for name in offer_list:
        print(
            f"{name}, welcome to join us!"
        )
