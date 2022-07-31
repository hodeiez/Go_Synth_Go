package dsp

import "hodei/gosynthgo/synth/generator"

func PreMix(output []float32, buffered []*generator.Osc) []float32 {

	temp := float32(0.)
	for n := 0; n < len(buffered[0].Buf.Data); n++ {
		temp = 0.
		for i := range buffered {

			temp += float32(buffered[i].Buf.Data[n])
		}
		output[n] = temp

	}
	return output
}
func PreMixAfterFilt(output []float32, buffered []float32) []float32 {

	temp := float32(0.)
	for n := 0; n < len(buffered); n++ {
		temp = 0.
		for i := range buffered {

			temp += buffered[i]
		}
		output[n] = temp

	}
	return output
}

func PostMix() {

}
