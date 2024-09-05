package main

import "sort"

/*
Two strings are considered close if you can attain one from the other using the following operations:

Operation 1: Swap any two existing characters.
For example, abcde -> aecdb

Operation 2: Transform every occurrence of one existing character into another existing character, and do
the same with the other character.
For example, aacabb -> bbcbaa (all a's turn into b's, and all b's turn into a's)

You can use the operations on either string as many times as necessary.

Given two strings, word1 and word2, return true if word1 and word2 are close, and false otherwise.
*/

func closeStrings(word1 string, word2 string) bool {
	if len(word1) != len(word2) {
		return false
	}

	charOcc1 := map[rune]int{}
	for _, c := range word1 {
		charOcc1[c]++
	}

	charOcc2 := map[rune]int{}
	for _, c := range word2 {
		charOcc2[c]++
	}

	return compareMapKeysEqual(charOcc1, charOcc2) && compareSortedMapValues(charOcc1, charOcc2)
}

func compareMapKeysEqual(m1, m2 map[rune]int) bool {
	for k := range m1 {
		if _, ok := m2[k]; !ok {
			return false
		}
	}

	for k := range m2 {
		if _, ok := m1[k]; !ok {
			return false
		}
	}

	return true
}

func compareSortedMapValues(m1, m2 map[rune]int) bool {
	var values1 []int
	for _, v := range m1 {
		values1 = append(values1, v)
	}

	var values2 []int
	for _, v := range m2 {
		values2 = append(values2, v)
	}

	sort.Ints(values1)
	sort.Ints(values2)

	for i := range values1 {
		if values1[i] != values2[i] {
			return false
		}
	}

	return true
}
