"""
Given an array with n objects colored red, white or blue, sort them in-place so that objects of the same color are adjacent, with the colors in the order red, white and blue.
Here, we will use the integers 0, 1, and 2 to represent the color red, white, and blue respectively.
Note: You are not suppose to use the library's sort function for this problem.
"""

from typing import List


class Solution:
    def sortColors(self, nums: List[int]) -> None:
        """O(n) time and O(1) space"""
        if not nums:
            return
        cidx = [0, 0, len(nums) - 1]
        while cidx[1] <= cidx[2]:
            if nums[cidx[1]] == 0:
                nums[cidx[1]], nums[cidx[0]] = nums[cidx[0]], nums[cidx[1]]
                cidx[0] += 1
                cidx[1] += 1
            elif nums[cidx[1]] == 1:
                cidx[1] += 1
            elif nums[cidx[1]] == 2:
                nums[cidx[1]], nums[cidx[2]] = nums[cidx[2]], nums[cidx[1]]
                cidx[2] -= 1


if __name__ == "__main__":
    sol = Solution()
    nums = [1, 2, 2, 2, 2, 0, 0, 0, 1, 1]
    sol.sortColors(nums)
    print(nums)
