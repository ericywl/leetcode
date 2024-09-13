package main

import (
	"strconv"
	"strings"
	"unicode"
)

/*
Given an encoded string, return its decoded string.

The encoding rule is: k[encoded_string], where the encoded_string inside the square brackets is being repeated exactly k times.
Note that k is guaranteed to be a positive integer.

You may assume that the input string is always valid; there are no extra white spaces, square brackets are well-formed, etc.
Furthermore, you may assume that the original data does not contain any digits and that digits are only for those repeat numbers, k.
For example, there will not be input like 3a or 2[4].
*/

type RepeatString struct {
	Repeat int
	Chars  []rune
}

func (r RepeatString) String() string {
	return strings.Repeat(string(r.Chars), r.Repeat)
}

func decodeString(s string) string {
	var currentRepeat string
	stack := []RepeatString{{Chars: []rune{}}}

	for _, c := range s {
		// If it is a digit, add it to repeat
		if unicode.IsDigit(c) {
			currentRepeat += string(c)
			continue
		}

		// If it is a left square bracket '[' we add the repeat number
		if c == '[' {
			repeat, _ := strconv.Atoi(currentRepeat) // Assume no error
			stack = append(stack, RepeatString{Repeat: repeat})
			currentRepeat = ""
			continue
		}

		// If it is a right square bracket ']', we will build the new substring
		if c == ']' {
			// Pop last substring from stack
			repeatStr := stack[len(stack)-1].String()
			stack = stack[:len(stack)-1]
			// Add the repeated substring to the top element in stack
			stack[len(stack)-1].Chars = append(stack[len(stack)-1].Chars, []rune(repeatStr)...)
			continue
		}

		// Add new character to the top element in stack
		stack[len(stack)-1].Chars = append(stack[len(stack)-1].Chars, c)
	}

	return string(stack[0].Chars)
}
