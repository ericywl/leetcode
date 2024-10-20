package main

const (
	Radiant = 'R'
	Dire    = 'D'
)

func predictPartyVictory(senate string) string {
	// Initialize both Radiant and Dire indices in separate arrays
	radiant := []int{}
	dire := []int{}
	for i, party := range senate {
		if party == Radiant {
			radiant = append(radiant, i)
		} else {
			dire = append(dire, i)
		}
	}

	for len(radiant) > 0 && len(dire) > 0 {
		// Find the party that comes out earlier, that party will ban the opposing one (done in dequeue below)
		// Note: We add len(senate) to ensure the indices are strictly increasing for comparison
		if radiant[0] < dire[0] {
			radiant = append(radiant, radiant[0]+len(senate))
		} else {
			dire = append(dire, dire[0]+len(senate))
		}

		// Dequeue both elements
		radiant = radiant[1:]
		dire = dire[1:]
	}

	if len(radiant) > 0 {
		return "Radiant"
	}
	return "Dire"
}
