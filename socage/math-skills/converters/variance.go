package converters

func Variance(nums []float64, mean float64) float64 {

	var deviation float64
	var squareDeviation float64
	var variance float64
	var sum float64

	// Return 0 if Data sheet is empty
	if len(nums) == 0 {
		return 0
	}

	for i := 0; i < len(nums); i++ {
		deviation = nums[i] - mean // mean = average
		squareDeviation = deviation * deviation
		sum += squareDeviation // sum of all squaredDeviations
	}
	variance = sum / float64(len(nums))
	return variance
}
