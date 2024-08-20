package main

func findMaxAverage(nums []int, k int) float64 {
	maxSum := sum(nums[0:k])
	currSum := maxSum
	for i := k; i < len(nums); i++ {
		currSum = currSum - nums[i-k] + nums[i]
		if currSum > maxSum {
			maxSum = currSum
		}
	}

	return float64(maxSum) / float64(k)
}

func sum(nums []int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return sum
}
