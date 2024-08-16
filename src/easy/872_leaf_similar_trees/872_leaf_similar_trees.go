package main

import (
	"slices"

	"github.com/ericywl/leetcode/src/common"
)

/*
Consider all the leaves of a binary tree, from left to right order, the values of those leaves form a leaf value sequence.

For example, in the given tree above, the leaf value sequence is (6, 7, 4, 9, 8).

Two binary trees are considered leaf-similar if their leaf value sequence is the same.

Return true if and only if the two given trees with head nodes root1 and root2 are leaf-similar.
*/

func leafSimilar(root1 *common.TreeNode, root2 *common.TreeNode) bool {
	var sequence1 []int
	leafDFS(root1, &sequence1)
	var sequence2 []int
	leafDFS(root2, &sequence2)

	return slices.Equal(sequence1, sequence2)
}

func leafDFS(root *common.TreeNode, sequence *[]int) {
	if root.Left == nil && root.Right == nil {
		*sequence = append(*sequence, root.Val)
		return
	}

	if root.Left != nil {
		leafDFS(root.Left, sequence)
	}
	if root.Right != nil {
		leafDFS(root.Right, sequence)
	}
}
