package main

func maxVowels(s string, k int) int {
	maxNumVowels := 0
	currNumVowels := 0
	for i := 0; i < len(s); i++ {
		if isVowel(rune(s[i])) {
			currNumVowels += 1
		}
		if i >= k && isVowel(rune(s[i-k])) {
			currNumVowels -= 1
		}
		if currNumVowels > maxNumVowels {
			maxNumVowels = currNumVowels
		}
	}

	return maxNumVowels
}

func isVowel(character rune) bool {
	return character == 'a' || character == 'e' || character == 'i' || character == 'o' || character == 'u'
}
