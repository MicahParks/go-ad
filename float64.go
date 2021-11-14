package ad

// AD represents the state of an Accumulation/Distribution Line.
type AD struct {
	prev float64
}

// New creates a new AD data structure and returns the initial result.
func New(close, low, high, volume float64) (ad *AD, result float64) {
	ad = &AD{}
	return ad, ad.Calculate(close, low, high, volume)
}

// Calculate produces the next point on the Accumulation/Distribution Line given the current period's information.
func (ad *AD) Calculate(close, low, high, volume float64) float64 {
	ad.prev = (((close - low) - (high - close)) / (high - low) * volume) + ad.prev
	return ad.prev
}
