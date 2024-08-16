package main

func canVisitAllRooms(rooms [][]int) bool {
	visited := map[int]bool{}
	toVisit := []int{0}

	for len(toVisit) > 0 {
		numToVisit := len(toVisit)
		for i := 0; i < numToVisit; i++ {
			// Pop from toVisit
			keyToRoom := toVisit[0]
			toVisit = toVisit[1:]

			// Visit room that has not been visited before
			if !visited[keyToRoom] {
				visited[keyToRoom] = true
				if len(visited) == len(rooms) {
					return true
				}
				// Add all keys / rooms into toVisit
				toVisit = append(toVisit, rooms[keyToRoom]...)
			}
		}
	}

	return len(visited) == len(rooms)
}
