package ad

type ADFloat struct {
	prev float64
}

func (ad *ADFloat) Calculate(mvf float64) float64 {
	ad.prev = mvf + ad.prev
	return ad.prev
}

func MFMFloat(close, low, high float64) float64 {
	return ((close - low) - (high - close)) / (high - low)
}

func MFVFloat(mfm, volume float64) float64 {
	return mfm * volume
}
