package medium

// canJump solves the problem below:
// You are initially positioned at the array's first index, and each element in the array represents your maximum
// jump length at that position.
// Return true if you can reach the last index, else false.
// E.g. nums = [2,3,1,1,4] => true
//		nums = [3,2,1,0,4] => false
func canJump(nums []int) bool {
	if len(nums) == 0 {
		return false
	}

	if len(nums) == 1 {
		// Already at the end
		return true
	}

	// Start from the end of the list and check if any steps in-front of it can reach it
	// If reachable, use that index as the new ending and repeat the process
	// The algorithm will terminate after we reach the beginning of the list
	currEnd := len(nums) - 1
	currStart := currEnd
	for currStart >= 0 {
		if currStart+nums[currStart] >= currEnd {
			currEnd = currStart
		}
		currStart -= 1
	}

	// If by the time we reach the beginning of the list, our existing end pointer also reaches the beginning,
	// it means that the first element in the list can reach the end by jumping
	return currEnd == 0
}
