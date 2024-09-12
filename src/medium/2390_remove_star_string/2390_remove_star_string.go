package main

func removeStars(s string) string {
	var stack []rune
	for _, c := range s {
		if c == '*' {
			// Remove character from stack
			stack = stack[:len(stack)-1]
		} else {
			// Add character to stack
			stack = append(stack, c)
		}
	}

	return string(stack)
}
