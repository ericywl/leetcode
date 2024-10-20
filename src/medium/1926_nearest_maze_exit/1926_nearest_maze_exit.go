package main

type Coordinates struct {
	Row, Column int
}

type CoordinatesWithSteps struct {
	Coord Coordinates
	Steps int
}

func nearestExit(maze [][]byte, entrance []int) int {
	entranceCoord := Coordinates{
		Row:    entrance[0],
		Column: entrance[1],
	}
	visited := map[Coordinates]bool{entranceCoord: true}
	queue := []CoordinatesWithSteps{{Coord: entranceCoord, Steps: 0}}

	for len(queue) > 0 {
		// Dequeue coordinates from queue
		current := queue[0]
		queue = queue[1:]

		// Check cardinal directions
		dy := []int{1, -1, 0, 0}
		dx := []int{0, 0, -1, 1}
		for i := range dy {
			next := CoordinatesWithSteps{
				Coord: Coordinates{
					Row:    current.Coord.Row + dy[i],
					Column: current.Coord.Column + dx[i],
				},
				Steps: current.Steps + 1,
			}

			if visited[next.Coord] {
				continue
			}

			// Invalid coordinates or it is a wall, ignore it
			if !canMoveIntoCoord(next.Coord, maze) {
				continue
			}

			// The next step goes to an exit, and since we use BFS it will be the shortest distance from entrance
			if isExit(next.Coord, maze) {
				return next.Steps
			}

			// Enqueue and mark as visited since we already checked it
			queue = append(queue, next)
			visited[next.Coord] = true
		}
	}

	return -1
}

func canMoveIntoCoord(coord Coordinates, maze [][]byte) bool {
	mazeHeight := len(maze)
	mazeWidth := len(maze[0])

	if coord.Row < 0 || coord.Row >= mazeHeight {
		return false
	}

	if coord.Column < 0 || coord.Column >= mazeWidth {
		return false
	}

	return maze[coord.Row][coord.Column] == '.'
}

func isExit(coord Coordinates, maze [][]byte) bool {
	mazeHeight := len(maze)
	mazeWidth := len(maze[0])
	return coord.Row == 0 || coord.Row == mazeHeight-1 || coord.Column == 0 || coord.Column == mazeWidth-1
}
