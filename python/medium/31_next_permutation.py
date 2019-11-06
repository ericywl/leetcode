"""
Implement next permutation, which rearranges numbers into the lexicographically next greater permutation of numbers.
If such arrangement is not possible, it must rearrange it as the lowest possible order (ie, sorted in ascending order).
The replacement must be in-place and use only constant extra memory.
"""

from typing import List


class Solution:
    def nextPermutation(self, nums: List[int]) -> None:
        """
        O(n log n) time with O(1) space
        * The worse time is due to unnecessary sorting, see below
        """
        lim = -1
        for i in range(len(nums) - 1, 0, -1):
            if nums[i] > nums[i - 1]:
                lim = i - 1
                break
        if lim == -1:
            nums.sort()
            return
        next_smallest = float('inf')
        ns_idx = len(nums)
        for i in range(len(nums) - 1, lim, -1):
            if nums[i] > nums[lim] and nums[i] < next_smallest:
                next_smallest = nums[i]
                ns_idx = i
        if ns_idx < len(nums):
            nums[lim], nums[ns_idx] = nums[ns_idx], nums[lim]
            nums[lim + 1:] = sorted(nums[lim + 1:])


class Solution2:
    def nextPermutation(self, nums: List[int]) -> None:
        """O(n) time with O(1) space"""
        x = 0
        for x in range(len(nums) - 2, -2, -1):
            if nums[x + 1] > nums[x]:
                break
        if x < 0:
            nums.reverse()
            return
        y = 0
        for y in range(len(nums) - 1, x, -1):
            if nums[y] > nums[x]:
                nums[x], nums[y] = nums[y], nums[x]
                break
        nums[x + 1:] = nums[x + 1:][::-1]


if __name__ == "__main__":
    sol = Solution2()
    # arr = [1, 4, 3, 2, 1]
    arr = [3, 2, 1]
    sol.nextPermutation(arr)
    print(arr)
