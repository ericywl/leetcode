package main

import (
	"github.com/ericywl/leetcode/src/common"
)

func oddEvenList(head *common.ListNode) *common.ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	// Two tail pointers for the odd and even node groups
	odd := head
	even := head.Next
	// Keep the even head since we need the odd tail to point to it
	evenHead := head.Next

	for even != nil && even.Next != nil {
		odd.Next = even.Next
		odd = odd.Next
		even.Next = odd.Next
		even = even.Next
	}

	// Point the odd tail to even head
	odd.Next = evenHead
	return head
}
