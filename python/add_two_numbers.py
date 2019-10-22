"""
You are given two non-empty linked lists representing two non-negative integers. The digits are stored in reverse order and each of their nodes contain a single digit. Add the two numbers and return it as a linked list.

You may assume the two numbers do not contain any leading zero, except the number 0 itself.
"""

# Definition for singly-linked list.
class ListNode:
    def __init__(self, x):
        self.val = x
        self.next = None


class Solution:
    def addTwoNumbers(self, l1: ListNode, l2: ListNode) -> ListNode:
        summ = l1.val + l2.val
        res = ListNode(summ % 10)
        carry = summ // 10
        l1 = l1.next
        l2 = l2.next
        curr = res

        while l1 and l2:
            summ = l1.val + l2.val + carry
            curr.next = ListNode(summ % 10)
            carry = summ // 10
            l1 = l1.next
            l2 = l2.next
            curr = curr.next

        if l1:
            rem = l1
        else:
            rem = l2

        while rem:
            summ = rem.val + carry
            curr.next = ListNode(summ % 10)
            carry = summ // 10
            rem = rem.next
            curr = curr.next

        if carry:
            curr.next = ListNode(carry)

        return res
