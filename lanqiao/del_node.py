from linkedlist import LinkedList


class MyLinkedList(LinkedList):

    def delete_node(self, node):
        if len(self) == 0:
            return

        if self.head is node:
            self.head = self.head.next_node
            return

        if node is None:
            return

        next_node = node.next_node
        node.data = next_node.data
        node.next_node = next_node.next_node
        del next_node


def main():
    linked_list = MyLinkedList()
    linked_list.append(1)
    linked_list.append(2)
    linked_list.append(3)
    linked_list.append(4)
    linked_list.append(5)
    linked_list.append(6)
    print(linked_list.get_all_data())
    linked_list.delete_node(linked_list.head.next_node.next_node)
    linked_list.delete_node(None)
    print(linked_list.get_all_data())


if __name__ == '__main__':
    main()
