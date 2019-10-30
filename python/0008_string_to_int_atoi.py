"""
Implement atoi which converts a string to an integer.

The function first discards as many whitespace characters as necessary until the first non-whitespace character is found. Then, starting from this character, takes an optional initial plus or minus sign followed by as many numerical digits as possible, and interprets them as a numerical value.

The string can contain additional characters after those that form the integral number, which are ignored and have no effect on the behavior of this function.

If the first sequence of non-whitespace characters in str is not a valid integral number, or if no such sequence exists because either str is empty or it contains only whitespace characters, no conversion is performed.

If no valid conversion could be performed, a zero value is returned.
"""


class Solution:
    INT_MIN = -(2 ** 31)
    INT_MAX = (2 ** 31) - 1

    def myAtoi(self, s: str) -> int:
        s = s.strip()
        # Check empty string
        if not s:
            return 0
        # Check starting character matches [0123456789+-]
        if not s[0].isdigit() and s[0] not in ["-", "+"]:
            return 0
        # Check +- sign
        multi = 1
        if s[0] in ["-", "+"]:
            multi = -1 if s[0] == "-" else 1
            s = s[1:]
        # Check for digits in remaining string 
        new_s = ""
        for char in s:
            if not char.isdigit():
                break
            new_s += char
        if not new_s:
            return 0
        res = multi * int(new_s)
        # Check for overflow
        if res < self.INT_MIN:
            return self.INT_MIN
        if res > self.INT_MAX:
            return self.INT_MAX
        return res


if __name__ == "__main__":
    sol = Solution()
    print(sol.myAtoi("-91283472332"))
