# math-skills

A small Go CLI that reads a file containing one number per line and computes:

- Average
- Median
- Variance 
- Standard Deviation

Results are printed as rounded integers.

## Requirements

- Go 1.24+

## Project Structure

- `math-skills.go`: entry point, file reading, pipeline orchestration, final printing
- `converters/average.go`: average calculation
- `converters/median.go`: median calculation
- `converters/variance.go`: variance calculation
- `converters/standard-deviation.go`: standard deviation calculation
- `data.txt`: sample input file

## Input Format

The input file must contain one numeric value per line:

- Example

```txt
189
113
121
114
145
110
```

## Run

```bash
go run math-skills.go data.txt
```

## Expected Output Format

```txt
Average: 132
Median: 118
Variance: 785
Standard Deviation: 28
```

Notes:
- The exact numbers depend on your input file.
- Values are rounded for output.

## Math Formulas Used

For values `x1, x2, ..., xn`:

- Average (mean):
  - `mean = (x1 + x2 + ... + xn) / n`
- Median:
  - Sort values, then pick the middle value (or average of 2 middle values if `n` is even)
- Variance:
  - `variance = Σ(xi - mean)^2 / n`
- Standard deviation:
  - `stddev = sqrt(variance)`

## Development Notes

Current implementation follows a simple pipeline flow:

1. Read file
2. Convert each line to `float64`
3. Compute statistics in `converters`
4. Round and print results `(convert float64 to int for rounding)`
