/**
 * You are given an integer array height of length n. There are n vertical lines drawn such that the two endpoints of the ith line are (i, 0) and (i, height[i]).
 * Find two lines that together with the x-axis form a container, such that the container contains the most water.
 * 
 * Return the maximum amount of water a container can store.
 * 
 * Notice that you may not slant the container.
 */

import { textSpanContainsPosition } from "typescript";

function maxArea(heights: number[]): number {
    let left: number = 0;
    let right: number = heights.length - 1;
    let area: number = 0;

    while (left < right) {
        const newArea = (right - left) * Math.min(heights[left], heights[right]);
        if (newArea > area) {
            area = newArea;
        }

        if (heights[left] < heights[right]) {
            // The right one is taller, hence it will yield the most possible area, move the left instead
            left++;
        } else {
            // The left one is taller, hence it will yield the most possible area, move the right instead
            right--;
        }
    }

    return area;
};

import { test, expect } from "bun:test";

test("example 1", () => {
    expect(maxArea([1, 8, 6, 2, 5, 4, 8, 3, 7])).toBe(49);
})

test("example 2", () => {
    expect(maxArea([1, 1])).toBe(1);
})
