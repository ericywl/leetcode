"""
Given a collection of intervals, merge all overlapping intervals.
"""

from typing import List


class Solution:
    def merge(self, intervals: List[List[int]]) -> List[List[int]]:
        """O(n^3) time and O(n) space"""
        def checkOverlap(inter1: List[int], inter2: List[int]) -> bool:
            return (any(x in range(inter2[0], inter2[1] + 1) for x in inter1)
                    or any(x in range(inter1[0], inter1[1] + 1) for x in inter2))

        temp = intervals[:]
        i = 0
        while i < len(temp) - 1:
            for j in range(i + 1, len(temp)):
                if checkOverlap(temp[i], temp[j]):
                    start = min(temp[i], temp[j], key=lambda x: x[0])[0]
                    end = max(temp[i], temp[j], key=lambda x: x[1])[1]
                    temp[i] = [start, end]
                    del temp[j]
                    i -= 1
                    break
            i += 1
        return temp


class Solution2:
    def merge(self, intervals):
        """O(n log n) time and O(n) space"""
        # if we sort the intervals by their start value,
        # then each set of intervals that can be merged will
        # appear as a contiguous "run" in the sorted list.
        intervals.sort(key=lambda x: x.start)

        merged = []
        for interval in intervals:
            # if the list of merged intervals is empty or if the current
            # interval does not overlap with the previous, simply append it.
            if not merged or merged[-1].end < interval.start:
                merged.append(interval)
            else:
                # otherwise, there is overlap, so we merge the current and previous
                # intervals.
                merged[-1].end = max(merged[-1].end, interval.end)

        return merged


if __name__ == "__main__":
    sol = Solution()
    arr = [[1, 3], [2, 6], [8, 10], [15, 18]]
    arr2 = [[1, 4], [0, 2], [3, 5]]
    print(sol.merge(arr2))
