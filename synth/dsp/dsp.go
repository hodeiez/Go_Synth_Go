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
func RunDSP(dspConf DspConf, osc generator.Osc, noize generator.Osc, cutFreq *float64, resoVal *float64) {

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
	oscs := []generator.Osc{osc, noize}
	for {
		fillBuffers(oscs)

		out = Mixing(out, dspConf, oscs, *cutFreq, *resoVal)
		// write to the stream
		if err := stream.Write(); err != nil {
			log.Printf("error writing to stream : %v\n", err)
		}

	}

}
func fillBuffers(oscs []generator.Osc) {
	for _, o := range oscs {
		if err := o.Osc.Fill(o.Buf); err != nil {
			log.Printf("error filling up the buffer")
		}

	}

}

func Mixing(dst []float32, src DspConf, oscs []generator.Osc, cutFreq float64, resoVal float64) []float32 {

	PreMix(dst, oscs)

	dst = post_audio.Lowpass(dst, cutFreq, 0.0001, 44100)
	//dst = post_audio.Bandpass(dst, cutFreq, 0.0001, 44100, resoVal)
	return dst

}
