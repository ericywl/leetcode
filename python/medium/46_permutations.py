"""
Given a collection of distinct integers, return all possible permutations.
"""

from typing import List

class Solution:
    def permute(self, nums: List[int]) -> List[List[int]]:
        if len(nums) == 1:
            return [nums]
        if len(nums) == 2:
            return [nums, nums[::-1]]
        results = []
        for num in nums:
            new_nums = [n for n in nums if n != num]
            for tmp in self.permute(new_nums):
                results.append([num] + tmp)
        return results
        
        
if __name__ == "__main__":
    sol = Solution()
    print(sol.permute([1, 2, 3]))
