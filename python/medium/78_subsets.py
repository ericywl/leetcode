"""
Given a set of distinct integers, nums, return all possible subsets (the power set).
Note: The solution set must not contain duplicate subsets.
"""

import itertools
from typing import List


class Solution:
    def subsets(self, nums: List[int]) -> List[List[int]]:
        res = []
        for i in range(len(nums) + 1):
            res += list(map(list, itertools.combinations(nums, i)))
        return res


if __name__ == "__main__":
    sol = Solution()
    nums = [1, 2, 3]
    print(sol.subsets(nums))
