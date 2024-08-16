"""
A robot is located at the top-left corner of a m x n grid (marked 'Start' in the diagram below).
The robot can only move either down or right at any point in time. The robot is trying to reach the bottom-right corner of the grid (marked 'Finish' in the diagram below).
How many possible unique paths are there?
"""


class Solution:
    memo = {}

    def uniquePaths(self, m: int, n: int) -> int:
        """O(mn) time and O(mn) space, due to memoization"""
        if not m or not n:
            return 0
        if m == 1 or n == 1:
            return 1
        if (m, n) in self.memo:
            return self.memo[(m, n)]
        if (n, m) in self.memo:
            return self.memo[(n, m)]
        a = self.uniquePaths(m - 1, n)
        self.memo[(m - 1, n)] = a
        b = self.uniquePaths(m, n - 1)
        self.memo[(m, n - 1)] = b
        return a + b


if __name__ == "__main__":
    sol = Solution()
    print(sol.uniquePaths(23, 12))
