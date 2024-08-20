package main

func longestOnes(nums []int, k int) int {
	windowStart := 0
	windowEnd := 0

	for ; windowEnd < len(nums); windowEnd++ {
		// The end of the window is always moving forward.
		// If we encounter a zero, we flip it by decrementing k.
		if nums[windowEnd] == 0 {
			k--
		}

		// We decremented k but found that we have actually used up all the flips.
		// What we can do is to continue moving the window rightwards without shrinking / expanding it.
		// This means that we will only expand the window if we can find one that is larger
		// than the current one.
		if k < 0 {
			if nums[windowStart] == 0 {
				// Recover back 1 flip if the start of window is 0.
				k++
			}
			// Move the window rightwards (instead of expanding).
			windowStart++
		}

		// Otherwise, we still have more flips left, so we can expand the window.
	}

	return windowEnd - windowStart
}
