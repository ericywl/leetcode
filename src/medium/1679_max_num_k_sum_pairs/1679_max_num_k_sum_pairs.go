/**
 * You are given an integer array nums and an integer k.
 * In one operation, you can pick two numbers from the array whose sum equals k and remove them from the array.
 * Return the maximum number of operations you can perform on the array.
 */

package main

import (
	"slices"
)

func maxOperations(nums []int, k int) int {
	slices.Sort(nums)
	operations := 0
	i := 0
	j := len(nums) - 1

	for i < j {
		sum := nums[i] + nums[j]
		if sum == k {
			operations++
			i++
			j--
		} else if sum > k {
			j--
		} else {
			i++
		}
	}

	return operations
}
