class ListNode:
    def __init__(self, x):
        self.val = x
        self.next = None


class Solution:
    def ReverseList(self, head: ListNode) -> ListNode:
        if head is None or head.next is None:
            return head
        newHead = self.ReverseList(head.next)
        head.next.next = head
        head.next = None
        return newHead


def test1():
    solution = Solution()
    head = node1 = ListNode(1)
    node2 = ListNode(2)
    node3 = ListNode(3)
    node1.next = node2
    node2.next = node3
    new_head = solution.ReverseList(head)
    node = new_head
    values = []
    while node is not None:
        values.append(node.val)
        node = node.next
    print(values)


def test2():
    solution = Solution()
    head = None
    new_head = solution.ReverseList(head)
    node = new_head
    values = []
    while node is not None:
        values.append(node.val)
        node = node.next
    print(values)


if __name__ == '__main__':
    test1()
    test2()
