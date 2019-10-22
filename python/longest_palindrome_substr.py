"""
Given a string s, find the longest palindromic substring in s. You may assume that the maximum length of s is 1000.
"""


class Solution:
    def addBorders(self, s: str) -> str:
        return "#" + "#".join([c for c in s]) + "#"

    def removeBorders(self, s: str) -> str:
        return "".join(s.split("#"))

    def longestPalindrome(self, s: str) -> str:
        if not s:
            return ""

        pali = s[0]
        s = self.addBorders(s)
        for i in range(1, len(s) - 1):
            for l in range(i - 1, -1, -1):
                r = 2 * i - l
                if r >= len(s) or s[l] != s[r]:
                    break
                pali = s[l:r + 1] if len(s[l:r + 1]) > len(pali) else pali
        return self.removeBorders(pali)


if __name__ == "__main__":
    sol = Solution()
    print(sol.longestPalindrome("adfnwejrdabbad"))
