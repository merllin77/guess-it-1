package converters

func Median(nums []float64) float64 {

	var median float64
	var middle int
	even := false
	selectedMiddleNums := []float64{}

	// Check if the sum of numbers are even or odd
	if len(nums)%2 == 0 {
		even = true
	} else {
		even = false
	}

	// Sorting []nums in order min > max

	for i := 0; i < len(nums)-1; i++ {
		for j := 0; j < len(nums)-i-1; j++ {
			if nums[j] > nums[j+1] { // swaping positions
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}

	// If the lentgh of numbers are EVEN, take the average of 2 middle numbers

	if even {
		middle = len(nums) / 2
		selectedMiddleNums = append(selectedMiddleNums, nums[middle-1], nums[middle])
		median = Average(selectedMiddleNums)
	}

	// If the lentgh of numbers are ODD, take just the middle number
	if !even {
		middle = len(nums) / 2
		median = nums[middle]
	}

	return median
}
