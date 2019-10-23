"""
Given an array nums of n integers and an integer target, are there elements a, b, c, and d in nums such that a + b + c + d = target? Find all unique quadruplets in the array which gives the sum of target.
"""

from typing import List


class Solution1:
    """
    Make use of previously coded 3sum to do 4sum
    """

    def threeSum(self, nums: List[int], target: int) -> List[List[int]]:
        # Assume sorted list
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
                if summ == target:
                    res.append([n1, n2, n3])
                    j += 1
                    k -= 1
                    # Skip duplicate numbers
                    while j < k and nums[j] == nums[j - 1]:
                        j += 1
                    while j < k and nums[k] == nums[k + 1]:
                        k -= 1
                # Summation more than zero, decrease back pointer
                elif summ > target:
                    k -= 1
                # Summation less than zero, increased front pointer
                else:
                    j += 1
        return res

    def fourSum(self, nums: List[int], target: int) -> List[List[int]]:
        nums.sort()
        res = []
        for i in range(len(nums) - 1, 2, -1):
            if i < len(nums) - 1 and nums[i] == nums[i + 1]:
                continue
            new_target = target - nums[i]
            tmp = self.threeSum(nums[:i], new_target)
            for arr in tmp:
                res.append(arr + [nums[i]])
        return res


class Solution2:
    """
    Pre-compute all possible 2sums and store in dictionary, 
    then go through them to find 4sums
    """

    def fourSum(self, nums: List[int], target: int) -> List[List[int]]:
        nums.sort()
        twoSums = {}
        for i in range(0, len(nums) - 1):
            for j in range(i + 1, len(nums)):
                summ = nums[i] + nums[j]
                if summ not in twoSums:
                    twoSums[summ] = set()
                # Use frozenset to ensure no duplicates like (i, j) and (j, i)
                twoSums[summ].add(frozenset((i, j)))
        # This part can be improved on, using two-for loops on the nums
        res = set()
        for k in list(twoSums.keys()):
            diff = target - k
            if diff in twoSums:
                for a in twoSums[k]:
                    for b in twoSums[diff]:
                        if not any(x == y for x in a for y in b):
                            fs = a.union(b)
                            # Use sorted tuple to ensure no duplicates
                            res.add(tuple(sorted(nums[idx] for idx in fs)))
            del twoSums[k]
        return [list(x) for x in res]


if __name__ == "__main__":
    sol = Solution2()
    print(sol.fourSum([0, 0, 0, 0], 0))
    print(sol.fourSum([5, 5, 3, 5, 1, -5, 1, -2], 4))
    print(sol.fourSum([0, 4, -5, 2, -2, 4, 2, -1, 4], 12))
