class Solution:
    def lengthOfLongestSubstring(self, s: str) -> int:
        chars = {}
        res = i = 0
        for j in range(0, len(s)):
            if s[j] in chars:
                i = max(chars[s[j]] + 1, i)
            res = max(res, j - i + 1)
            chars[s[j]] = j
        return res


if __name__ == "__main__":
    sol = Solution()
    print(sol.lengthOfLongestSubstring("tmmzuxt"))
    print(sol.lengthOfLongestSubstring("pwzwew"))
