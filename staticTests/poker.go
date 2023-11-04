package staticTests

import (
	"awesomeProject/operations"
	"math"
)

func CheckPoker(array []bool) bool {
	const (
		pokerLowerBound = 1.03
		pokerUpperBound = 57.4
	)

	var (
		m = 4
		k = len(array) / m
		n = make(map[string]int)
	)

	for i := 0; i < k; i++ {
		block := operations.TakeStringValue(i*m, (i+1)*m, array)
		n[block]++
	}

	sum := 0.0
	for _, count := range n {
		sum += float64(count * count)
	}

	X3 := math.Pow(2, float64(m))/float64(k)*sum - float64(k)

	return X3 > pokerLowerBound && X3 < 57.4
}
