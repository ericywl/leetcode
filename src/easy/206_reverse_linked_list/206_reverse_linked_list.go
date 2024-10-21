package main

import (
	"github.com/ericywl/leetcode/src/common"
)

func reverseList(head *common.ListNode) *common.ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	prev, curr := head, head.Next
	// Make sure previous node points to nothing
	prev.Next = nil
	for curr != nil {
		tmp := curr.Next
		curr.Next = prev
		prev = curr
		curr = tmp
	}

	return prev
}

// With the example (1) -> (2) -> (3), we can visualize what happens.
//
// Round 1:
// This is when the call stack reaches the end i.e. (3), since (3) points to nil.
// Final: (3)
//
// Round 2:
// Since (3) is returned from our call, `node` is (3) and `head` is (2).
// Begin: (2H) -> (3N)
// We point `head`'s next node towards `head`, in this case that next node is (3).
// So we are pointing (3) towards (2).
// Next:  (2H) <-> (3N)
// Finally, we make sure that `head` points to nil.
// Final: (2H) <- (3N)
//
// Round 3:
// Now, `node` is still (3) but `head` is (1).
// Begin: (1H) -> (2) <- (3N)
// We do the same thing as we have done before.
// Next:  (1H) <-> (2) <- (3N)
// Final: (1H) <- (2) <- (3N)
func reverseListRecursive(head *common.ListNode) *common.ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	node := reverseListRecursive(head.Next)
	head.Next.Next = head
	head.Next = nil
	return node
}
