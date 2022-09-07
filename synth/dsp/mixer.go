package dsp

import (
	"hodei/gosynthgo/synth/generator"
)

//TODO: refactor noize volume control to its place
func PreMix(output []float64, buffered []*generator.Tone, voice *generator.Voice) []float64 {

	temp := float64(0.)
	for n := 0; n < len(buffered[0].Osc.Buf.Data); n++ {
		temp = 0.
		for i := range buffered {
			if buffered[i].Type == generator.Noize {
				buffered[i].Osc.Buf.Data[n] *= (*voice.ControlValues.Noize / 100)
			}
			temp += buffered[i].Osc.Buf.Data[n]
		}
		output[n] = temp

	}
	return output
}

// func PreMix(output []float32, tones []*generator.Tone, voice *generator.Voice) []float32 {

// 	temp := float32(0.)
// 	for n := 0; n < len(tones[0].Osc.Buf.Data); n++ {
// 		temp = 0.
// 		for i := range tones {
// 			if tones[i].Type == generator.Noize {
// 				tones[i].Osc.Buf.Data[n] *= (*voice.ControlValues.Noize / 100)
// 			}
// 			temp += float32(tones[i].Osc.Buf.Data[n])
// 		}
// 		output[n] = temp

// 	}
// 	return output
// }

func PostMix() {

}
