package ad

import (
	"math/big"
)

type ADBig struct {
	prev *big.Float
}

func (ad *ADBig) Calculate(mvf *big.Float) *big.Float {
	if ad.prev == nil {
		ad.prev = big.NewFloat(0)
	}
	ad.prev.Add(mvf, ad.prev)
	return new(big.Float).Copy(ad.prev)
}

func MFMBig(close, low, high *big.Float) *big.Float {
	return new(big.Float).Quo(new(big.Float).Sub(new(big.Float).Sub(close, low), new(big.Float).Sub(high, close)), new(big.Float).Sub(high, low))
}

func MFVBig(mfm, volume *big.Float) *big.Float {
	return new(big.Float).Mul(mfm, volume)
}
