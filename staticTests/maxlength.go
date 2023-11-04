package staticTests

func CheckRunLengthOne(array []bool) bool {
	num1, num2, num3, num4, num5, num6 := 0, 0, 0, 0, 0, 0
	currentLength := 0
	inSeries := false

	for i := range array {
		if array[i] == true {
			currentLength++
			inSeries = true
		} else if inSeries {
			inSeries = false
			switch currentLength {
			case 1:
				num1++
			case 2:
				num2++
			case 3:
				num3++
			case 4:
				num4++
			case 5:
				num5++
			default:
				num6++
			}
			currentLength = 0
		}
	}

	return (2267 <= num1 && num1 <= 2733) &&
		(1079 <= num2 && num2 <= 1421) &&
		(502 <= num3 && num3 <= 748) &&
		(223 <= num4 && num4 <= 402) &&
		(90 <= num5 && num5 <= 223) &&
		(90 <= num6 && num6 <= 223)
}

func CheckRunLengthZero(array []bool) bool {
	num1, num2, num3, num4, num5, num6 := 0, 0, 0, 0, 0, 0
	currentLength := 0
	inSeries := false

	for i := range array {
		if array[i] == false {
			currentLength++
			inSeries = true
		} else if inSeries {
			inSeries = false
			switch currentLength {
			case 1:
				num1++
			case 2:
				num2++
			case 3:
				num3++
			case 4:
				num4++
			case 5:
				num5++
			default:
				num6++
			}
			currentLength = 0
		}
	}

	return (2267 <= num1 && num1 <= 2733) &&
		(1079 <= num2 && num2 <= 1421) &&
		(502 <= num3 && num3 <= 748) &&
		(223 <= num4 && num4 <= 402) &&
		(90 <= num5 && num5 <= 223) &&
		(90 <= num6 && num6 <= 223)
}
