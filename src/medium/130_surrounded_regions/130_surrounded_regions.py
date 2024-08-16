"""
Given a 2D board containing 'X' and 'O' (the letter O), capture all regions surrounded by 'X'.
A region is captured by flipping all 'O's into 'X's in that surrounded region.

Surrounded regions shouldn’t be on the border, which means that any 'O' on the border of the 
board are not flipped to 'X'. Any 'O' that is not on the border and it is not connected to an 'O' 
on the border will be flipped to 'X'. Two cells are connected if they are adjacent cells 
connected horizontally or vertically.
"""


from pprint import pprint
from typing import List


class Solution:
    def solve(self, board: List[List[str]]) -> None:
        """
        Do not return anything, modify board in-place instead.
        """
        if not any(board):
            return

        tmp = []
        for j in range(len(board)):
            for i in range(len(board[j])):
                if i == 0 or j == 0 or i == len(board[j]) - 1 or j == len(board) - 1:
                    tmp.append((j, i))
        while tmp:
            j, i = tmp.pop()
            if j < 0 or j >= len(board):
                continue
            if i < 0 or i >= len(board[j]):
                continue
            if board[j][i] == "O":
                board[j][i] = "#"
                tmp += [(j - 1, i), (j + 1, i), (j, i - 1), (j, i + 1)]
        
        for j in range(len(board)):
            for i in range(len(board[j])):
                if board[j][i] == "#":
                    board[j][i] = "O"
                else:
                    board[j][i] = "X"


if __name__ == "__main__":
    sol = Solution()
    board = [["X", "O", "X", "O", "X", "O", "O", "O", "X", "O"],
             ["X", "O", "O", "X", "X", "X", "O", "O", "O", "X"],
             ["O", "O", "O", "O", "O", "O", "O", "O", "X", "X"],
             ["O", "O", "O", "O", "O", "O", "X", "O", "O", "X"],
             ["O", "O", "X", "X", "O", "X", "X", "O", "O", "O"],
             ["X", "O", "O", "X", "X", "X", "O", "X", "X", "O"],
             ["X", "O", "X", "O", "O", "X", "X", "O", "X", "O"],
             ["X", "X", "O", "X", "X", "O", "X", "O", "O", "X"],
             ["O", "O", "O", "O", "X", "O", "X", "O", "X", "O"],
             ["X", "X", "O", "X", "X", "X", "X", "O", "O", "O"]]
    expected = [["X", "O", "X", "O", "X", "O", "O", "O", "X", "O"],
                ["X", "O", "O", "X", "X", "X", "O", "O", "O", "X"],
                ["O", "O", "O", "O", "O", "O", "O", "O", "X", "X"],
                ["O", "O", "O", "O", "O", "O", "X", "O", "O", "X"],
                ["O", "O", "X", "X", "O", "X", "X", "O", "O", "O"],
                ["X", "O", "O", "X", "X", "X", "X", "X", "X", "O"],
                ["X", "O", "X", "X", "X", "X", "X", "O", "X", "O"],
                ["X", "X", "O", "X", "X", "X", "X", "O", "O", "X"],
                ["O", "O", "O", "O", "X", "X", "X", "O", "X", "O"],
                ["X", "X", "O", "X", "X", "X", "X", "O", "O", "O"]]
    sol.solve(board)
    pprint(board)
    if board != expected:
        for j in range(len(board)):
            for i in range(len(board[j])):
                if board[j][i] != expected[j][i]:
                    print(j, i)
