package common

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func ArrayToTree(arr []any) *TreeNode {
	root, _ := ArrayToTreeWithNodes(arr)
	return root
}

func ArrayToTreeWithNodes(arr []any) (*TreeNode, []*TreeNode) {
	nodes := createNodes(arr)
	return arrayNode(nodes, 0), nodes
}

func createNodes(arr []any) []*TreeNode {
	nodes := make([]*TreeNode, 0, len(arr))
	for _, ele := range arr {
		if ele == nil {
			nodes = append(nodes, nil)
			continue
		}

		intVal, ok := ele.(int)
		if !ok {
			panic("invalid array element")
		}

		nodes = append(nodes, &TreeNode{Val: intVal})
	}

	return nodes
}

func arrayNode(nodes []*TreeNode, idx int) *TreeNode {
	if idx >= len(nodes) || nodes[idx] == nil {
		return nil
	}

	nodes[idx].Left = arrayNode(nodes, (idx+1)*2-1)
	nodes[idx].Right = arrayNode(nodes, (idx+1)*2)
	return nodes[idx]
}
