"""
Given a binary tree, return the level order traversal of its nodes' values. (ie, from left to right, level by level).
"""


class TreeNode:
    # Definition for a binary tree node.
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None


class Solution:
    """O(n) time and O(n) space"""
    def levelOrder(self, root: TreeNode) -> List[List[int]]:
        if not root:
            return []
        result, next_level = [], [root]
        while next_level:
            curr_level, temp = [], []
            for nd in next_level:
                curr_level.append(nd.val)
                if nd.left:
                    temp.append(nd.left)
                if nd.right:
                    temp.append(nd.right)
            result.append(curr_level)
            next_level = temp
        return result
