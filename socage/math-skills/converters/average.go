package converters

func Average(nums []float64) float64 {

	var sum float64
	var average float64

	for _, n := range nums {
		sum += n
	}
	if sum == 0 { // if data is empty
		return 0
	}
	average = sum / float64(len(nums)) // Convert the len(nums) which return an int, to float64(len(nums) which return float64 for compatibility
	return average
}
