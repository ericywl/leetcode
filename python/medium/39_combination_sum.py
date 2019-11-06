"""
Given a set of candidate numbers (candidates) (without duplicates) and a target number (target), 
find all unique combinations in candidates where the candidate numbers sums to target.
The same repeated number may be chosen from candidates unlimited number of times.
"""

from typing import List


class Solution:
    def combinationSum(self, candidates: List[int], target: int) -> List[List[int]]:
        """O(n) time and O(n) space"""
        def combiRecurse(candidates: List[int], start_idx: int, target: int) -> List[List[int]]:
            if target <= 0:
                return []

            combinations = []
            for i, x in enumerate(candidates):
                if i < start_idx:
                    continue
                if x > target:
                    break
                if x == target:
                    combinations.append([x])
                else:
                    for results in combiRecurse(candidates, i, target - x):
                        combinations.append([x] + results)
            return combinations

        candidates.sort()
        return combiRecurse(candidates, 0, target)


if __name__ == "__main__":
    sol = Solution()
    candidates = [2, 3, 6, 7]
    target = 8
    print(sol.combinationSum(candidates, target))
