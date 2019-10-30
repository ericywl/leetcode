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
        first = second = head
        for _ in range(n):
            second = second.next
        if not second:
            return head.next
        while second.next:
            first = first.next
            second = second.next
        first.next = first.next.next
        return head
