package dsp

import (
	"hodei/gosynthgo/synth/generator"

	"log"

	"github.com/gordonklaus/portaudio"
)

//** for now we run just one oscillator
type DspConf struct {
	BufferSize int
}

//TODO: review and fix the volume and amplitude
func RunDSP(dspConf DspConf, voice generator.Voice) {

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
		fillBuffers(voice)

		out = Mixing(out, dspConf, voice)
		// write to the stream
		if err := stream.Write(); err != nil {
			log.Printf("error writing to stream : %v\n", err)
		}

	}

}
func fillBuffers(voice generator.Voice) {
	oscs := []generator.Osc{*voice.Osc[0], *voice.Noize[0]}
	for _, o := range oscs {
		if err := o.Osc.Fill(o.Buf); err != nil {
			log.Printf("error filling up the buffer")
		}

	}

}

func Mixing(dst []float32, src DspConf, voice generator.Voice) []float32 {
	oscs := []generator.Osc{*voice.Osc[0], *voice.Noize[0]}
	PreMix(dst, oscs)

	dst = voice.Filter.RunFilter(dst, 0.0001, 44100)
	return dst

}
