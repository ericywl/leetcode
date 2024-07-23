package main

func longestCommonSubsequence(text1 string, text2 string) int {
	dp := make([][]int, len(text1))
	for i := 0; i < len(text1); i++ {
		dp[i] = make([]int, len(text2))
	}

	return lcs(text1, text2, dp, 0, 0)
}

func lcs(text1 string, text2 string, dp [][]int, i, j int) int {
	// Reached the end of one of the texts
	if len(text1) == i || len(text2) == j {
		return 0
	}

	// We have a memoized solution
	if dp[i][j] != 0 {
		return dp[i][j]
	}

	// Equal characters, advance both
	if text1[i] == text2[j] {
		dp[i][j] = 1 + lcs(text1, text2, dp, i+1, j+1)
		return dp[i][j]
	}

	// Advance one each and compare both to see which is better
	a := lcs(text1, text2, dp, i+1, j)
	b := lcs(text1, text2, dp, i, j+1)
	if a > b {
		dp[i][j] = a
	} else {
		dp[i][j] = b
	}
	return dp[i][j]
}
