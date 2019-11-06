"""
Implement pow(x, n), which calculates x raised to the power n (x^n).
"""


class Solution:
    def myPow(self, x: float, n: int) -> float:
        if n < 0:
            x, n = 1.0 / x, -n
        result = 1.0
        while n:
            # Generalizing, if the LSB of y is 0, the result is (x^(y / 2))^2
            # Otherwise it is x * (x^(y / 2))^2
            if n & 1:
                result *= x
            x, n = x * x, n >> 1
        return result
        

if __name__ == "__main__":
    sol = Solution()
    print(sol.myPow(2, 0))
