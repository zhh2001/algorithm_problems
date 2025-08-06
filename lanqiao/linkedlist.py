class Node(object):

    def __init__(self, data, next_node=None):
        self.data = data
        self.next_node = next_node

    def __str__(self) -> str:
        return str(self.data)


class LinkedList(object):

    def __init__(self, head=None):
        self.head = head

    def __len__(self) -> int:
        length = 0
        node = self.head
        while node is not None:
            length += 1
            node = node.next_node
        return length

    def insert_to_front(self, data):
        if data is None:
            return None
        head = Node(data, self.head)
        self.head = head
        return self.head

    def append(self, data):
        if data is None:
            return None
        last_node = Node(data)
        if self.head is None:
            self.head = last_node
            return last_node
        node = self.head
        while node.next_node is not None:
            node = node.next_node
        node.next_node = last_node
        return last_node

    def find(self, data):
        node = self.head
        while node is not None:
            if node.data == data:
                return node
            node = node.next_node
        return None

    def delete(self, data):
        if self.head is None:
            return None
        if self.head.data == data:
            return self.head.next_node
        prev_node = self.head
        next_node = self.head.next_node
        while next_node is not None:
            if next_node.data == data:
                prev_node.next_node = next_node.next_node
                break
            prev_node = next_node
            next_node = next_node.next_node
        return None

    def print_list(self):
        node = self.head
        while node is not None:
            print(node)
            node = node.next_node

    def get_all_data(self):
        all_data = []
        node = self.head
        while node is not None:
            all_data.append(node.data)
            node = node.next_node
        return all_data


def main():
    head = Node(1)
    linked_list = LinkedList(head)
    linked_list.append(2)
    linked_list.append(3)
    linked_list.insert_to_front(0)
    linked_list.delete(2)
    print('len:', len(linked_list))
    print('all data:', linked_list.get_all_data())
    linked_list.print_list()
    print(linked_list.find(1))
    print(linked_list.find(2))


if __name__ == '__main__':
    main()
