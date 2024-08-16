package medium

// longestConsecutive returns the length of the longest consecutive elements sequence given unsorted list of integers.
// E.g. nums = [100,4,200,1,3,2] => 4
//		nums = [0,3,7,2,5,8,4,6,0,1] => 9
func longestConsecutive(nums []int) int {
	numsMap := map[int]bool{}
	for _, n := range nums {
		numsMap[n] = true
	}

	longestStreak := 0
	for n := range numsMap {
		// If the n-1 exists, the number is not the smallest in the sequence, so we ignore
		if !numsMap[n-1] {
			currNum := n
			currStreak := 1

			// Keep iterating and checking if the current number + 1 exists
			for numsMap[currNum+1] {
				currNum += 1
				currStreak += 1
			}

			if currStreak > longestStreak {
				longestStreak = currStreak
			}
		}
	}

	return longestStreak
}
