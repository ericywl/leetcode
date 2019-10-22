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