package main

import (
	"sort"
	"strconv"
	"strings"
)

// largestNumber will arrange the given list of non-negative integers such that they form the largest number
// and returns it as a string.
// E.g. nums = [10,2] => "210"
//
//	nums = [3,30,34,5,9] => "9534330"
func largestNumber(nums []int) string {
	strs := make([]string, 0, len(nums))
	for _, n := range nums {
		strs = append(strs, strconv.Itoa(n))
	}

	sort.Slice(strs, func(i, j int) bool {
		// Test both combinations to see which should go in-front i.e. create a larger number
		order1 := strs[i] + strs[j]
		order2 := strs[j] + strs[i]
		return strings.Compare(order1, order2) > 0
	})

	if strs[0] == "0" {
		return "0"
	}

	return strings.Join(strs, "")
}
