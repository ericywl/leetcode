"""
Given an array of integers nums sorted in ascending order, find the starting and ending position of a given target value.
Your algorithm's runtime complexity must be in the order of O(log n).
If the target is not found in the array, return [-1, -1].
"""


class Solution:
    def searchRange(self, nums: List[int], target: int) -> List[int]:
        # Finding first index
        start, end = 0, len(nums) - 1
        result = [-1, -1]
        while start <= end:
            mid = (start + end) // 2
            if target == nums[mid]:
                result[0] = mid
                end = mid - 1
            elif target < nums[mid]:
                end = mid - 1
            else:
                start = mid + 1
        # Finding last index
        start, end = 0, len(nums) - 1
        while start <= end:
            mid = (start + end) // 2
            if target == nums[mid]:
                result[1] = mid
                start = mid + 1
            elif target < nums[mid]:
                end = mid - 1
            else:
                start = mid + 1
        return result
