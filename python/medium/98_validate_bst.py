"""
Given a binary tree, determine if it is a valid binary search tree (BST).

Assume a BST is defined as follows:
    1. The left subtree of a node contains only nodes with keys less than the node's key.
    2. The right subtree of a node contains only nodes with keys greater than the node's key.
    3. Both the left and right subtrees must also be binary search trees.
"""

import sys


class TreeNode:
    # Definition for a binary tree node.
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None


class Solution:
    def isValidBST(self, root: TreeNode) -> bool:

        def isValidBSTRecurse(root: TreeNode, minVal: int, maxVal: int) -> bool:
            if not root:
                return True
            if root.val >= maxVal or root.val <= minVal:
                return False
            return (
                isValidBSTRecurse(root.left, minVal, root.val) and
                isValidBSTRecurse(root.right, root.val, maxVal)
            )
            
        return isValidBSTRecurse(root, -(sys.maxsize - 1), sys.maxsize)
