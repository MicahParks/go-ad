package ad

func MFM(close, low, high float64) float64 {
	return ((close - low) - (high - close)) / (high - low)
}

func MFV(mfm, volume float64) float64 {
	return mfm * volume
}
