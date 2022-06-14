package dsp

func PreMix(output []float32, buffered [][]float64) {
	temp := float32(0.)
	for n := 0; n < len(buffered[0]); n++ {
		temp = 0.
		for i := range buffered {

			temp += float32(buffered[i][n])
		}
		output[n] = temp

	}
}

func PostMix() {

}
