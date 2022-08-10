package post_audio

func Amp(src []float32, val float32) []float32 {
	for i := range src {
		src[i] *= val
	}
	return src
}
