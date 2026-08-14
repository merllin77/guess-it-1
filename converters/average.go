package converters

func Average(numbers []float64) float64 {
	sum := 0.0

	for _, num := range numbers {
		sum += num
	}
	result := sum / float64(len(numbers))

	return result
}
