package main

import (
	"github.com/ericywl/leetcode/src/common"
)

func pairSum(head *common.ListNode) int {
	middle, tail := tortoiseAndHare(head)

	// Reverse the second half of the linked list
	var prev *common.ListNode
	curr := middle // middle.Next is not nil because the constraint is that there are >= 2 nodes
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}

	// Now that we have pointers at the head and tail, and the back half is reversed,
	// we can calculate twin sums while progressing the pointers
	maxTwinSum := 0
	start, end := head, tail
	for start != nil {
		twinSum := start.Val + end.Val
		if twinSum > maxTwinSum {
			maxTwinSum = twinSum
		}

		start = start.Next
		end = end.Next
	}

	return maxTwinSum
}

func tortoiseAndHare(head *common.ListNode) (middle *common.ListNode, tail *common.ListNode) {
	slow := head
	fast := head

	// Keep progressing forward until fast can no longer progress by 2 steps
	for fast != nil && fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	// Make sure that fast is always the tail by continuously progressing it to the end
	for fast.Next != nil {
		fast = fast.Next
	}

	return slow, fast
}
