package main

import "github.com/ericywl/leetcode/src/common"

func maxLevelSum(root *common.TreeNode) int {
	if root == nil {
		return 0
	}

	depth := 1
	maxSum := root.Val
	maxDepth := depth
	queue := []*common.TreeNode{root}

	for len(queue) > 0 {
		qLen := len(queue)
		sum := 0
		for i := 0; i < qLen; i++ {
			// Pop node from front of queue
			node := queue[0]
			queue = queue[1:]
			// Compute sum of this level
			sum += node.Val

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		if sum > maxSum {
			maxDepth = depth
			maxSum = sum
		}
		depth++
	}

	return maxDepth
}
