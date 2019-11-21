"""
Given a binary tree, return the inorder traversal of its nodes' values.
"""


class TreeNode:
    # Definition for a binary tree node.
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None


class Solution:
    """Recursive solution"""

    def inorderTraversal(self, root: TreeNode) -> List[int]:
        if not root:
            return

        def inorderRecurse(root: TreeNode, result: List[int]):
            if not root:
                return
            inorderRecurse(root.left, result)
            result.append(root.val)
            inorderRecurse(root.right, result)

        result = []
        inorderRecurse(root, result)
        return result


class Solution2:
    """Iterative solution"""

    def inorderTraversal(self, root: TreeNode) -> List[int]:
        if not root:
            return
        stack, result = [], []
        nd = root
        while True:
            if nd:
                stack.append(nd)
                nd = nd.left
            elif stack:
                nd = stack.pop()
                result.append(nd.val)
                nd = nd.right
            else:
                break
        return result
