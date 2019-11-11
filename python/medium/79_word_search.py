"""
Given a 2D board and a word, find if the word exists in the grid.
The word can be constructed from letters of sequentially adjacent cell, where "adjacent" cells 
are those horizontally or vertically neighboring. The same letter cell may not be used more than once.
"""

from typing import List, Tuple, Set


class Solution:
    def exist(self, board: List[List[str]], word: str) -> bool:
        """O(mn * 4^s) time and O(s) space, s being the length of the word"""
        def existRecurse(board: List[List[str]], j: int, i: int,
                         word: str, visited_indices: Set[Tuple[int, int]]) -> bool:
            if len(word) == 0:
                return True
            if i < 0 or j < 0:
                return False
            if j >= len(board) or i >= len(board[0]):
                return False
            if word[0] != board[j][i]:
                return False
            if (j, i) in visited_indices:
                return False
            for next_idx in [(j + 1, i), (j - 1, i), (j, i + 1), (j, i - 1)]:
                visited_indices.add((j, i))
                if existRecurse(board, next_idx[0], next_idx[1], word[1:], visited_indices):
                    return True
                visited_indices.remove((j, i))
            return False

        if not board or not word:
            return False
        for j in range(len(board)):
            for i in range(len(board[j])):
                if existRecurse(board, j, i, word, set()):
                    return True
        return False


if __name__ == "__main__":
    sol = Solution()
    matrix = [["A", "B", "C", "E"], ["S", "F", "C", "S"], ["A", "D", "E", "E"]]
    word = "ABCCED"
    print(sol.exist(matrix, word))
