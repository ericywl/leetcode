"""
Given a string containing digits from 2-9 inclusive, return all possible letter combinations that the number could represent.

A mapping of digit to letters (just like on the telephone buttons) is given below. Note that 1 does not map to any letters.
"""

from typing import List
import itertools


class Solution:
    mappings = {
        "2": "abc",
        "3": "def",
        "4": "ghi",
        "5": "jkl",
        "6": "mno",
        "7": "pqrs",
        "8": "tuv",
        "9": "wxyz"
    }

    def letterCombinations(self, digits: str) -> List[str]:
        if not digits.isdigit() or any(d in digits for d in ["0", "1"]):
            return []
        tmp = [self.mappings[d] for d in digits]
        res = ["".join(a) for a in itertools.product(*tmp)]
        return res

if __name__ == "__main__":
    sol = Solution()
    print(sol.letterCombinations("23"))
