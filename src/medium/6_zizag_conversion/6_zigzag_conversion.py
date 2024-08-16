"""
The string "PAYPALISHIRING" is written in a zigzag pattern on a given number of rows like this: (you may want to display this pattern in a fixed font for better legibility)

P   A   H   N
A P L S I I G
Y   I   R

And then read line by line: "PAHNAPLSIIGYIR"

Write the code that will take a string and make this conversion given a number of rows:
"""


class Solution:
    def convert(self, s: str, numRows: int) -> str:
        if numRows == 1:
            return s

        arr = [[] for _ in range(numRows)]
        idx = 0
        minus = False
        for char in s:
            arr[idx].append(char)
            if idx == 0:
                minus = False
            elif idx == numRows - 1:
                minus = True
            idx += -1 if minus else 1
            idx = idx % numRows
        return "".join(["".join(inner) for inner in arr])


if __name__ == "__main__":
    sol = Solution()
    print(sol.convert("PAYPALISHIRING", 1))
