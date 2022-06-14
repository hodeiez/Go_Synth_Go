package dsp

import (
	"hodei/gosynthgo/synth/generator"
	"hodei/gosynthgo/synth/post_audio"
	"log"

	"github.com/gordonklaus/portaudio"
)

//** for now we run just one oscillator
type DspConf struct {
	BufferSize int
}

//TODO: review and fix the volume and amplitude
func RunDSP(dspConf DspConf, osc generator.Osc, noize generator.Noise, cutFreq *float64, resoVal *float64) {

	portaudio.Initialize()
	api, _ := portaudio.HostApis()

	for _, ap := range api {
		log.Println(*ap)
	}
	defer portaudio.Terminate()
	out := make([]float32, dspConf.BufferSize)

	stream, err := portaudio.OpenDefaultStream(0, 2, 44100, len(out), &out)
	if err != nil {
		log.Fatal(err)
	}

	defer stream.Close()

	if err := stream.Start(); err != nil {
		log.Fatal(err)
	}
	defer stream.Stop()

	for {

		if err := osc.Osc.Fill(osc.Buf); err != nil {
			log.Printf("error filling up the buffer")
		}
		if err := noize.Osc.Fill(noize.Buf); err != nil {
			log.Printf("error filling up the buffer")
		}
		// populate the out buffer

		// for _, oscillators := range dspConf.VM.Voices {
		// 	if err := oscillators.Oscillator.Osc.Fill(oscillators.Oscillator.Buf); err != nil {
		// 		log.Printf("error filling up the buffer")
		// 	}

		// }

		// for _, oscillators2 := range dspConf.VM.Voices {
		// 	if err := oscillators2.Oscillator2.Osc.Fill(oscillators2.Oscillator2.Buf); err != nil {
		// 		log.Printf("error filling up the buffer")
		// 	}

		// }

		out = Mixing(out, dspConf, osc, noize, *cutFreq, *resoVal)
		// write to the stream
		if err := stream.Write(); err != nil {
			log.Printf("error writing to stream : %v\n", err)
		}

	}

}

func Mixing(dst []float32, src DspConf, osc generator.Osc, noize generator.Noise, cutFreq float64, resoVal float64) []float32 {
	el := [][]float64{osc.Buf.Data, noize.Buf.Data} //TODO: wrapUp in a struct type??
	PreMix(dst, el)

	dst = post_audio.Lowpass(dst, cutFreq, 0.0001, 44100)
	//dst = post_audio.Bandpass(dst, cutFreq, 0.0001, 44100, resoVal)
	return dst

}
