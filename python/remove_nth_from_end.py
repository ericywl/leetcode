"""
Given a linked list, remove the n-th node from the end of list and return its head.
"""


# Definition for singly-linked list.
class ListNode:
    def __init__(self, x):
        self.val = x
        self.next = None


class Solution:
    def removeNthFromEnd(self, head: ListNode, n: int) -> ListNode:
        curr = head
        first = None
        while True:
            second = curr
            for _ in range(n):
                second = second.next
            if second:
                first = curr
                curr = curr.next
            else:
                break
        if not first:
            if head.next:
                return head.next
            return None
        first.next = first.next.next
        return head
