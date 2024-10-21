package main

import "github.com/ericywl/leetcode/src/common"

/*
Given a binary tree root, a node X in the tree is named good if in the path from root to X there are no nodes with a value greater than X.

Return the number of good nodes in the binary tree.
*/

func goodNodes(root *common.TreeNode) int {
	if root == nil {
		return 0
	}

	return goodNodesDFS(root, root.Val, 0)
}

func goodNodesDFS(node *common.TreeNode, maxSoFar int, count int) int {
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

type NodeWithCount struct {
	Node     *common.TreeNode
	MaxSoFar int
}

func goodNodesIter(root *common.TreeNode) int {
	stack := []*NodeWithCount{{Node: root, MaxSoFar: root.Val}}
	count := 0
	for len(stack) > 0 {
		// Pop node from stack
		nc := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		maxSoFar := nc.MaxSoFar
		// GTE because we use root.Val as the starting max, and we need to count it
		if nc.Node.Val >= nc.MaxSoFar {
			count++
			maxSoFar = nc.Node.Val
		}

		// Add right and left children to stack
		if nc.Node.Right != nil {
			stack = append(stack, &NodeWithCount{
				Node:     nc.Node.Right,
				MaxSoFar: maxSoFar,
			})
		}
		if nc.Node.Left != nil {
			stack = append(stack, &NodeWithCount{
				Node:     nc.Node.Left,
				MaxSoFar: maxSoFar,
			})
		}
	}

	return count
}
