package main

func exist(board [][]byte, word string) bool {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if existDFS(board, i, j, word, 0) {
				return true
			}
		}
	}

	return false
}

const visitedChar = '#'

func existDFS(board [][]byte, row int, col int, word string, index int) bool {
	if index >= len(word) {
		return true
	}

	height := len(board)
	width := len(board[0])

	// Invalid cases
	if row < 0 || row >= height || col < 0 || col >= width {
		return false
	}

	// Character has been visited already
	char := board[row][col]
	if char == visitedChar {
		return false
	}

	// Character is not the next in the word
	if char != word[index] {
		return false
	}

	// Mark current cell as visited to pass into the DFS
	board[row][col] = visitedChar

	// Visit other directions
	dy := []int{1, -1, 0, 0}
	dx := []int{0, 0, -1, 1}
	for i := range dy {
		if existDFS(board, row+dy[i], col+dx[i], word, index+1) {
			return true
		}
	}

	// Reset the board character
	board[row][col] = char
	return false
}
