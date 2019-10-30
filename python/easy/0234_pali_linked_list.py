"""
Given a singly linked list, determine if it is a palindrome.
"""


class Solution:
    def reverse_part_list(self, head: ListNode, until: int) -> ListNode:
        """O(n) time and O(1) space"""
        prev, nxt = None, None
        curr = head
        while curr != until:
            nxt = curr.next
            curr.next = prev
            prev = curr
            curr = nxt
        return prev

    def check_same(self, n1: ListNode, n2: ListNode) -> bool:
        """O(n) time and O(1) space"""
        h1, h2 = n1, n2
        while h1 and h2:
            if h1.val != h2.val:
                return False
            h1, h2 = h1.next, h2.next
        # This checks if the two are the same length
        return not h1 and not h2

    def isPalindrome(self, head: ListNode) -> bool:
        """O(n) time and O(1) space"""
        h1 = h2 = head
        # Assign h1 to 1-after halfway of list
        # h2 moving 2x faster so it will reach the end
        while h2 and h2.next:
            h2 = h2.next.next
            h1 = h1.next
        # When h2 reach the end, h1 should be in middle
        mid = h1
        # If length is odd, we advance h1 to cross the middle
        if h2:
            h1 = h1.next
        # Reverse the part of list from 0 to halfway
        h2 = self.reverse_part_list(head, mid)
        return self.check_same(h1, h2)
