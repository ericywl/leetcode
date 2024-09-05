package main

/*
Given an array of integers arr, return true if the number of occurrences of each value in
the array is unique or false otherwise.
*/

func uniqueOccurrences(arr []int) bool {
	numOcc := map[int]int{}
	for _, n := range arr {
		numOcc[n]++
	}

	uniqueOcc := map[int]bool{}
	for _, occ := range numOcc {
		if uniqueOcc[occ] {
			return false
		}

		uniqueOcc[occ] = true
	}

	return true
}
