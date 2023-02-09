[![Go Reference](https://pkg.go.dev/badge/github.com/MicahParks/go-ad.svg)](https://pkg.go.dev/github.com/MicahParks/go-ad) [![Go Report Card](https://goreportcard.com/badge/github.com/MicahParks/go-ad)](https://goreportcard.com/report/github.com/MicahParks/go-ad)
# go-ad
The Accumulation Distribution Line technical analysis algorithm implemented in Golang.

```go
import "github.com/MicahParks/go-ad"
```

# Usage
For full examples, please see the `examples` directory.

## Step 1
Gather the input for the current period.
```go
input := ad.Input{
	Close:  closePrices[i],
	Low:    low[i],
	High:   high[i],
	Volume: volume[i],
}
```

## Step 2
Create the A/D data structure and get the first result
```go
adLine, result := ad.New(input)
```

## Step 3
Use the next period's data to calculate the next point on the A/D line. Repeat for all periods.
```go
result = adLine.Calculate(input)
```

# Somewhat complete example (without data)
```go
package main

import (
	"log"
	"os"

	"github.com/MicahParks/go-ad"
)

func main() {
	// Create a logger.
	logger := log.New(os.Stdout, "", 0)

	// Iterate through the rest of the periods' data and calculate the A/D line's point for the given period.
	var adLine *ad.AD
	var result float64
	for i := range open {
		input := ad.Input{
			Close:  closePrices[i],
			Low:    low[i],
			High:   high[i],
			Volume: volume[i],
		}

		if adLine == nil {
			adLine, result = ad.New(input)
		} else {
			result = adLine.Calculate(input)
		}
		logger.Printf("Index: %d AD: %.4f", i, result)
	}
}
```

# Testing
There is 100% test coverage and benchmarks for this project. Here is an example benchmark result:
```
$ go test -bench .
goos: linux
goarch: amd64
pkg: github.com/MicahParks/go-ad
cpu: Intel(R) Core(TM) i5-9600K CPU @ 3.70GHz
BenchmarkAD_Calculate-6         1000000000               0.0000010 ns/op
BenchmarkBigAD_Calculate-6      1000000000               0.0000765 ns/op
PASS
ok      github.com/MicahParks/go-ad     0.004s
```

# Other Technical Algorithms
Looking for some other technical analysis algorithms? Here are some other ones I've implemented:
* Accumulation/Distribution (A/D): [go-ad](https://github.com/MicahParks/go-ad)
* Chaikin: [go-chaikin](https://github.com/MicahParks/go-chaikin)
* Moving Average Convergence Divergence (MACD), Exponential Moving Average (EMA), Simple Moving Average (SMA):
  [go-ma](https://github.com/MicahParks/go-ma)
* Relative Strength Index (RSI): [go-rsi](https://github.com/MicahParks/go-rsi)

# Resources
I built and tested this package using these resources:
* [Investopedia](https://www.investopedia.com/terms/a/accumulationdistribution.asp)
* [Invest Excel](https://investexcel.net/accumulation-distribution-line/)
