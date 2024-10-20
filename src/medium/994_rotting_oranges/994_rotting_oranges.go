package main

type Coordinates struct {
	Row, Column int
}

const (
	EmptyCell    = 0
	FreshOrange  = 1
	RottenOrange = 2
)

func orangesRotting(grid [][]int) int {
	numFreshOranges := 0
	rotten := []Coordinates{}
	for i := range grid {
		for j := range grid[i] {
			switch grid[i][j] {
			case RottenOrange:
				rotten = append(rotten, Coordinates{Row: i, Column: j})
			case FreshOrange:
				numFreshOranges++
			}
		}
	}

	// Edge case where there are no fresh oranges
	if numFreshOranges == 0 {
		return 0
	}

	minutesPassed := 0
	// Note: We can optimize this away by modifying the matrix directly to keep track of
	// visited cells, but that would ruin the original matrix
	visited := map[Coordinates]bool{}

	for len(rotten) > 0 {
		// Visit only the ones that were originally in the queue
		num := len(rotten)
		for x := 0; x < num; x++ {
			// Dequeue rotting orange coordinates
			currentCoord := rotten[0]
			rotten = rotten[1:]

			// Enqueue next rotten oranges
			dy := []int{1, -1, 0, 0}
			dx := []int{0, 0, -1, 1}
			for i := range dy {
				nextCoord := Coordinates{Row: currentCoord.Row + dy[i], Column: currentCoord.Column + dx[i]}
				// We only enqueue if we have not already visited and the coordinates are valid and it contains fresh orange
				if !visited[nextCoord] && isValidCoordinates(nextCoord, grid) && grid[nextCoord.Row][nextCoord.Column] == FreshOrange {
					rotten = append(rotten, nextCoord)
					// Mark as visited
					visited[nextCoord] = true
					numFreshOranges--
				}
			}
		}

		// Increment minutes after each round
		minutesPassed++
		// If we have no more fresh oranges, return answer
		if numFreshOranges == 0 {
			return minutesPassed
		}
	}

	return -1
}

func isValidCoordinates(coord Coordinates, grid [][]int) bool {
	gridHeight := len(grid)
	gridWidth := len(grid[0])

	return coord.Row >= 0 && coord.Row < gridHeight && coord.Column >= 0 && coord.Column < gridWidth
}
