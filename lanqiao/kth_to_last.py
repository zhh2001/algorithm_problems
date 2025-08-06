from linkedlist import LinkedList


class MyLinkedList(LinkedList):

    def kth_to_last_elem(self, k):
        length = len(self)

        if length == 0:
            return None

        if k >= length:
            return None

        prev_node = cur_node = self.head
        for _ in range(k):
            cur_node = cur_node.next_node

        while cur_node.next_node is not None:
            prev_node = prev_node.next_node
            cur_node = cur_node.next_node

        return prev_node.data


def main():
    linked_list = MyLinkedList()
    linked_list.append(1)
    linked_list.append(2)
    linked_list.append(3)
    linked_list.append(4)
    linked_list.append(5)
    linked_list.append(6)
    print(linked_list.kth_to_last_elem(3))
    print(linked_list.kth_to_last_elem(6))
    print(linked_list.kth_to_last_elem(7))


if __name__ == '__main__':
    main()
