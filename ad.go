package ad

type AD struct {
	prev float64
}

func (ad *AD) Calculate(mvf float64) float64 {
	ad.prev = mvf + ad.prev
	return ad.prev
}
