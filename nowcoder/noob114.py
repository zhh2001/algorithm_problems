from typing import List


class ListNode:
    def __init__(self, x):
        self.val = x
        self.next = None


class Solution:
    def listnodeToVector(self, head: ListNode) -> List[int]:
        """

        :param head: ListNode类
        :return: int整型一维数组
        """
        vector = []
        node = head
        while node is not None:
            vector.append(node.val)
            node = node.next
        return vector


def main():
    head = ListNode(1)
    head.next = ListNode(1)
    head.next.next = ListNode(4)
    head.next.next.next = ListNode(5)
    head.next.next.next.next = ListNode(1)
    head.next.next.next.next.next = ListNode(4)
    print(Solution().listnodeToVector(head))


if __name__ == '__main__':
    main()
