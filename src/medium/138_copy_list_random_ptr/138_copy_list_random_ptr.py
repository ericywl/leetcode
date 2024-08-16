"""
A linked list is given such that each node contains an additional random pointer which could point to 
any node in the list or null.

Return a deep copy of the list.
"""


class Node:
    # Definition for a Node.
    def __init__(self, val, next, random):
        self.val = val
        self.next = next
        self.random = random


class Solution:
    def copyRandomList(self, head: 'Node') -> 'Node':
        """O(n) time and O(1) space"""
        if not head:
            return None
        # Copy linked list and put side by side 
        # ie. 1 -> 1c -> 2 -> 2c -> 3 -> 3c
        curr = head
        while curr:
            copy = Node(curr.val, curr.next, curr.random)
            curr.next = copy
            curr = copy.next
        # Assign copy's random pointer to another copy
        curr = head
        while curr and curr.next:
            copy = curr.next
            if copy.random:
                copy.random = copy.random.next
            curr = copy.next
        # Split the two linked list 
        curr = head
        new_head = curr.next
        while curr:
            copy = curr.next
            curr.next = copy.next
            if copy.next:
                copy.next = copy.next.next
            curr = curr.next

        return new_head
        
            
