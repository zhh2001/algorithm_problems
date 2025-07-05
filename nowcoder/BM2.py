class ListNode:
    def __init__(self, x):
        self.val = x
        self.next = None


class Solution:
    def reverseBetween(self, head: ListNode, m: int, n: int) -> ListNode:
        res = ListNode(-1)
        res.next = head
        pre = res
        cur = head
        for _ in range(1, m):
            pre = cur
            cur = cur.next
        for _ in range(m, n):
            temp = cur.next
            cur.next = temp.next
            temp.next = pre.next
            pre.next = temp
        return res.next
