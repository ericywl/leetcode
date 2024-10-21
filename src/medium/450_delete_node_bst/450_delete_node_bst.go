package main

import "github.com/ericywl/leetcode/src/common"

func deleteNode(root *common.TreeNode, key int) *common.TreeNode {
	if root == nil {
		return nil
	}

	if root.Val < key {
		// Key greater than current, we go down right side to delete node
		root.Right = deleteNode(root.Right, key)
	} else if root.Val > key {
		// Key less than current, we go down left side to delete node
		root.Left = deleteNode(root.Left, key)
	} else {
		// Key equals current, we need to consider some cases
		if root.Left == nil && root.Right == nil {
			// Both children are empty, safe to delete current node
			return nil
		}

		if root.Left == nil {
			// Only left child is empty, return right
			return root.Right
		}

		if root.Right == nil {
			// Only right child is empty, return left
			return root.Left
		}

		// Both children exists, we go towards right side (seems like the answer wants this)
		// and replace the current value with minimum value
		root.Val = findMin(root.Right)
		// Delete the duplicate minimum node
		root.Right = deleteNode(root.Right, root.Val)
	}

	return root
}

func findMin(node *common.TreeNode) int {
	curr := node
	minVal := curr.Val

	for curr != nil {
		if curr.Val < minVal {
			minVal = curr.Val
		}
		curr = curr.Left
	}

	return minVal
}
