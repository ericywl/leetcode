package main

const (
	undefinedAnswer = -1.0
)

func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	// Store the divisions, inverse divisions and self divisions in a map for quick access later on.
	// First key is the dividend variable, second is the divisor variable, and finally the value is the quotient.
	divisionMap := map[string]map[string]float64{}
	for i, equation := range equations {
		if _, ok := divisionMap[equation[0]]; !ok {
			divisionMap[equation[0]] = map[string]float64{}
		}

		if _, ok := divisionMap[equation[1]]; !ok {
			divisionMap[equation[1]] = map[string]float64{}
		}

		// Division by each other
		divisionMap[equation[0]][equation[1]] = values[i]
		divisionMap[equation[1]][equation[0]] = 1.0 / values[i]
		// Division by the same variable
		divisionMap[equation[0]][equation[0]] = 1.0
		divisionMap[equation[1]][equation[1]] = 1.0
	}

	answers := make([]float64, 0, len(queries))
	for _, query := range queries {
		answers = append(answers, resolveQuery(query[0], query[1], divisionMap))
	}

	return answers
}

type VariableWithIntermediateAnswer struct {
	Variable           string
	IntermediateAnswer float64
}

func resolveQuery(first, second string, divisionMap map[string]map[string]float64) float64 {
	checkedVariables := map[string]bool{}
	// Initialize stack with first variable and intermediate answer of 1.0
	toCheckStack := []VariableWithIntermediateAnswer{
		{
			Variable:           first,
			IntermediateAnswer: 1.0,
		},
	}

	for len(toCheckStack) > 0 {
		// Pop variable from stack
		dividendVar := toCheckStack[len(toCheckStack)-1]
		toCheckStack = toCheckStack[:len(toCheckStack)-1]

		// Mark current variable as checked
		checkedVariables[dividendVar.Variable] = true

		// Get the divisors of the dividend and check if the variable even exists
		divisorsToQuotient, exist := divisionMap[dividendVar.Variable]
		if !exist {
			return undefinedAnswer
		}

		// Go through all the divisors of the current dividend variable
		for divisor, quotientValue := range divisorsToQuotient {
			if divisor == second {
				return dividendVar.IntermediateAnswer * quotientValue
			}

			// We have not seen this before, add it to stack
			if !checkedVariables[divisor] {
				toCheckStack = append(toCheckStack, VariableWithIntermediateAnswer{
					Variable:           divisor,
					IntermediateAnswer: dividendVar.IntermediateAnswer * quotientValue,
				})
			}
		}
	}

	return undefinedAnswer
}
