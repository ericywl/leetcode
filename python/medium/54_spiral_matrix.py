"""
Given a matrix of m x n elements (m rows, n columns), return all elements of the matrix in spiral order.
"""

from typing import List
import math


class Solution:
    def spiralOrder(self, matrix: List[List[int]]) -> List[int]:
        if not matrix:
            return []
        result = []
        w, h = len(matrix[0]), len(matrix)
        for k in range(math.ceil(max(w, h) / 2)):
            i, j = k, k
            tmp_w, tmp_h = w - 2 * i, h - 2 * j
            if tmp_w < 1 or tmp_h < 1:
                break
            # Move rightwards on 1st row till (k + tmp_w - 1)
            for i in range(i, i + tmp_w):
                result.append(matrix[j][i])
            # Move down from 1st row till (k + tmp_h - 1)
            # Note: after previous loop, i is now (k + tmp_w - 1)
            for j in range(j + 1, j + tmp_h):
                result.append(matrix[j][i])
            # If the remaining matrix is flat,
            # we don't need to loop back
            if tmp_w > 1 and tmp_h > 1:
                # Move leftwards from last element in remaining matrix
                # Note: i is still(k + tmp_w - 1) here
                for i in range(i - 1, k - 1, -1):
                    result.append(matrix[j][i])
                # Move upwards after reaching start of remaining matrix
                # Note: i is now k
                for j in range(j - 1, k, -1):
                    result.append(matrix[j][i])
        return result


if __name__ == "__main__":
    sol = Solution()
    matrix1 = [
        [1, 2, 3, 4, 5],
        [6, 7, 8, 9, 10],
        [11, 12, 13, 14, 15],
        [16, 17, 18, 19, 20],
        [21, 22, 23, 24, 25]
    ]
    matrix2 = [
        [1, 2, 3, 4],
        [5, 6, 7, 8],
        [9, 10, 11, 12]
    ]
    print(sol.spiralOrder(matrix1))
    print(sol.spiralOrder(matrix2))
