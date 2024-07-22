package main

/*
Given the root of a binary tree and an integer targetSum, return the number of paths where the sum of the values along the path equals targetSum.

The path does not need to start or end at the root or a leaf, but it must go downwards (i.e., traveling only from parent nodes to child nodes).
*/

func pathSum(root *TreeNode, targetSum int) int {
	cache := map[int]int{0: 1}
	return pathSumDFS(root, targetSum, 0, cache)
}

func pathSumDFS(node *TreeNode, targetSum int, currPathSum int, cache map[int]int) int {
	if node == nil {
		return 0
	}

	currPathSum += node.Val
	count := cache[currPathSum-targetSum]
	cache[currPathSum]++

	count += pathSumDFS(node.Left, targetSum, currPathSum, cache)
	count += pathSumDFS(node.Right, targetSum, currPathSum, cache)

	cache[currPathSum]--
	return count
}
