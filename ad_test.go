package ad_test

import (
	"math/big"
	"testing"

	"github.com/MicahParks/go-ad"
)

func BenchmarkAD_Calculate(b *testing.B) {
	var a *ad.AD
	for i := range open {
		input := ad.Input{
			Close:  closePrices[i],
			Low:    low[i],
			High:   high[i],
			Volume: volume[i],
		}

		if a == nil {
			a, _ = ad.New(input)
		} else {
			_ = a.Calculate(input)
		}
	}
}

func BenchmarkBigAD_Calculate(b *testing.B) {
	var a *ad.BigAD
	for i := range bigOpen {
		input := ad.BigInput{
			Close:  bigClose[i],
			Low:    bigLow[i],
			High:   bigHigh[i],
			Volume: bigVolume[i],
		}

		if a == nil {
			a, _ = ad.NewBig(input)
		} else {
			_ = a.Calculate(input)
		}
	}
}

func TestAD_Calculate(t *testing.T) {
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
		if result != adResults[i] {
			t.FailNow()
		}
	}
}

func TestBigAD_Calculate(t *testing.T) {
	var a *ad.BigAD
	var result *big.Float
	for i := range bigOpen {
		input := ad.BigInput{
			Close:  bigClose[i],
			Low:    bigLow[i],
			High:   bigHigh[i],
			Volume: bigVolume[i],
		}

		if a == nil {
			a, result = ad.NewBig(input)
		} else {
			result = a.Calculate(input)
		}
		if result.Cmp(bigADResults[i]) != 0 {
			t.FailNow()
		}
	}
}
