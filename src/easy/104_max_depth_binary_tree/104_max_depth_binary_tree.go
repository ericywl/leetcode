package main

import "github.com/ericywl/leetcode/src/common"

/*
Given the root of a binary tree, return its maximum depth.

A binary tree's maximum depth is the number of nodes along the longest path from the root node down to the farthest leaf node.
*/

func maxDepthDFS(root *common.TreeNode) int {
	if root == nil {
		return 0
	}

	leftMaxDepth := maxDepthDFS(root.Left)
	rightMaxDepth := maxDepthDFS(root.Right)
	if leftMaxDepth > rightMaxDepth {
		return leftMaxDepth + 1
	}
	return rightMaxDepth + 1
}

func maxDepthBFS(root *common.TreeNode) int {
	if root == nil {
		return 0
	}

	depth := 0
	var queue []*common.TreeNode
	queue = append(queue, root)

	for len(queue) > 0 {
		depth++

		// Calculate the length only once before going into loop,
		// since we only want to pop the old nodes in the queue
		n := len(queue)
		for i := 0; i < n; i++ {
			// Pop the node
			node := queue[0]
			queue = queue[1:]

			// Append new left and right nodes to the queue
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
	}

	return depth
}
