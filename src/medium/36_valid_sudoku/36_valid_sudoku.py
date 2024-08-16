"""
Determine if a 9x9 Sudoku board is valid. Only the filled cells need to be validated according to the following rules:
1. Each row must contain the digits 1-9 without repetition.
2. Each column must contain the digits 1-9 without repetition.
3. Each of the 9 3x3 sub-boxes of the grid must contain the digits 1-9 without repetition.
"""

from typing import List


class Solution:
    def isValidSudoku(self, board: List[List[str]]) -> bool:

        def checkRow(board: List[List[str]], row: int) -> bool:
            numbers = set()
            for x in board[row]:
                if not x.isdigit():
                    continue
                if x in numbers:
                    return False
                numbers.add(x)
            return True

        def checkCol(board: List[List[str]], col: int) -> bool:
            numbers = set()
            for j in range(len(board)):
                x = board[j][col]
                if not x.isdigit():
                    continue
                if x in numbers:
                    return False
                numbers.add(x)
            return True

        def checkSubBox(board: List[List[str]], box: int) -> bool:
            starting = ((box // 3) * 3, (box % 3) * 3)
            numbers = set()
            for j in range(3):
                for i in range(3):
                    x = board[starting[0] + j][starting[1] + i]
                    if not x.isdigit():
                        continue
                    if x in numbers:
                        return False
                    numbers.add(x)
            return True

        for i in range(9):
            if checkRow(board, i) and checkCol(board, i) and checkSubBox(board, i):
                continue
            return False
        return True


if __name__ == "__main__":
    sol = Solution()
    board = [
        [".", ".", ".", ".", "5", ".", ".", "1", "."],
        [".", "4", ".", "3", ".", ".", ".", ".", "."],
        [".", ".", ".", ".", ".", "3", ".", ".", "1"],
        ["8", ".", ".", ".", ".", ".", ".", "2", "."],
        [".", ".", "2", ".", "7", ".", ".", ".", "."],
        [".", "1", "5", ".", ".", ".", ".", ".", "."],
        [".", ".", ".", ".", ".", "2", ".", ".", "."],
        [".", "2", ".", "9", ".", ".", ".", ".", "."],
        [".", ".", "4", ".", ".", ".", ".", ".", "."]
    ]
    print(sol.isValidSudoku(board))
