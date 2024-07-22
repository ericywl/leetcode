package main

/*
Given a binary tree root, a node X in the tree is named good if in the path from root to X there are no nodes with a value greater than X.

Return the number of good nodes in the binary tree.
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func goodNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return goodNodesDFS(root, root.Val, 0)
}

func goodNodesDFS(node *TreeNode, maxSoFar int, count int) int {
	if node == nil {
		return count
	}

	if node.Val >= maxSoFar {
		count++
		maxSoFar = node.Val
	}

	count = goodNodesDFS(node.Left, maxSoFar, count)
	count = goodNodesDFS(node.Right, maxSoFar, count)
	return count
}
