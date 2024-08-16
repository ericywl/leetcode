"""
Suppose an array sorted in ascending order is rotated at some pivot unknown to you beforehand.
(i.e., [0,1,2,4,5,6,7] might become [4,5,6,7,0,1,2]).
You are given a target value to search. If found in the array return its index, otherwise return -1.
You may assume no duplicate exists in the array.
Your algorithm's runtime complexity must be in the order of O(log n).
"""

from typing import List
import math


class Solution:
    def search(self, nums: List[int], target: int) -> int:
        if not nums:
            return -1
        start = 0
        end = len(nums) - 1
        while start <= end:
            mid = (start + end) // 2
            if target == nums[mid]:
                return mid
            if target < nums[mid]:
                # Left-half is sorted, but target is not inside
                if nums[start] <= nums[mid] and target < nums[start]:
                    start = mid + 1
                else:
                    end = mid - 1
            else:
                # Right-half is sorted, but target is not inside
                if nums[mid] <= nums[end] and target > nums[end]:
                    end = mid - 1
                else:
                    start = mid + 1
        return -1


if __name__ == "__main__":
    sol = Solution()
    arr = [4, 5, 6, 7, 0, 1, 2]
    print(sol.search(arr, 0))
