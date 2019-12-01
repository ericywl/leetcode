"""
There are N gas stations along a circular route, where the amount of gas at station i is gas[i].

You have a car with an unlimited gas tank and it costs cost[i] of gas to travel from station i to 
its next station (i+1). You begin the journey with an empty tank at one of the gas stations.

Return the starting gas station's index if you can travel around the circuit once in the clockwise 
direction, otherwise return -1.
"""


from typing import List

class Solution:
    def canCompleteCircuit(self, gas: List[int], cost: List[int]) -> int:
        """O(n) time and O(1) space"""
        if not gas or not cost or sum(gas) < sum(cost):
            return -1
            
        tank, start = 0, 0
        for i in range(len(gas)):
            tank += gas[i] - cost[i]
            if tank < 0:
                tank = 0
                start = (i + 1) % len(gas)
        return start
        

class Solution2:
    def canCompleteCircuit(self, gas: List[int], cost: List[int]) -> int:
        """O(n^2) time and O(1) space, time exceeds limit"""
        for i in range(len(gas)):
            if cost[i] > gas[i]:
                continue
            tank = gas[i] - cost[i]
            j = (i + 1) % len(gas)
            while tank >= 0:
                tank += gas[j] - cost[j]
                if j == i:
                    return i
                j = (j + 1) % len(gas)
        return -1
        



if __name__ == "__main__":
    sol = Solution()
    gas = [5, 1, 2, 3, 4]
    cost = [4, 4, 1, 5, 1]
    diff = [1, -3, 1, -2, 3]
    print(sol.canCompleteCircuit(gas, cost))
