"""
You are given a perfect binary tree where all leaves are on the same level, and every parent has two children.
Populate each next pointer to point to its next right node. If there is no next right node, 
the next pointer should be set to NULL.
Initially, all next pointers are set to NULL.
"""

class Node:
    # Definition for a Node.
    def __init__(self, val, left, right, next):
        self.val = val
        self.left = left
        self.right = right
        self.next = next

class Solution:
    def connect(self, root: 'Node') -> 'Node':
        start = root
        while start:
            curr = start
            while curr:
                if curr.left:
                    curr.left.next = curr.right
                if curr.right and curr.next:
                    curr.right.next = curr.next.left
                curr = curr.next
            start = start.left
        return root