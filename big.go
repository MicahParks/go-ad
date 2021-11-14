package ad

import (
	"math/big"
)

// BigAD represents the state of an Accumulation/Distribution Line.
type BigAD struct {
	prev *big.Float
}

// NewBig creates a new AD data structure and returns the initial result.
func NewBig(close, low, high, volume *big.Float) (ad *BigAD, result *big.Float) {
	ad = &BigAD{prev: big.NewFloat(0)}
	return ad, ad.Calculate(close, low, high, volume)
}

// Calculate produces the next point on the Accumulation/Distribution Line given the current period's information.
func (ad *BigAD) Calculate(close, low, high, volume *big.Float) *big.Float {
	ad.prev.Add(new(big.Float).Mul(new(big.Float).Quo(new(big.Float).Sub(new(big.Float).Sub(close, low), new(big.Float).Sub(high, close)), new(big.Float).Sub(high, low)), volume), ad.prev)
	return new(big.Float).Copy(ad.prev)
}
