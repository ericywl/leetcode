"""
Given an array of non-negative integers, you are initially positioned at the first index of the array.
Each element in the array represents your maximum jump length at that position.
Determine if you are able to reach the last index.
"""

from typing import List


class Solution:
    def canJump(self, nums: List[int]) -> bool:
        """O(n) time and O(1) space"""
        if not nums:
            return False
        if len(nums) == 1:
            return True

        i = len(nums) - 1
        j = i
        while j >= 0:
            if j + nums[j] >= i:
                i = j
            j -= 1
        return i == 0


if __name__ == "__main__":
    sol = Solution()
    arr = [3, 0, 8, 2, 0, 0, 1]
    print(sol.canJump(arr))
