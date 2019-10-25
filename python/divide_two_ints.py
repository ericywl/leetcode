"""
Given two integers dividend and divisor, divide two integers without using multiplication, division and mod operator.

Return the quotient after dividing dividend by divisor.

The integer division should truncate toward zero.
"""

class Solution:
    INT_MIN = -(2 ** 31)
    INT_MAX = 2 ** 31 - 1

    def divide(self, dividend: int, divisor: int) -> int:
        def get_bit(value: int, idx: int) -> int:
            return 1 if (value & (1 << idx)) else 0
            
        def set_bit(value: int, idx: int, new_bit: int) -> int:
            return value | (new_bit << idx)

        negative = (dividend * divisor) < 0
        divisor = abs(divisor)
        dividend = abs(dividend)

        quotient = 0
        remainder = 0
        for i in range(dividend.bit_length() - 1, -1, -1):
            remainder <<= 1
            bit_i = get_bit(dividend, i)
            remainder = set_bit(remainder, 0, bit_i) 
            if remainder >= divisor:
                remainder -= divisor
                quotient = set_bit(quotient, i, 1) 
        quotient = -quotient if negative else quotient
        return max(min(quotient, self.INT_MAX), self.INT_MIN)
        

if __name__ == "__main__":
    sol = Solution()
    res = sol.divide(-7, 3)
    print(res)
