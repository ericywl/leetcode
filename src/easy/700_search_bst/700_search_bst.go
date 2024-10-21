package main

import "github.com/ericywl/leetcode/src/common"

func searchBST(root *common.TreeNode, val int) *common.TreeNode {
	curr := root
	for curr != nil {
		if val == curr.Val {
			return curr
		} else if val < curr.Val {
			curr = curr.Left
		} else {
			curr = curr.Right
		}
	}

	return nil
}
