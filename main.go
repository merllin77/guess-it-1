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

func calculateRange(userInput []float64, numbersLen int) (int, int) {
	average := converters.Average(userInput)
	stdDeviation := converters.StdDeviation(userInput, average)
	lower, upper := converters.PredictRange(average, stdDeviation, numbersLen)
	return int(math.Floor(lower)), int(math.Ceil(upper)) // rounding to lower decimal (math.Floor) and to upper decimal (math.Ceil)
}

func main() {
	var (
		userInput  []float64
		numbersLen int
	)
	scanner := bufio.NewScanner(os.Stdin) // "scanner" reads the user's input from CLI (using os.Stdin)

	// Read one input line at a time from user input
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		if line == "" { // Bad input
			fmt.Println("Wrong Input")
			return
		}

		num, err := strconv.ParseFloat(line, 64) // convert string to float64
		if err != nil {
			fmt.Println("Error in number conversion")
			return
		}
		numAccepted := true // initialize every new num as accepted
		if numbersLen >= 1 {
			deviation := num - userInput[numbersLen-1]
			if (int(deviation) >= 1000 && int(deviation) >= -1000) || (int(deviation) <= 1000 && int(deviation) <= -1000) {
				// userInput = userInput // Do nothing (reject random range variation of 1000 from current accepted value)
				numAccepted = false // if num is out of filtering range mark it as false
			}

			if numAccepted {
				userInput = append(userInput, num) // append only the accepted (numAccepted) user input to history
			}

		}

		// Add the first input to accepted history
		if numbersLen == 0 {
			userInput = append(userInput, num)
		}

		numbersLen = len(userInput) // length of userInput's history
		lower, upper := calculateRange(userInput, numbersLen)

		// If the number is accepted, ensure the final range includes the current number.
		if numAccepted {
			if int(num) <= lower {
				lower = int(num)
			} else if int(num) >= upper {
				upper = int(num)
			}
		}

		fmt.Println(lower, upper)
	}

	// Check if the scanner has completed successfully
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}
