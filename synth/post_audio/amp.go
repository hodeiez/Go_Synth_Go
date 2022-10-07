package post_audio

func Amp(src []float64, val float64) []float64 {
	for i := range src {
		src[i] *= val

	}
	return src
}
