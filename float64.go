package ad

// AD represents the state of an Accumulation Distribution Line.
type AD struct {
	prev float64
}

// Input is the required input to calculate a point on the Accumulation Distribution Line.
type Input struct {
	Close  float64
	Low    float64
	High   float64
	Volume float64
}

// New creates a new AD data structure and returns the initial result.
func New(initial Input) (ad *AD, result float64) {
	ad = &AD{}
	return ad, ad.Calculate(initial)
}

// Calculate produces the next point on the Accumulation Distribution Line given the current period's information.
func (ad *AD) Calculate(next Input) float64 {
	if next.High == next.Low || next.Volume == 0 {
		return ad.prev
	}
	ad.prev = (((next.Close - next.Low) - (next.High - next.Close)) / (next.High - next.Low) * next.Volume) + ad.prev
	return ad.prev
}
