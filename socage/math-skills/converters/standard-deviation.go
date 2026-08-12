package converters

import "math"

func StandardDeviation(nums []float64, variance float64) float64 {

	stdDev := math.Sqrt(variance)
	return stdDev
}
