"""
A message containing letters from A-Z is being encoded to numbers using the following mapping:

    'A' -> 1
    'B' -> 2
    ...
    'Z' -> 26

Given a non-empty string containing only digits, determine the total number of ways to decode it.
"""


class Solution:
    """O(n) time and O(n) space"""
    def numDecodings(self, s: str) -> int:
        if not s:
            return 0

        def isValid(s: str, i: int):
            if i >= len(s):
                return False
            return s[i] == "1" or (s[i] == "2" and s[i + 1] < "7")

        memo = [0] * (len(s) + 1)
        memo[len(s)] = 1
        for i in range(len(s) - 1, -1, -1):
            if s[i] != "0":
                memo[i] = memo[i + 1]
                if i < len(s) - 1 and isValid(s, i):
                    memo[i] += memo[i + 2]
        return memo[0]        


class Solution2:
    def numDecodings(self, s: str) -> int:
        """O(2^n) time, too slow"""
        if not s:
            return 0
        if s[0] == "0":
            return 0
        if len(s) == 1:
            return 1
        if len(s) == 2:
            if int(s) > 26:
                return 0 if s[1] == "0" else 1
            return 1 if s[1] == "0" else 2
        val = 0
        if s[1] != "0":
            val += self.numDecodings(s[1:])
        if s[:2] < "27":
            val += self.numDecodings(s[2:])
        return val


if __name__ == "__main__":
    sol = Solution()
    print(sol.numDecodings("302"))
