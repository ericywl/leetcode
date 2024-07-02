/**
 * Given an integer array nums, return true if there exists a triple of indices (i, j, k) such that i < j < k and nums[i] < nums[j] < nums[k]. 
 * If no such indices exists, return false.
 * 
 * Example 1:
 * Input: nums = [1,2,3,4,5]
 * Output: true
 * Explanation: Any triplet where i < j < k is valid.
 * 
 * Example 2:
 * Input: nums = [5,4,3,2,1]
 * Output: false
 * Explanation: No triplet exists.
 * 
 * Example 3:
 * Input: nums = [2,1,5,0,4,6]
 * Output: true
 * Explanation: The triplet (3, 4, 5) is valid because nums[3] == 0 < nums[4] == 4 < nums[5] == 6.
 */

const MAX = 2 ** 31 - 1;

/**
 * Algorithm walkthrough:
 * We keep track of the current best increasing tuple subsequence and the smallest number seen so far.
 * 
 * Given the input...
 *    nums:           [20,  100,  10,  12,  5, 13]
 * The state of all variables after each loop are...
 *    bestTuple[0]:   [20,   20,  10,  20, 10, 10]
 *    bestTuple[1]:   [MAX, 100, 100, 100, 12, 12]
 * 
 * Note: If bestTuple-1 is ever set, it means that there already was a case where bestTuple[0] < bestTuple[1].
 *       So it doesn't matter if we were to overwrite bestTuple[0] with something smaller later on.
 */
function increasingTriplet(nums: number[]): boolean {
    if (nums.length < 3) {
        return false;
    }

    let bestTuple: [number, number] = [MAX, MAX];
    for (let num of nums) {
        if (num <= bestTuple[0]) {
            bestTuple[0] = num;
        } else if (num <= bestTuple[1]) {
            bestTuple[1] = num;
        } else {
            return true;
        }
    }

    return false;
};

import { test, expect } from "bun:test";

test("true", () => {
    expect(increasingTriplet([2, 1, 5, 0, 4, 6])).toBe(true);
})

test("true 2", () => {
    expect(increasingTriplet([20, 100, 10, 12, 5, 13])).toBe(true);
})

test("true 3", () => {
    expect(increasingTriplet([1, 5, 0, 4, 1, 3])).toBe(true);
})

test("false", () => {
    expect(increasingTriplet([5, 4, 3, 2, 1])).toBe(false);
})