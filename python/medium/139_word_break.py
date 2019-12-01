"""
Given a non-empty string s and a dictionary wordDict containing a list of non-empty words, determine if s can be segmented into a space-separated sequence of one or more dictionary words.

Note:
 - The same word in the dictionary may be reused multiple times in the segmentation.
 - You may assume the dictionary does not contain duplicate words.
"""

from typing import List


class Solution:
    def wordBreak(self, s: str, wordDict: List[str]) -> bool:
        """O(min(nm, n^2)) time and O(n) space"""
        memo = [False] * (len(s) + 1)
        memo[0] = True
        wordSet = frozenset(wordDict)
        if len(s) > len(wordDict):
            # O(nm) time
            for i in range(len(s)):
                for word in wordSet:
                    if s[i:i + len(word)] == word and memo[i]:
                        memo[i + len(word)] = True
        else:
            # O(n^2) time
            for i in range(1, len(s) + 1):
                for j in range(i):
                    if memo[j] and s[j:i] in wordSet:
                        memo[i] = True
        return memo[-1]
        
        
def test1(sol):
    s = "applepenapple"
    wordDict = ["apple", "pen"]
    print(sol.wordBreak(s, wordDict))


def test2(sol):
    s = "leetcode"
    wordDict = ["leet", "code"]
    print(sol.wordBreak(s, wordDict))


def test3(sol):
    s = "catsandog"
    wordDict = ["cats", "dog", "sand", "and", "cat"]
    print(sol.wordBreak(s, wordDict))


if __name__ == "__main__":
    sol = Solution()
    test1(sol)
    test2(sol)
    test3(sol)
