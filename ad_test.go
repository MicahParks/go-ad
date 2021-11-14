package ad_test

import (
	"math/big"
	"testing"

	"github.com/MicahParks/go-ad"
)

func BenchmarkAD_Calculate(b *testing.B) {
	var adLine *ad.AD
	for i := range open {
		input := ad.Input{
			Close:  closePrices[i],
			Low:    low[i],
			High:   high[i],
			Volume: volume[i],
		}

		if adLine == nil {
			adLine, _ = ad.New(input)
		} else {
			_ = adLine.Calculate(input)
		}
	}
}

func BenchmarkBigAD_Calculate(b *testing.B) {
	var adLine *ad.BigAD
	for i := range bigOpen {
		input := ad.BigInput{
			Close:  bigClose[i],
			Low:    bigLow[i],
			High:   bigHigh[i],
			Volume: bigVolume[i],
		}

		if adLine == nil {
			adLine, _ = ad.NewBig(input)
		} else {
			_ = adLine.Calculate(input)
		}
	}
}

func TestAD_Calculate(t *testing.T) {
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
		if result != adResults[i] {
			t.FailNow()
		}
	}
}

func TestBigAD_Calculate(t *testing.T) {
	var adLine *ad.BigAD
	var result *big.Float
	for i := range bigOpen {
		input := ad.BigInput{
			Close:  bigClose[i],
			Low:    bigLow[i],
			High:   bigHigh[i],
			Volume: bigVolume[i],
		}

		if adLine == nil {
			adLine, result = ad.NewBig(input)
		} else {
			result = adLine.Calculate(input)
		}
		if result.Cmp(bigADResults[i]) != 0 {
			t.FailNow()
		}
	}
}
