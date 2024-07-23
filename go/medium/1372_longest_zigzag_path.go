package main

func longestZigZagDFS(root *TreeNode) int {
	return zigZagDFS(root).maxZigZagWhole
}

type Result struct {
	// maxZigZagLeft is the maximum zigzag length in direction of node.Left
	maxZigZagLeft int
	// maxZigZagRight is the maximum zigzag length in direction of node.Right
	maxZigZagRight int
	// maxZigZagWhole is the maximum zigzag length in the whole sub-tree
	maxZigZagWhole int
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func zigZagDFS(node *TreeNode) Result {
	if node == nil {
		return Result{maxZigZagLeft: -1, maxZigZagRight: -1, maxZigZagWhole: -1}
	}

	// Run DFS on both left and right nodes
	leftRes := zigZagDFS(node.Left)
	rightRes := zigZagDFS(node.Right)

	// There are two scenarios:
	// 	1. `continueZigZagLen` - the zigzag continues from the top, we add one to the length (happens when L -> R or R -> L)
	//  2. `newZigZagLen` - the zigzag starts from current node, we start anew (happens when L -> L or R -> R)
	continueZigZagLen := max(leftRes.maxZigZagRight, rightRes.maxZigZagLeft) + 1
	newZigZagLen := max(leftRes.maxZigZagLeft, rightRes.maxZigZagRight)
	// Find the sub-tree maximum between `continueZigZagLen` and `newZigZagLen`
	newMaxWhole := max(continueZigZagLen, newZigZagLen)

	// The `maxZigZagLeft` will only increase if there are any results to its right.
	// Similarly, the `maxZigZagRight` will only increase if there are any results to its left.
	return Result{maxZigZagLeft: leftRes.maxZigZagRight + 1, maxZigZagRight: rightRes.maxZigZagLeft + 1, maxZigZagWhole: newMaxWhole}
}

type Element struct {
	node      *TreeNode
	fromLeft  bool
	zigZagLen int
}

func longestZigZagBFS(root *TreeNode) int {
	if root == nil {
		return 0
	}

	var queue []Element
	maxZigZagLen := 0

	if root.Left != nil {
		queue = append(queue, Element{node: root.Left, fromLeft: true, zigZagLen: 1})
	}
	if root.Right != nil {
		queue = append(queue, Element{node: root.Right, fromLeft: false, zigZagLen: 1})
	}

	for len(queue) > 0 {
		// Pop element
		ele := queue[0]
		queue = queue[1:]
		// Update max zigzag length
		if ele.zigZagLen > maxZigZagLen {
			maxZigZagLen = ele.zigZagLen
		}

		if ele.node.Left != nil {
			if ele.fromLeft { // Came from left, no zigzag
				queue = append(queue, Element{node: ele.node.Left, fromLeft: true, zigZagLen: 1})
			} else { // Came from right, continue zigzag
				queue = append(queue, Element{node: ele.node.Left, fromLeft: true, zigZagLen: ele.zigZagLen + 1})
			}
		}

		if ele.node.Right != nil {
			if ele.fromLeft { // Came from left, continue zigzag
				queue = append(queue, Element{node: ele.node.Right, fromLeft: false, zigZagLen: ele.zigZagLen + 1})
			} else { // Came from right, no zigzag
				queue = append(queue, Element{node: ele.node.Right, fromLeft: false, zigZagLen: 1})
			}
		}
	}

	return maxZigZagLen
}
