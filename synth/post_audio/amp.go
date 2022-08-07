package post_audio

func Amp(src []float32, val float32) []float32 {
	// out := make([]float32, len(src))
	for i := range src {
		src[i] *= val
	}
	return src
}
