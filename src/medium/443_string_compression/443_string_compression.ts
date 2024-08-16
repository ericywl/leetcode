/**
 * Given an array of characters chars, compress it using the following algorithm:
 * 
 * Begin with an empty string s. For each group of consecutive repeating characters in chars:
 *  - If the group's length is 1, append the character to s
 *  - Otherwise, append the character followed by the group's length.
 * The compressed string s should not be returned separately, but instead, be stored in the input character array chars. 
 * Note that group lengths that are 10 or longer will be split into multiple characters in chars.
 * After you are done modifying the input array, return the new length of the array.
 * 
 * You must write an algorithm that uses only constant extra space.
 */

import { expect, test } from "bun:test";

function compress(chars: string[]): number {
    let readIdx: number = 0;
    let writeIdx: number = 0;

    while (readIdx < chars.length) {
        let char = chars[readIdx];
        let count = 0;

        // Count repeating characters
        while (readIdx < chars.length && chars[readIdx] === char) {
            readIdx++;
            count++;
        }

        // Write character
        chars[writeIdx++] = char;

        // Write count
        if (count > 1) {
            const countStr = count.toString();
            for (let digit of countStr) {
                chars[writeIdx] = digit;
                writeIdx++;
            }
        }
    }

    return writeIdx;
};

test("aabbccc", () => {
    let list = ["a", "a", "b", "b", "c", "c", "c"];
    const output = compress(list);
    expect(output).toBe(6);
    expect(list.slice(0, output)).toStrictEqual(["a", "2", "b", "2", "c", "3"]);
})

test("a", () => {
    let list = ["a"];
    const output = compress(list);
    expect(output).toBe(1);
    expect(list.slice(0, output)).toStrictEqual(["a"]);
})

test("abbbbbbbbbbbb", () => {
    let list = ["a", "b", "b", "b", "b", "b", "b", "b", "b", "b", "b", "b", "b"];
    const output = compress(list);
    expect(output).toBe(4);
    expect(list.slice(0, output)).toStrictEqual(["a", "b", "1", "2"]);
})
