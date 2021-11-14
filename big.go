package ad

import (
	"math/big"
)

type BigAD struct {
	prev *big.Float
}

func (ad *BigAD) Calculate(mvf *big.Float) *big.Float {
	if ad.prev == nil {
		ad.prev = big.NewFloat(0)
	}
	ad.prev.Add(mvf, ad.prev)
	return new(big.Float).Copy(ad.prev)
}

func BigMFM(close, low, high *big.Float) *big.Float {
	return new(big.Float).Quo(new(big.Float).Sub(new(big.Float).Sub(close, low), new(big.Float).Sub(high, close)), new(big.Float).Sub(high, low))
}

func BigMFV(mfm, volume *big.Float) *big.Float {
	return new(big.Float).Mul(mfm, volume)
}
