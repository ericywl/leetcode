package main

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	var result []int
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		qLen := len(queue)
		// Append last node value to result
		result = append(result, queue[qLen-1].Val)

		for i := 0; i < qLen; i++ {
			// Pop node from front of queue
			node := queue[0]
			queue = queue[1:]

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
	}

	return result
}
