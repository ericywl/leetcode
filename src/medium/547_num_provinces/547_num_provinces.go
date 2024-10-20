package main

func findCircleNum(isConnected [][]int) int {
	citiesVisited := map[int]bool{}
	numProvinces := 0
	// Loop through all cities as they can all potentially be its own province
	for city := 0; city < len(isConnected); city++ {
		// If city has been visited before, it means that they were already part
		// of another province so we can ignore it
		if citiesVisited[city] {
			continue
		}

		// Perform DFS with the city as root
		citiesToVisitStack := []int{city}
		for len(citiesToVisitStack) > 0 {
			// Pop city from stack
			cityToVisit := citiesToVisitStack[len(citiesToVisitStack)-1]
			citiesToVisitStack = citiesToVisitStack[:len(citiesToVisitStack)-1]

			// City already been visited before, ignore it
			if citiesVisited[cityToVisit] {
				continue
			}

			// Visit the city and add new connections to be visited at the top of stack
			citiesVisited[cityToVisit] = true
			cityConnections := isConnected[cityToVisit]
			for city, isConnected := range cityConnections {
				if isConnected > 0 && !citiesVisited[city] {
					citiesToVisitStack = append(citiesToVisitStack, city)
				}
			}
		}

		// Increment number of provinces once we are done with DFS
		numProvinces++
	}

	return numProvinces
}
