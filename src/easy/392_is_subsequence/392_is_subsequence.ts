function isSubsequence(s: string, t: string): boolean {
    let idx: number = 0;
    for (let char of t) {
        if (s[idx] === char) {
            idx++;
        }
    }

    return idx === s.length;
};

import { expect, test } from "bun:test";

test("true", () => {
    expect(isSubsequence("abc", "ahbgdc")).toBe(true);
})

test("false", () => {
    expect(isSubsequence("axc", "ahbgdc")).toBe(false);
})