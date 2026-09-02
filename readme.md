# Guess It 1

A Go command-line program that reads numbers continuously from standard input and prints a prediction range for the next number after every input.

## How it predicts

The program keeps accepted input values and calculates a range centred on their average.

- With one accepted value, it uses a safe range of 50 below to 50 above that value.
- Afterwards, the range is the average plus or minus 1.58 standard deviations.
- The bounds are rounded outward: down for the lower bound and up for the upper bound. This prevents an otherwise valid decimal range from becoming narrower when printed as integers.
- Statistics use all accepted values collected during the current run.
- A value whose distance from the latest accepted value is at least 1,000 is treated as an abrupt outlier and is not added to the statistical history. The program still prints a prediction after it.
- When the current value is accepted, the final printed interval is expanded if necessary to include that value.

This balances two goals from the exercise: covering the next value and keeping ranges narrow enough to score well.

## Input and output

Provide one number per line. For each received number, the program writes one line with two integers separated by a space: the lower and upper bounds for the next input.

Example:

```text
189
139 239
113
101 201
```

The exact predicted bounds change as more values arrive.

## Run locally

```sh
go run .
```

Then type numbers line by line. End the input with `Ctrl+D` on Linux/macOS.

To build an executable:

```sh
go build -o guess-it .
```

## Test

Run the Go checks with:

```sh
go test ./...
```

The supplied browser-based tester is in `guess-it-dockerized/`. It expects the compiled Linux executable and launcher at `guess-it-dockerized/student/`.

```sh
cd guess-it-dockerized
docker compose up --build
```

Open `http://localhost:3000` and compare the **Final Result** as well as coverage. A wider interval can improve correct guesses while lowering the score, so Final Result is the main measure for tuning.
