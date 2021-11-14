package ad

type FloatAD struct {
	prev float64
}

func (ad *FloatAD) Calculate(mvf float64) float64 {
	ad.prev = mvf + ad.prev
	return ad.prev
}

func FloatMFM(close, low, high float64) float64 {
	return ((close - low) - (high - close)) / (high - low)
}

func FloatMFV(mfm, volume float64) float64 {
	return mfm * volume
}
