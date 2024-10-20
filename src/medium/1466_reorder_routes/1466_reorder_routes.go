package main

func minReorder(_ int, connections [][]int) int {
	// Note: We didn't use n because n is always just len(connections) + 1.
	// Each entry in connections map stores a list of cities that it can go to along the direction of the road.
	connectionsMap := map[int][]int{}
	// The reverse connections map, on the other hand, stores the list of cities it can go to
	// if the direction of the road is to be reversed.
	reverseConnectionsMap := map[int][]int{}
	for _, connection := range connections {
		connectionsMap[connection[0]] = append(connectionsMap[connection[0]], connection[1])
		reverseConnectionsMap[connection[1]] = append(reverseConnectionsMap[connection[1]], connection[0])
	}

	reorders := 0
	citiesToVisitStack := []int{0}
	visitedCities := map[int]bool{}

	for len(citiesToVisitStack) > 0 {
		// Pop city from stack.
		city := citiesToVisitStack[len(citiesToVisitStack)-1]
		citiesToVisitStack = citiesToVisitStack[:len(citiesToVisitStack)-1]

		// Mark current city as visited.
		visitedCities[city] = true

		// We first check the reverse connections, these cities already lead to zero (since we started from zero).
		reverseConnections := reverseConnectionsMap[city]
		for _, nextCity := range reverseConnections {
			if !visitedCities[nextCity] {
				citiesToVisitStack = append(citiesToVisitStack, nextCity)
			}
		}

		// Next we check the non-reversed connections, these cities will need their road reversed to go to zero.
		connections := connectionsMap[city]
		for _, nextCity := range connections {
			if !visitedCities[nextCity] {
				reorders++
				citiesToVisitStack = append(citiesToVisitStack, nextCity)
			}
		}
	}

	return reorders
}
