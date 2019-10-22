"""
Given n non-negative integers a1, a2, ..., an , where each represents a point at coordinate (i, ai). n vertical lines are drawn such that the two endpoints of line i is at (i, ai) and (i, 0). Find two lines, which together with x-axis forms a container, such that the container contains the most water.

Note: You may not slant the container and n is at least 2.
"""


class Solution:
    def maxArea(self, height: [int]) -> int:
        i = 0
        j = len(height) - 1
        max_area = 0
        while i != j:
            area = min(height[i], height[j]) * (j - i)
            max_area = max(max_area, area)
            if height[i] < height[j]:
                i += 1
            else:
                j -= 1
        return max_area



if __name__ == "__main__":
    sol = Solution()
    print(sol.maxArea([1,2,100,1,1,1,1,1,2]))
