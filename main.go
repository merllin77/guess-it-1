package main

import (
	"bufio"
	"fmt"
	"guess-it/converters"
	"math"
	"os"
	"strconv"
	"strings"
)

func calculateRange(userInput []float64, numbersLen int) {
	average := converters.Average(userInput)
	stdDeviation := converters.StdDeviation(userInput, average)
	lower, upper := converters.PredictRange(average, stdDeviation, numbersLen)
	fmt.Println(int(math.Floor(lower)), int(math.Ceil(upper))) // roundation to lower decimal (math.Lower) and to upper decimal (math.Ceil)
}

func main() {
	var (
		userInput  []float64
		numbersLen int
	)
	scanner := bufio.NewScanner(os.Stdin) // using of os.Stdin for user input from CLI

	// Read everyline from user input
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		if line == "" { // Bad input
			fmt.Println("Wrong Input")
			return
		}

		num, err := strconv.ParseFloat(line, 64)
		if err != nil {
			fmt.Println("Error in number conversion")
			return
		}

		if numbersLen >= 1 {
			deviation := num - userInput[numbersLen-1]
			if (int(deviation) >= 1000 && int(deviation) >= -1000) || (int(deviation) <= 1000 && int(deviation) <= -1000) {
				// userInput = userInput // Do nothing (skip if a random with big range difference value found)
			} else {
				userInput = append(userInput, num) // save all the user input
			}
		}

		// Add the first num
		if numbersLen == 0 {
			userInput = append(userInput, num)
		}

		numbersLen = len(userInput) // length of userInput
		calculateRange(userInput, numbersLen)

		// // Calculation for the FIRST 10 numbers
		// if numbersLen <= 10 {
		// 	calculateRange(userInput, numbersLen)
		// }

		// // Calculation for the LAST 10 numbers
		// if numbersLen > 10 {
		// 	recentUserInput := userInput[len(userInput)-10:]
		// 	calculateRange(recentUserInput, 10)
		// }
	}
}
