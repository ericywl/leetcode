"""
Given an array nums of n integers, are there elements a, b, c in nums such that a + b + c = 0? Find all unique triplets in the array which gives the sum of zero.

Note:

The solution set must not contain duplicate triplets.
"""

from typing import List


class Solution:
    def threeSum(self, nums: List[int]) -> List[List[int]]:
        nums.sort()
        res = []
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
                if summ == 0:
                    res.append([n1, n2, n3])
                    j += 1
                    k -= 1
                    # Skip duplicate numbers
                    while j < k and nums[j] == nums[j - 1]:
                        j += 1
                    while j < k and nums[k] == nums[k + 1]:
                        k -= 1
                # Summation more than zero, decrease back pointer
                elif summ > 0:
                    k -= 1
                # Summation less than zero, increased front pointer
                else:
                    j += 1
        return res


if __name__ == "__main__":
    sol = Solution()
    print(sol.threeSum([-2, 0, 0, 2, 2]))
