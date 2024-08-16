package main

import "github.com/ericywl/leetcode/src/common"

/*
Given the root of a binary tree and an integer targetSum, return the number of paths where the sum of the values along the path equals targetSum.

The path does not need to start or end at the root or a leaf, but it must go downwards (i.e., traveling only from parent nodes to child nodes).
*/

func pathSum(root *common.TreeNode, targetSum int) int {
	cache := map[int]int{0: 1}
	return pathSumDFS(root, targetSum, 0, cache)
}

func pathSumDFS(node *common.TreeNode, targetSum int, currPathSum int, cache map[int]int) int {
	if node == nil {
		return 0
	}

	// Traversing through the tree, at each node, we can get the `currPathSum` (from root to current node).
	// If within this path, there is a valid solution, then there must be an `oldPathSum` such that `currPathSum - oldPathSum = targetSum`.
	currPathSum += node.Val
	oldPathSum := currPathSum - targetSum

	// If `oldPathSum` exists in cache, it means that we have some consecutive nodes that sum to `targetSum`.
	count := cache[oldPathSum]
	// Save `currPathSum` into cache.
	cache[currPathSum]++

	// Recursively run the algorithm for left and right nodes.
	count += pathSumDFS(node.Left, targetSum, currPathSum, cache)
	count += pathSumDFS(node.Right, targetSum, currPathSum, cache)

	// Remove `currPathSum` from the cache since we are done with it.
	cache[currPathSum]--
	return count
}
