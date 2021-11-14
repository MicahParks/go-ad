package ad_test

import (
	"math/big"
	"testing"

	"github.com/MicahParks/go-ad"
)

func BenchmarkAD_Calculate(b *testing.B) {
	var a *ad.AD
	for i := range open {
		if a == nil {
			a, _ = ad.New(closePrices[i], low[i], high[i], volume[i])
		} else {
			_ = a.Calculate(closePrices[i], low[i], high[i], volume[i])
		}
	}
}

func BenchmarkBigAD_Calculate(b *testing.B) {
	var a *ad.BigAD
	for i := range open {
		if a == nil {
			a, _ = ad.NewBig(bigClose[i], bigLow[i], bigHigh[i], bigVolume[i])
		} else {
			_ = a.Calculate(bigClose[i], bigLow[i], bigHigh[i], bigVolume[i])
		}
	}
}

func TestAD_Calculate(t *testing.T) {
	var a *ad.AD
	var result float64
	for i := range open {
		if a == nil {
			a, result = ad.New(closePrices[i], low[i], high[i], volume[i])
		} else {
			result = a.Calculate(closePrices[i], low[i], high[i], volume[i])
		}
		if result != adResults[i] {
			t.FailNow()
		}
	}
}

func TestBigAD_Calculate(t *testing.T) {
	var a *ad.BigAD
	var result *big.Float
	for i := range open {
		if a == nil {
			a, result = ad.NewBig(bigClose[i], bigLow[i], bigHigh[i], bigVolume[i])
		} else {
			result = a.Calculate(bigClose[i], bigLow[i], bigHigh[i], bigVolume[i])
		}
		if result.Cmp(bigADResults[i]) != 0 {
			t.FailNow()
		}
	}
}
