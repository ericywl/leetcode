import itertools
from typing import List


class Solution1:
    def generateParenthesis(self, n: int) -> List[str]:
        res = []

        def recurse(s: str, openParen: int, closeParen: int):
            if openParen + closeParen == 0:
                res.append(s)
                return
            if openParen > 0:
                recurse(s + "(", openParen - 1, closeParen)
            if closeParen > openParen:
                recurse(s + ")", openParen, closeParen - 1)

        recurse("", n, n)
        return res
        
class Solution2:
    def generateParenthesis(self, n: int) -> List[str]:
        if n == 0:
            return [""]

        res = []
        for i in range(n):
            for left in self.generateParenthesis(i):
                for right in self.generateParenthesis(n - i - 1):
                    res.append("(" + left + ")" + right)
        return res


if __name__ == "__main__":
    sol = Solution1()
    print(sol.generateParenthesis(4))
