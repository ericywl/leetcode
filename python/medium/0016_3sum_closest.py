"""
Given an array nums of n integers and an integer target, find three integers in nums such that the sum is closest to target. Return the sum of the three integers. You may assume that each input would have exactly one solution.
"""

from typing import List


class Solution:
    def threeSumClosest(self, nums: List[int], target: int) -> int:
        nums.sort()
        res = nums[0] + nums[1] + nums[2]
        for i in range(0, len(nums) - 2):
            n1 = nums[i]
            # Skip duplicate numbers
            if i > 0 and n1 == nums[i - 1]:
                continue
            j = i + 1
            k = len(nums) - 1
            while j < k:
                n2 = nums[j]
                n3 = nums[k]
                summ = n1 + n2 + n3
                # Summation is zero, add to result
                if abs(summ - target) < abs(res - target):
                    res = summ
                if summ == target:
                    return target
                # Summation more than zero, decrease back pointer
                if summ > target:
                    k -= 1
                # Summation less than zero, increased front pointer
                else:
                    j += 1
        return res
