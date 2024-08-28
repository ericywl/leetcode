package main

/*
Given two 0-indexed integer arrays nums1 and nums2, return a list answer of size 2 where:
- answer[0] is a list of all distinct integers in nums1 which are not present in nums2.
- answer[1] is a list of all distinct integers in nums2 which are not present in nums1.

Note that the integers in the lists may be returned in any order.
*/

func findDifference(nums1 []int, nums2 []int) [][]int {
	map1 := map[int]bool{}
	for _, n := range nums1 {
		map1[n] = true
	}

	map2 := map[int]bool{}
	for _, n := range nums2 {
		map2[n] = true
	}

	answer := make([][]int, 2)
	for n := range map1 {
		if !map2[n] {
			answer[0] = append(answer[0], n)
		}
	}

	for n := range map2 {
		if !map1[n] {
			answer[1] = append(answer[1], n)
		}
	}

	return answer
}
