package main

/*
Given an array of integers nums, calculate the pivot index of this array.

The pivot index is the index where the sum of all the numbers strictly to the left of the index
is equal to the sum of all the numbers strictly to the index's right.

If the index is on the left edge of the array, then the left sum is 0 because there are no
elements to the left. This also applies to the right edge of the array.

Return the leftmost pivot index. If no such index exists, return -1.
*/

func pivotIndex(nums []int) int {
	// Get the total sum
	rightSum := 0
	for _, n := range nums {
		rightSum += n
	}

	leftSum := 0
	// Loop through the numbers, exclude the current element from right sum.
	// If left and right sum are equal, return index.
	// Otherwise, continue and update the left sum.
	for i, n := range nums {
		rightSum -= n
		if leftSum == rightSum {
			return i
		}
		leftSum += n
	}

	return -1
}
