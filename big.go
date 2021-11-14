package ad

import (
	"math/big"
)

// BigAD represents the state of an Accumulation Distribution Line.
type BigAD struct {
	prev *big.Float
}

// BigInput is the required input to calculate a point on the Accumulation Distribution Line.
type BigInput struct {
	Close  *big.Float
	Low    *big.Float
	High   *big.Float
	Volume *big.Float
}

// NewBig creates a new AD data structure and returns the initial result.
func NewBig(input BigInput) (ad *BigAD, result *big.Float) {
	ad = &BigAD{prev: big.NewFloat(0)}
	return ad, ad.Calculate(input)
}

// Calculate produces the next point on the Accumulation Distribution Line given the current period's information.
func (ad *BigAD) Calculate(next BigInput) *big.Float {
	ad.prev.Add(new(big.Float).Mul(new(big.Float).Quo(new(big.Float).Sub(new(big.Float).Sub(next.Close, next.Low), new(big.Float).Sub(next.High, next.Close)), new(big.Float).Sub(next.High, next.Low)), next.Volume), ad.prev)
	return new(big.Float).Copy(ad.prev)
}
