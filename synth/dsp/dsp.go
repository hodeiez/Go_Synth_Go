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

func RunDSP(dspConf DspConf, voices []*generator.Voice) {

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

		fillBuffers(voices)

		Mixing(out, dspConf, voices)
		// write to the stream
		if err := stream.Write(); err != nil {
			log.Printf("error writing to stream : %v\n", err)
		}

	}

}
func fillBuffers(voices []*generator.Voice) {

	for _, v := range voices {
		for _, o := range v.Tones {
			if err := o.Osc.Osc.Fill(o.Osc.Buf); err != nil {
				log.Printf("error filling up the buffer")
			}

		}
	}

}

//TODO: abstract channels to its own object and methods and fix mixing
func Mixing(dst []float32, src DspConf, voices []*generator.Voice) []float32 {

	var oscs []*generator.Tone

	// var temp []float32
	//m := dst
	var audioChannels [][]float32
	audioChannel := make([]float32, len(dst))

	for _, v := range voices {

		oscs = append(oscs, v.Tones...)

		// oscs = append(oscs, v.Noize...)
		//TODO:refactor PREMIX, to BUFFER MIX, and do inside all the channel separation
		// audioChannel = PreMix(dst, oscs, v)
		//these has to go inside premix,
		audioChannel = PreMix(dst, oscs, v)
		audioChannel = v.Filter.RunFilter(audioChannel, 0.0001, 44100)
		audioChannel = post_audio.Amp(audioChannel, float32(*v.ControlValues.Vol/100))
		audioChannels = append(audioChannels, audioChannel)
	}
	for _, a := range audioChannels {
		for i, _ := range a {

			dst[i] += a[i]
		}
	}
	return dst

}
