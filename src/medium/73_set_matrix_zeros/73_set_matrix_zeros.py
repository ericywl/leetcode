"""
Given a m x n matrix, if an element is 0, set its entire row and column to 0. Do it in-place.
"""

from typing import List


class Solution:
    def setZeroes(self, matrix: List[List[int]]) -> None:
        """O(nm) time and O(1) space"""
        def setZeroesCol(matrix: List[List[int]], col: int) -> None:
            for j in range(len(matrix)):
                matrix[j][col] = 0

        def setZeroesRow(matrix: List[List[int]], row: int) -> None:
            for i in range(len(matrix[0])):
                matrix[row][i] = 0

        first_row_zero = any(x == 0 for x in matrix[0])
        first_col_zero = any(matrix[j][0] == 0 for j in range(len(matrix)))
        # Store zero decisions in first row and first column
        for j in range(1, len(matrix)):
            for i in range(1, len(matrix[j])):
                if matrix[j][i] == 0:
                    matrix[j][0] = 0
                    matrix[0][i] = 0
        # Set columns to zero
        for i in range(1, len(matrix[0])):
            if matrix[0][i] == 0:
                setZeroesCol(matrix, i)
        # Set rows to zero
        for j in range(1, len(matrix)):
            if matrix[j][0] == 0:
                setZeroesRow(matrix, j)
        # Zero out first row and column if necessary
        if first_row_zero:
            setZeroesRow(matrix, 0)
        if first_col_zero:
            setZeroesCol(matrix, 0)


if __name__ == "__main__":
    sol = Solution()
    matrix = [[0, 1, 2, 0], [3, 4, 5, 2], [1, 3, 1, 5]]
    sol.setZeroes(matrix)
    print(matrix)
