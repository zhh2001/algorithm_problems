from linkedlist import LinkedList


class MyLinkedList(LinkedList):

    def remove_dupes(self):
        if len(self) == 0:
            return None
        s = {self.head.data}
        prev_node = self.head
        cur_node = self.head.next_node
        while cur_node is not None:
            if cur_node.data in s:
                prev_node.next_node = cur_node.next_node
                cur_node = prev_node.next_node
            else:
                s.add(cur_node.data)
                prev_node = cur_node
                cur_node = cur_node.next_node
        return None


def main():
    linked_list = MyLinkedList()
    linked_list.remove_dupes()
    print(linked_list.get_all_data())
    linked_list.append(1)
    linked_list.append(1)
    linked_list.append(1)
    linked_list.remove_dupes()
    print(linked_list.get_all_data())
    linked_list.append(2)
    linked_list.append(2)
    linked_list.append(3)
    linked_list.append(4)
    linked_list.append(4)
    linked_list.append(4)
    print(linked_list.get_all_data())
    linked_list.remove_dupes()
    print(linked_list.get_all_data())


if __name__ == '__main__':
    main()
