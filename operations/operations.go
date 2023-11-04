package operations

import "math/rand"

func GenerateRandomBitSequence(size int) []bool {
	bitSequence := make([]bool, size)

	for i := 0; i < size; i++ {
		// Генеруємо випадкове число 0 або 1
		randomBit := rand.Intn(2) == 1
		bitSequence[i] = randomBit
	}

	return bitSequence
}

func TakeStringValue(n, m int, array []bool) string {
	value := ""

	for ; n < m; n++ {
		if array[n] {
			value += "1"
		} else {
			value += "0"
		}
	}
	return value
}
