package main

import "github.com/ericywl/leetcode/src/common"

func deleteMiddle(head *common.ListNode) *common.ListNode {
	var previousSlow *common.ListNode
	slow := head
	fast := head
	for fast != nil && fast.Next != nil {
		previousSlow = slow
		slow = slow.Next
		fast = fast.Next.Next
	}

	if previousSlow == nil {
		return nil
	}
	if slow != nil {
		previousSlow.Next = slow.Next
		slow = nil
	}
	return head
}
