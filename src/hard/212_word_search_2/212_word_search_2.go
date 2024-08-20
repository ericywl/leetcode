package hard

/*
Given an m x n board of characters and a list of strings words, return all words on the board.

Each word must be constructed from letters of sequentially adjacent cells, where adjacent cells
are horizontally or vertically neighboring.
The same letter cell may not be used more than once in a word.

Assume only lowercase English letters.
*/

func findWords(board [][]byte, words []string) []string {
	result := make([]string, 0)
	root := buildWordsTrie(words)
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			findWordsDFS(board, i, j, root, &result)
		}
	}

	return result
}

const visitedChar = '#'

func findWordsDFS(board [][]byte, row int, col int, node *TrieNode, result *[]string) {
	height := len(board)
	width := len(board[0])

	// Invalid cases
	if row < 0 || row >= height || col < 0 || col >= width {
		return
	}

	char := board[row][col]
	// Visited before, or current character does not exist in word trie
	if char == visitedChar || node.Next[int(char-'a')] == nil {
		return
	}

	nextNode := node.Next[int(char-'a')]
	// The new node is a leaf (which has a word)
	if nextNode.Word != "" {
		*result = append(*result, nextNode.Word)
		// Remove the word from the leaf node to deduplicate it
		nextNode.Word = ""
	}

	// Mark character as visited to pass into the next DFS
	board[row][col] = visitedChar

	// Check other directions
	dy := []int{1, -1, 0, 0}
	dx := []int{0, 0, -1, 1}
	for i := range dy {
		findWordsDFS(board, row+dy[i], col+dx[i], nextNode, result)
	}

	// Restore the character
	board[row][col] = char
}

type TrieNode struct {
	Next []*TrieNode
	Word string
}

func newTrieNode() *TrieNode {
	return &TrieNode{
		Next: make([]*TrieNode, 26), // 26 lowercase letters
		Word: "",
	}
}

func buildWordsTrie(words []string) *TrieNode {
	root := newTrieNode()

	for _, word := range words {
		tmp := root
		for _, char := range word {
			index := int(char - 'a')
			if tmp.Next[index] == nil {
				tmp.Next[index] = newTrieNode()
			}

			tmp = tmp.Next[index]
		}

		// Store word at the leaf
		tmp.Word = word
	}

	return root
}
