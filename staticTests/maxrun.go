package staticTests

func CheckMaxRun(sequence []bool) bool {
	maxZerosRun := 0
	maxOnesRun := 0
	currentZerosRun := 0
	currentOnesRun := 0

	for _, bit := range sequence {
		if bit {
			// Знайдено одиницю
			currentOnesRun++
			currentZerosRun = 0

			if currentOnesRun > maxOnesRun {
				maxOnesRun = currentOnesRun
			}
		} else {
			// Знайдено нуль
			currentZerosRun++
			currentOnesRun = 0

			if currentZerosRun > maxZerosRun {
				maxZerosRun = currentZerosRun
			}
		}
	}

	return maxZerosRun <= 36 && maxOnesRun <= 36
}
