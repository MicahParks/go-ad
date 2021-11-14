# go-ad
The Accumulation Distribution Line technical analysis algorithm implemented in Golang.

## Usage
For full examples, please see the `examples` directory.

### Step 1
Gather the input for the current period.
```go
input := ad.Input{
	Close:  closePrices[i],
	Low:    low[i],
	High:   high[i],
	Volume: volume[i],
}
```

### Step 2
Create the A/D data structure and get the first result
```go
a, result := ad.New(input)
```

### Step 3
Use the next period's data to calculate the next point on the A/D line. Repeat for all periods.
```go
result = a.Calculate(input)
```

## Partially complete example
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
	var a *ad.AD
	var result float64
	for i := range open {
		input := ad.Input{
			Close:  closePrices[i],
			Low:    low[i],
			High:   high[i],
			Volume: volume[i],
		}

		if a == nil {
			a, result = ad.New(input)
		} else {
			result = a.Calculate(input)
		}
		logger.Printf("Index: %d AD: %.4f", i, result)
	}
}
```

## Testing
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

## Resources
I built and tested this package using these resources:
* [Investopedia](https://www.investopedia.com/terms/a/accumulationdistribution.asp)
* [Invest Excel](https://investexcel.net/accumulation-distribution-line/)
