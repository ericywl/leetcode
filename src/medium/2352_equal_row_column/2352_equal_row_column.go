package main

func equalPairs(grid [][]int) int {
	columns := transpose(grid)
	equalCount := 0
	for _, row := range grid {
		for _, col := range columns {
			if sliceEqual(row, col) {
				equalCount++
			}
		}
	}

	return equalCount
}

func transpose(slice [][]int) [][]int {
	result := make([][]int, len(slice[0]))
	for i := range result {
		result[i] = make([]int, len(slice))
	}

	for i := 0; i < len(slice[0]); i++ {
		for j := 0; j < len(slice); j++ {
			result[i][j] = slice[j][i]
		}
	}

	return result
}

func sliceEqual(s1, s2 []int) bool {
	if len(s1) != len(s2) {
		return false
	}

	for i := range s1 {
		if s1[i] != s2[i] {
			return false
		}
	}

	return true
}
