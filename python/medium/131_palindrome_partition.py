"""
Given a string s, partition s such that every substring of the partition is a palindrome.
Return all possible palindrome partitioning of s.
"""

from pprint import pprint
from typing import List, Dict


class Solution:
    def partition(self, s: str) -> List[List[str]]:
        """O(n * 2^n) time and O(n^2) space, recursive"""
        def is_palindrome(s: str):
            """O(n) time and O(1) space"""
            for i in range(len(s) // 2):
                if s[i] != s[len(s) - 1 - i]:
                    return False
            return True

        def partition_recurse(s: str, memo: Dict = {}) -> List[List[str]]:
            """O(n * 2^n) time and O(n^2) space"""
            if not s:
                return [[]]
            if len(s) in memo:
                return memo[len(s)]

            res = []
            for i in range(len(s) - 1, -1, -1):
                curr = s[i:]
                if is_palindrome(curr):
                    tmp = partition_recurse(s[:i], memo)
                    res += [r + [curr] for r in tmp]
            memo[len(s)] = res
            return res

        return partition_recurse(s)


class Solution2:
    def partition(self, s: str) -> List[List[str]]:
        """O(n^4) time and O(n^2) space, iterative"""
        dp = [[] for _ in range(len(s) + 1)]
        dp[-1] = [[]]
        for i in range(len(s) - 1, -1, -1):
            for j in range(i + 1, len(s) + 1):
                curr = s[i:j]
                if curr == curr[::-1]:
                    for each in dp[j]:
                        dp[i].append([curr] + each)
        return dp[0]


if __name__ == "__main__":
    sol = Solution2()
    pprint(sol.partition("aaaa"))
