/**
 * Given an integer array nums, move all 0's to the end of it while maintaining the relative order of the non-zero elements.
 * Note that you must do this in-place without making a copy of the array.
 * 
 * Example 1:
 * Input: nums = [0,1,0,3,12]
 * Output: [1,3,12,0,0]
 * 
 * Example 2:
 * Input: nums = [0]
 * Output: [0]
 */

/**
 * Algorithm walkthrough:
 * We first initialize 2 pointers, i and j to 0 and 1. We will keep looping until j exceeds the end.
 * Note that j should always be ahead of i.
 * 
 * There are 4 scenarios that can happen:
 *     1. nums[i] == 0 and nums[j] == 0: 
 *          We should leave i where it is so that we can swap future non-zeroes to its position, increment j
 *     2. nums[i] != 0 and nums[j] != 0: 
 *          We don't want to swap since we need to preserve the original order, increment both
 *     3. nums[i] != 0 and nums[j] == 0: 
 *          We don't want to swap since that would bring zeroes to the front, increment both
 *     4. nums[i] == 0 and nums[j] != 0: 
 *          We swap numbers on i and j, to bring non-zeroes to the front, increment both
 */
function moveZeroes(nums: number[]): void {
    let i: number = 0;
    let j: number = 1;

    while (j < nums.length) {
        if (nums[i] === 0 && nums[j] === 0) {
            j++;
            continue;
        }

        if (nums[i] === 0 && nums[j] !== 0) {
            nums[i] = nums[j];
            nums[j] = 0;
        }

        i++;
        j++;
    }
}

import { test, expect } from "bun:test";

test("default", () => {
    let nums = [0, 1, 0, 3, 12];
    moveZeroes(nums);
    expect(nums).toStrictEqual([1, 3, 12, 0, 0]);
})

test("zero", () => {
    let nums = [0];
    moveZeroes(nums);
    expect(nums).toStrictEqual([0]);
})
