package common

type ListNode struct {
	Val  int
	Next *ListNode
}

func ArrayToListNodes(arr []int) *ListNode {
	var curr *ListNode
	for i := len(arr) - 1; i >= 0; i-- {
		node := &ListNode{Val: arr[i], Next: curr}
		curr = node
	}

	return curr
}

func ListNodesToArray(head *ListNode) []int {
	var arr []int
	for head != nil {
		arr = append(arr, head.Val)
		head = head.Next
	}
	return arr
}
