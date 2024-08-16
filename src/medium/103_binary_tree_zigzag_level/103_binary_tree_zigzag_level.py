"""
Given a binary tree, return the zigzag level order traversal of its nodes' values. 
(ie, from left to right, then right to left for the next level and alternate between).
"""

class TreeNode:
# Definition for a binary tree node.
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None

class Solution:
    def zigzagLevelOrder(self, root: TreeNode) -> List[List[int]]:
        """O(n) time and O(n) space"""
        if not root:
            return []
        result, next_level = [], [root]
        reverse = False
        while next_level:
            curr_level, temp = [], []
            for nd in next_level:
                curr_level.append(nd.val)
                if nd.left:
                    temp.append(nd.left)
                if nd.right:
                    temp.append(nd.right)
            # Reversing list is O(n)
            result.append(curr_level[::-1] if reverse else curr_level)
            next_level = temp
            reverse = not reverse
        return result