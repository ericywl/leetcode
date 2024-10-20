package main

import (
	"container/heap"
)

type IntHeap []int

func (h IntHeap) Len() int {
	return len(h)
}

func (h IntHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *IntHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func findKthLargest(nums []int, k int) int {
	// Initialize k elements into heap first.
	h := IntHeap(nums[:k])
	heap.Init(&h)

	// For each remaining num, we check if it is larger than the smallest num in heap.
	// If it is, we overwrite the smallest num and fix the heap.
	for _, num := range nums[k:] {
		if num > h[0] {
			// Note: This is equivalent to popping and then pushing the new num.
			// heap.Pop(&h)
			// heap.Push(&h, num)
			h[0] = num
			heap.Fix(&h, 0)
		}
	}

	// Since our heap is of size k, the smallest num in the heap in the end is the k-th largest.
	return h[0]
}
