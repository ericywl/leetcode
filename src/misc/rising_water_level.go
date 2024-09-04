/*
You are given a grid of land heights. For example:

6 5 4 5 5
4 2 5 1 1
5 5 2 1 5
2 3 2 4 4
5 4 5 5 6

The upper left corner is the start and the bottom riht corner is the end. You may assume that
the start and end cells are of equal height and are the highest points in the grid.

Initially the water level is 0. Due to global warming, the water level increases by 1 every
year. A cell if flooded if the water level is greater than or equal to its height.

How many years do we have before there is no longer a path from the start to the end cell?
Directions of movement: Up, Down, Left, Right.
*/

package main

func yearsLeft(grid [][]int) int {
	width := len(grid[0])
	height := len(grid)

	// Min time we have is if there is already no path to the end
	low := 0
	// Max time we have is if there is a path whose cells all have same height as the start
	high := grid[0][0]

	for low < high {
		mid := low + (high-low)/2
		visited := initGrid[bool](height, width)
		if endReachableDFS(grid, visited, mid, 0, 0) {
			// Still reachable, increase water height
			low = mid + 1
		} else {
			// No longer reachable, decrease water height
			high = mid
		}
	}

	return high
}

func endReachableDFS(grid [][]int, visited [][]bool, waterHeight int, row int, col int) bool {
	width := len(grid[0])
	height := len(grid)

	// Invalid cases
	if row < 0 || row >= height || col < 0 || col >= width {
		return false
	}

	// Visited before, or current cell already submerged
	if visited[row][col] || grid[row][col] <= waterHeight {
		return false
	}

	// Mark cell as visited
	visited[row][col] = true

	// We have reached the end cell
	if row == height-1 && col == width-1 {
		return true
	}

	// Check other directions if we can reach the end through them
	// Movement directions: Up, Down, Left, Right
	dy := []int{1, -1, 0, 0}
	dx := []int{0, 0, -1, 1}
	for i := range dy {
		if endReachableDFS(grid, visited, waterHeight, row+dy[i], col+dx[i]) {
			return true
		}
	}
	return false
}

func initGrid[T any](rows, cols int) [][]T {
	grid := make([][]T, rows)
	for i := range grid {
		grid[i] = make([]T, cols)
	}
	return grid
}
