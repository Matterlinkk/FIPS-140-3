package staticTests

func CheckMonobit(array []bool) bool {
	const minimal, maximum = 9654, 10346
	count := 0

	for i := range array {
		if array[i] {
			count++
		}
	}

	if count >= minimal && count <= maximum {
		return true
	}
	return false
}
