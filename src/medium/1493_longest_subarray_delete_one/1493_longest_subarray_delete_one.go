package main

func longestSubarray(nums []int) int {
	windowStart := 0
	windowEnd := 0
	deletes := 1

	for ; windowEnd < len(nums); windowEnd++ {
		if nums[windowEnd] == 0 {
			deletes--
		}

		if deletes < 0 {
			if nums[windowStart] == 0 {
				deletes++
			}
			windowStart++
		}
	}

	return windowEnd - windowStart - 1
}
