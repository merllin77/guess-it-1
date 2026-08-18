package converters

func PredictRange(average float64, stdDeviation float64, numbersLen int) (float64, float64) {
	var lower, upper float64

	if numbersLen == 1 {
		lower = average - 50
		upper = average + 50
		return lower, upper
	}

	stdDeviation *= 1.58 // double the deviation for openning prediction (dynamic range)
	lower = average - stdDeviation
	upper = average + stdDeviation
	return lower, upper
}
