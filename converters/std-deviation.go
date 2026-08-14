package converters

import "math"

func StdDeviation(numbers []float64, average float64) float64 {

	var sum, variation, squareVariation float64

	// Variance calculation
	for _, num := range numbers {
		variation = num - average
		squareVariation = variation * variation
		sum += squareVariation
	}

	// Standard Deviation calculation
	result := math.Sqrt(sum / float64(len(numbers)))
	return result
}
