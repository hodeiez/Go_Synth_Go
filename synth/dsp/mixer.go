package dsp

import (
	"hodei/gosynthgo/synth/generator"
)

//TODO: refactor noize volume control to its place
func PreMix(output []float32, buffered []*generator.Tone, voice *generator.Voice) []float32 {

	temp := float32(0.)
	for n := 0; n < len(buffered[0].Osc.Buf.Data); n++ {
		temp = 0.
		for i := range buffered {
			if buffered[i].Type == generator.Noize {
				buffered[i].Osc.Buf.Data[n] *= (*voice.ControlValues.Noize / 100)
			}
			temp += float32(buffered[i].Osc.Buf.Data[n])
		}
		output[n] = temp

	}
	return output
}

func PostMix() {

}
