package main

import (
	"bufio"
	"fmt"
	"math"
	"math-skills/converters"
	"os"
	"strconv"
)

// String to number conversion with strconv.Atoi
func strToNumber(s string) float64 {

	num, err := strconv.ParseFloat(s, 64)
	if err != nil {
		fmt.Println("Error in string to integer conversion: ", err)
	}
	return num
}

func main() {
	var averageHold float64
	var medianHold float64
	var varianceHold float64

	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Println("Please, check your arguments. You gave more or less!")
		return
	}
	fileName := args[0]

	// Opening file for reading
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	allInts := []float64{}
	scanner := bufio.NewScanner(file)

	// Scan data.txt file line by line for creating a new []float64 with numbers (allInts := []float64{})
	for scanner.Scan() {
		line := scanner.Text()
		allInts = append(allInts, strToNumber(line)) // create []int of all strings to integers conversion
	}
	// Compute Average
	averageHold = converters.Average(allInts)
	roundedAverage := int(math.Round(averageHold)) // rounded integer from float64 (or if num float64 >= 0 --> + 0.5 manually else -0.5 (for negative values)
	fmt.Println("Average: ", roundedAverage)

	// // Compute Median
	medianHold = converters.Median(allInts)
	roundedMedian := int(math.Round(medianHold))
	fmt.Println("Median: ", roundedMedian)

	// Compute Variance
	varianceHold = converters.Variance(allInts, averageHold)
	roundedVariance := int(math.Round(varianceHold))
	fmt.Println("Variance: ", roundedVariance)

	// Compute Standard Deviation
	stdDeviation := converters.StandardDeviation(allInts, varianceHold)
	roundedStdDeviation := math.Round(stdDeviation)
	fmt.Println("Standard Deviation: ", roundedStdDeviation)
}
