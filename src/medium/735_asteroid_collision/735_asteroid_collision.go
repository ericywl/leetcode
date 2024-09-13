package main

func asteroidCollision(asteroids []int) []int {
	var stack []int
	for _, asteroid := range asteroids {
		for {
			// No more asteroids in stack, just add the new one in
			if len(stack) == 0 {
				stack = append(stack, asteroid)
				break
			}

			previousAsteroid := stack[len(stack)-1]
			// Asteroids won't collide
			if !willCollide(previousAsteroid, asteroid) {
				stack = append(stack, asteroid)
				break
			}

			// Asteroids are moving into each other, will collide
			// Pop previous asteroid
			stack = stack[:len(stack)-1]
			asteroidSize, previousAsteroidSize := abs(asteroid), abs(previousAsteroid)
			// Check which asteroid to keep
			//  - If current > previous, keep new asteroid and continue checking
			//  - If previous > current, add back previous asteroid and stop
			//  - Otherwise, both exploded, simply stop
			if asteroidSize < previousAsteroidSize {
				stack = append(stack, previousAsteroidSize)
				break
			} else if asteroidSize == previousAsteroidSize {
				break
			}

		}
	}

	return stack
}

func willCollide(a, b int) bool {
	return a > 0 && b < 0
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
