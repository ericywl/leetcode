package main

import "github.com/ericywl/leetcode/src/common"

func lowestCommonAncestor(node, p, q *common.TreeNode) *common.TreeNode {
	if node == nil || node == p || node == q {
		return node
	}

	left := lowestCommonAncestor(node.Left, p, q)
	right := lowestCommonAncestor(node.Right, p, q)

	// Found `p` and `q` on both left and right, i.e. current node is LCA
	if left != nil && right != nil {
		return node
	}

	// Found `p` or `q` on the left only
	if right == nil {
		return left
	}

	// Found `p` or `q` on the right only OR
	// The node doesn't exist (guaranteed to not happen by the problem statement)
	return right
}
