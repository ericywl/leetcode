"""
Given preorder and inorder traversal of a tree, construct the binary tree.

Note:
You may assume that duplicates do not exist in the tree.
"""


class TreeNode:
    # Definition for a binary tree node.
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None


class Solution:
    def buildTree(self, preorder: List[int], inorder: List[int]) -> TreeNode:
        """Recursive solution"""
        if not preorder or not inorder:
            return None

        def buildTreeRecurse(preorder: List[int], p_start: int, inorder: List[int],
                             i_start: int, i_end: int) -> TreeNode:
            if p_start >= len(preorder) - 1 or i_start > i_end:
                return None
            val = preorder[p_start]
            root = TreeNode(val)
            mid_idx = inorder.index(val)
            if mid_idx < 0:
                raise Exception(f"Can't find value in inorder array: {val}")
            root.left = buildTreeRecurse(
                preorder, p_start + 1, inorder, i_start, mid_idx - 1)
            root.right = buildTreeRecurse(
                preorder, p_start - i_start + mid_idx + 1, inorder, mid_idx + 1, i_end)
            return root

        return buildTreeRecurse(preorder, 0, inorder, 0, len(inorder) - 1)


class Solution2:
    def buildTree(self, preorder: List[int], inorder: List[int]) -> TreeNode:
        """Iterative solution"""
        if not preorder or not inorder:
            return None
        root = TreeNode(preorder[0])
        stack = [root]
        p_idx, i_idx = 1, 0
        while p_idx < len(preorder):
            curr = TreeNode(preorder[p_idx])
            parent = None
            while stack and stack[-1].val == inorder[i_idx]:
                parent = stack.pop()
                i_idx += 1
            if parent:
                parent.right = curr
            else:
                stack[-1].left = curr
            p_idx += 1
            stack.append(curr)
        return root
