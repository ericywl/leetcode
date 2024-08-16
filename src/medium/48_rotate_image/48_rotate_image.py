"""
You are given an n x n 2D matrix representing an image.
Rotate the image by 90 degrees (clockwise).
Note:
You have to rotate the image in-place, which means you have to modify the input 2D matrix directly. DO NOT allocate another 2D matrix and do the rotation.
"""


from typing import List

class Solution:
    def rotate(self, matrix: List[List[int]]) -> None:
        """O(n^2) time and O(1) space"""
        n = len(matrix)
        for j in range(n // 2):
            for i in range(j, len(matrix[j]) - j - 1):
                tmp = matrix[j][i]
                matrix[j][i] = matrix[n - i - 1][j]
                matrix[n - i - 1][j] = matrix[n - j - 1][n - i - 1]
                matrix[n - j - 1][n - i - 1] = matrix[i][n - j - 1]
                matrix[i][n - j - 1] = tmp


if __name__ == "__main__":
    sol = Solution()
    matrix = [
        [5, 1, 9, 11],
        [2, 4, 8, 10],
        [13, 3, 6, 7],
        [15, 14, 12, 16]
    ]
    sol.rotate(matrix)
    print(matrix)